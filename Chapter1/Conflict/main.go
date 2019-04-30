package main

import (
	"fmt"
	"time"
)

func main() {
	isNotConflict()
	isConflict()
}

func isNotConflict() {
	var data int
	// dataを検証するのと同時にgo routineが動くのでdata変数の値を参照する際に競合状態にあると言える。
	go func() {
		data++
	}()
	if data == 0 {
		// 競合状態であるためdata == 0となる
		fmt.Println("There is zero in data")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}

func isConflict() {
	var data int
	go func() {
		data++
	}()
	// 待つことにより並列で動いている変数増減の評価を待つことができる。
	// しかしこの待ち方は悪手と言える
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Println("There is zero in data")
	} else {
		// 待機するためこちらが表示される
		fmt.Printf("the value is %v.\n", data)
	}
}
