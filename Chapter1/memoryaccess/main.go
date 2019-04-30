package main

/*
競合状態と不可分操作を確保するためにメモリアクセスの同期をおこなう
Goでの綺麗な書き方とはかけ離れているが、メモリの同期を簡潔に表れいるのでこのコードを使用する。
*/
import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex // sync.Mutexによってメモリの変数を制御する
	var data int
	go func() {
		memoryAccess.Lock() // go rutine内でメモリに対する排他的アクセス権を取得
		data++
		fmt.Printf("the value is %v.\n", data)
		memoryAccess.Unlock() //アクセス権の破棄
	}()
	memoryAccess.Lock()
	if data == 0 {
		fmt.Println("There is zero in data")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}
