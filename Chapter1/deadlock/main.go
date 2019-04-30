package main

import (
	"fmt"
	"sync"
	"time"
)

// 数値とメモリアクセス関連の構造体
type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()         // 流入してくるあたいのためにクリティカルセクションに入る
		defer v1.mu.Unlock() //printSumが終了する前にクリティカルセクションを抜ける

		//こいつが根本原因
		time.Sleep(2 * time.Second) //処理の負荷をシュミュレートするための一定時間スリープ
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b) //&aが第一引数でprintSumを実行する際に、内部で&aがlockをかける
	go printSum(&b, &a) //&bが第一引数でprintSumを実行する際に、内部で&bがlockをかける
	wg.Wait()
}

/*
エラーメッセージ
fatal error: all goroutines are asleep - deadlock!
*/
