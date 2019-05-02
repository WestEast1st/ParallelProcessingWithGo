package main

import (
	"fmt"
	"sync"
)

func main() {
	sampleOnce1()
	sampleOnce2()
	sampleOnceDedlock()
}

func sampleOnce1() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment) //関数を確実に一度だけ呼び出すことのできる機能を利用し関数を呼び出す
		}()
	}
	increments.Wait()

	fmt.Printf("Count is %d\n", count)
}

func sampleOnce2() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once

	once.Do(increment)
	once.Do(decrement)
	fmt.Printf("Count: %d\n", count)
}

func sampleOnceDedlock() {
	var onceA, onceB sync.Once
	var initB func()

	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) } //この関数の呼び出しは最後のonceA.DoがinitAを呼び出すまで値を返すことはしません。
	onceA.Do(initA)                    // deadlock！！！！！！！！！！！！！
}
