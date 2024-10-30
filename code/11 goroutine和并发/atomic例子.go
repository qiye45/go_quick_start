package main

//票数计数器
//想象一个场景，你和朋友们一起参加一个公益活动，每个人都在往一个募捐箱里投票，你们每投一票都要记录下来。这时就有个计数器来记录总票数。为了避免同时写入时发生错误（比如两个朋友同时读到一样的票数然后都写入，导致实际票数少了一票），我们可以用原子操作来确保每次加票数是安全的。

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var totalVotes int32  // 总票数
	var wg sync.WaitGroup // 用于等待所有 goroutine 完成
	unsafeTotalvotes := 0

	// 模拟 100 个朋友同时投票
	for i := 0; i < 1000; i++ {
		wg.Add(1) // 增加等待组计数
		go func() {
			defer wg.Done()                 // goroutine 结束后减少等待组计数
			atomic.AddInt32(&totalVotes, 1) // 原子增加票数
			unsafeTotalvotes += 1
		}()
	}

	wg.Wait()                            // 等待所有 goroutine 完成
	fmt.Printf("总票数为: %d\n", totalVotes) //1000
	fmt.Println(unsafeTotalvotes)        //987
}
