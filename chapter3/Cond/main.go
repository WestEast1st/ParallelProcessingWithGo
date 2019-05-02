package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Cond型
ゴルーチンが待機したりイベントの発生を知らせるためのランデブーポイントを提供する型

二つ以上のゴルーチン間で、それが発生したということ以外の情報がない任意のシグナルを指します。ゴルーチン上で処理を続ける前にこうした信号を受け取りたいということが非常によくあります。
*/
func main() {
	sampleBroadCast()
}

func sampleCond() {
	c := sync.NewCond(&sync.Mutex{})    //<1> 標準のsync.MutexをLockerとして使って条件を作成
	queue := make([]interface{}, 0, 10) //<2> キューを設定

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        //<8> 再度条件のクリティカルセクションに入って、条件に合った形でデータ修正
		queue = queue[1:] //<9> キューの取り出し
		fmt.Println("Removed from queue.")
		c.L.Unlock() //<10>取り出しが終了したのでロック解除
		c.Signal()   //<11>　3へシグナルを送信
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            //<3> 条件であるLockerのLockメソッドを呼び出してクリティカルセクションに入る
		for len(queue) == 2 { //<4> キューの内容量を確認し、必要に応じて待機をする。今回の場合は容量2の時に停止する
			c.Wait() //<5> waitでゴルーチンをブロックする
		}
		fmt.Println("Adding to queue.")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) //<6> １秒後にキューを排出する
		c.L.Unlock()                        //<7> キューが追加できたのでクリティカルセクションを解除
	}
}

func sampleBroadCast() {
	type Button struct { //<1>Clickedという条件を含んでいるButton型を製作
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) { //<2>条件別にシグナルを扱う関数を登録する関数。かくしゅはんどらーは、それぞれのゴルーチン場で動作する
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup //<3>全てのハンドラーにマウスのおボタンがクリックが行われたことを通知します
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() { //<4>
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() { //<5>
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() { //<6>
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast() //<7>

	clickRegistered.Wait()
}
