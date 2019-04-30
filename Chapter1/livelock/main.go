package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
ライブロック : 並行操作を行なっているものの、その操作はプログラムの状態を全く進めていない。
廊下や道でお互いが避けあって、同じ方向に避けてサッカーのフェイントみたいになるような状態を想像するといいかも
*/
func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()
	// 人間が同じ歩調で動くことをシュミレート
	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}
	//避ける動作を関数化
	//ある人がある方向に動いてみてうまく避けられたかを返す。
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, "  %v\n", dirName)
		atomic.AddInt32(dir, 1) // 最初にある方向に動こうとしていることを、その方向に動く人数を１増やすことで宣言します。
		takeStep()              // ライブロックの例を示すために書く人間は同じスピードで同じ歩調で動かなければならない。
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1) // 同じ方向に歩んでしまい諦めて一歩戻る動作
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:\n", name)
		for i := 0; i < 5; i++ { //ライブロックのあるプログラムは上限がないが故に問題が発生する
			if tryLeft(&out) || tryRight(&out) { //まずある人が左に行こうとしたとして、それが失敗すると右に行く
				return
			}
		}
		fmt.Fprintf(&out, "\n%v toses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup //この変数はプログラムが両方の人間がお互いにすれ違えるようになる、あるいはすれ違うのを諦めるまで待つ方法を提供する
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()

}
