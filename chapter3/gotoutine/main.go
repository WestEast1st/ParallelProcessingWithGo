package main

import (
	"fmt"
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
	sample2()
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
