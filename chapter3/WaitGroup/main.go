package main

/*
WaitGroupはひとまとまりの並行処理があった時、その結果を気にしない、もしくは他の結果を取集する手段がある場合に、
それらの処理の完了を待つ手段として非常に有効です。
どちらの前提にも当てはまらない場合には代わりにSelect文を使うことをお勧めします。
*/
import (
	"fmt"
	"sync"
	"time"
)

/*
<1>Addを引数として渡して呼び出し、1つのGorutineが起動したことを表す
<2>Doneはdeferキーワードを使って呼び出して、ゴルーチンのクロージャが終了する前にWaitGroupに終了することを伝えるようにする。
<3>Waitを呼び出し、全てのゴルーチンが終了するまでメインゴルーチンをブロックする
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(1) //<1>
	go func() {
		defer wg.Done() //<2>
		fmt.Println("1st goroutine sleeping ...")
		time.Sleep(1)
	}()
	wg.Add(1) //<1>
	go func() {
		defer wg.Done() //<2>
		fmt.Println("2nd goroutine sleeping ...")
		time.Sleep(2)
	}()
	wg.Wait() //<3>
	fmt.Println("All goroutine complete.")

	//Addの呼び出しはできる限り監視対象のゴルーチンの直前に書くことが慣習です。
	//しかし、時に関連するゴルーチンの呼び出しを一度に監視するために、Addの呼び出しが行われているのを見ることもあるでしょう。
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v\n", id)
	}
	const numGreeters = 5
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
