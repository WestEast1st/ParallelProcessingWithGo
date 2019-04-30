package main

import (
	"fmt"
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
	go hello()
	sayHello := func() {
		fmt.Println("HELLO!")
	}
	go sayHello()
}

func hello() {
	fmt.Println("Hello!")
}
