package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
goroutineについて学ぶための章
goroutineはGoのプログラムの中で最も基本的な構成単位である。したがって、それがなぜ動くのかを理解するのは大変重要なことである。事実として、どのようなgoのプログラムの中にでも最低一つは含まれていると考えられる。

goroutineとは、
単純に言えば、他のコードに対して並行に実行している関数のことである(必ずしも並列ではない)。
```
go function()
```
で起動できる
*/

func main() {
	sample4()
}

func sample1() {
	go hello()
	sayHello := func() {
		fmt.Println("HELLO!")
	}
	go sayHello()
}

func hello() {
	fmt.Println("Hello!")
}

func sample2() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done() //終了を伝える
		fmt.Println("hello!")
	}
	fmt.Println("開始")
	wg.Add(1)
	go sayHello()
	wg.Wait() //終了の待機
	fmt.Println("終了")
}

//go routineのアドレス空間の確認
func sample3() {
	/*
	  go routineを利用したクロージャでのアドレス空間が同一であることを確認するためのサンプル
	*/
	var wg sync.WaitGroup
	fmt.Println("アドレス空間が変更される場合")
	fmt.Println("START")
	for _, salutation := range []string{"hello", "greetings", "good bay"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
	/*
	  good bay
	  good bay
	  good bay
	  と表示される
	  go routineが開始する前にループが終了するのでsalutationが書き換わってしまう。
	*/
	fmt.Println("END")
	fmt.Println("クロージャに変数として渡しアドレス空間を別に用意する方法")
	fmt.Println("START")
	for _, salutation := range []string{"hello", "greetings", "good bay"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
	fmt.Println("END")
}

func sample4() {
	memConsmed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } //計算のために沢山のgo routineをメモリ上に配置、絶対に終了しないgo routineが必要。
	const numGoroutines = 1e4         // ここで定義するgoroutineの数を定義する。漸近的にルーティンの数を増やす
	wg.Add(numGoroutines)
	before := memConsmed() //生成前のメモリ消費量
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsmed() //生成後のメモリ消費量
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
