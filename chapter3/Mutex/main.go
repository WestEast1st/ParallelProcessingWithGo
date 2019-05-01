package main

import (
	"fmt"
	"sync"
)

/*
Mutexとは？
相互排除を意味する"Mutual Exclusion"の略で、プログラム内のクリティカルセクションを保護する方法の一つです。

*/

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()         //Mutexインスタンスで保護されたクリティカルセクションの戦友を要求する。今回の場合はcount変数である。
		defer lock.Unlock() //lock保護をしているクリティカルセクションを解放する
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	//インクリメント
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	//デクリメント
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}
