**1. 使用 channel 进行数据转发：**

````go
import (
	"io"
	"net"
)

// forward 函数将数据从 src 连接转发到 dst 连接，并反之。
// 它使用两个 goroutine 并通过 channel 来等待转发完成。
func forward(src, dst net.Conn) {
	// 创建一个无缓冲的 bool 类型 channel，用于在转发完成后发送信号。
	done := make(chan bool)

	// 启动第一个 goroutine，负责将数据从 src 连接复制到 dst 连接。
	go func() {
		// io.Copy 函数从 src 读取所有数据并写入到 dst，直到遇到 EOF 或发生错误。
		// 它会阻塞直到复制完成。
		_, err := io.Copy(dst, src)
		if err != nil {
			// 在实际应用中，你可能需要记录或处理这个错误。
			// fmt.Println("io.Copy(dst, src) 错误:", err)
		}
		// 当从 src 到 dst 的复制完成后，向 done channel 发送 true。
		// 这表示这个方向的转发已经完成。
		done <- true
	}()

	// 启动第二个 goroutine，负责将数据从 dst 连接复制到 src 连接。
	go func() {
		// 同样使用 io.Copy 将数据从 dst 读取并写入到 src。
		_, err := io.Copy(src, dst)
		if err != nil {
			// 同样，在实际应用中可能需要处理这个错误。
			// fmt.Println("io.Copy(src, dst) 错误:", err)
		}
		// 当从 dst 到 src 的复制完成后，向 done channel 发送 true。
		// 这表示这个方向的转发也已经完成。
		done <- true
	}()

	// 主 goroutine 会在这里阻塞，直到从 done channel 接收到两个 true 值。
	// 这确保了两个方向的数据转发 goroutine 都已经完成。
	<-done
	<-done // 接收两次，因为有两个 goroutine 发送信号
	// 在实际应用中，可以只接收一次，然后在另一个 goroutine 关闭连接。
	// 这里为了简化示例，等待两个 goroutine 都完成。
}
````

**注释解释：**

*   **`done := make(chan bool)`**: 创建了一个无缓冲的 channel。无缓冲的 channel 意味着发送操作会阻塞，直到有接收者准备好接收。在这里，它用于在两个数据转发的 goroutine 完成后通知主 goroutine。
*   **`go func() { ... }()`**: 启动一个新的 goroutine。这允许 `io.Copy` 在后台运行，不会阻塞主线程。
*   **`io.Copy(dst, src)`**: 这是 Go 标准库 `io` 包中的一个函数，它将 `src`（一个 `io.Reader`，在这里是 `net.Conn`）中的所有数据读取出来，并写入到 `dst`（一个 `io.Writer`，在这里也是 `net.Conn`）。这个函数会一直执行，直到遇到 `src` 的 EOF (End Of File) 或者发生错误。
*   **`done <- true`**: 当 `io.Copy` 完成（无论是正常结束还是因为错误），这个 goroutine 会向 `done` channel 发送一个 `true` 值。这表明这个方向的数据转发已经结束。
*   **`<-done`**: 主 goroutine 在这里等待从 `done` channel 接收一个值。由于我们启动了两个 goroutine，并且每个都在完成后向 `done` channel 发送一个值，所以我们需要接收两次才能确保两个方向的转发都已完成。

**2. 使用 sync.WaitGroup 等待 goroutine 完成：**

````go
import (
	"io"
	"net"
	"sync"
)

// forward 函数将数据从 src 连接转发到 dst 连接，并反之。
// 它使用 sync.WaitGroup 来等待转发的两个 goroutine 完成。
func forward(src, dst net.Conn) {
	// 创建一个 sync.WaitGroup 类型的变量 wg。
	// WaitGroup 用于等待一组 goroutine 完成。
	var wg sync.WaitGroup

	// 使用 wg.Add(2) 来增加等待组的计数器，表示我们有两个 goroutine 需要等待。
	wg.Add(2)

	// 启动第一个 goroutine，负责将数据从 src 连接复制到 dst 连接。
	go func() {
		// 使用 defer wg.Done() 来确保在函数退出时将 WaitGroup 的计数器减一。
		// 即使函数中发生 panic，defer 也会执行。
		defer wg.Done()
		// io.Copy 函数从 src 读取所有数据并写入到 dst，直到遇到 EOF 或发生错误。
		_, err := io.Copy(dst, src)
		if err != nil {
			// 在实际应用中，你可能需要记录或处理这个错误。
			// fmt.Println("io.Copy(dst, src) 错误:", err)
		}
		// 当 io.Copy 完成并且 wg.Done() 被调用后，WaitGroup 的计数器会减一。
	}()

	// 启动第二个 goroutine，负责将数据从 dst 连接复制到 src 连接。
	go func() {
		// 同样使用 defer wg.Done() 来在函数退出时减少 WaitGroup 的计数器。
		defer wg.Done()
		// io.Copy 函数从 dst 读取所有数据并写入到 src，直到遇到 EOF 或发生错误。
		_, err := io.Copy(src, dst)
		if err != nil {
			// 同样，在实际应用中可能需要处理这个错误。
			// fmt.Println("io.Copy(src, dst) 错误:", err)
		}
		// 当 io.Copy 完成并且 wg.Done() 被调用后，WaitGroup 的计数器会再次减一。
	}()

	// wg.Wait() 会阻塞当前 goroutine（在这里是 forward 函数的主 goroutine），
	// 直到 WaitGroup 的计数器变为零。这表示所有通过 wg.Add 增加的 goroutine 都已经调用了 wg.Done()。
	wg.Wait()
	// 当 wg.Wait() 返回时，表示两个数据转发的 goroutine 都已经执行完毕。
}
````

**注释解释：**

*   **`var wg sync.WaitGroup`**: 声明一个 `sync.WaitGroup` 类型的变量。`WaitGroup` 提供了一种简单的机制来等待一组 goroutine 完成。
*   **`wg.Add(2)`**: 将 `WaitGroup` 的内部计数器增加 2，表示有两个 goroutine 需要等待。
*   **`defer wg.Done()`**: `defer` 关键字用于延迟一个函数的执行，直到周围的函数返回。在这里，无论 `io.Copy` 是否成功完成或者发生错误，`wg.Done()` 都会被调用。`wg.Done()` 会将 `WaitGroup` 的计数器减 1。
*   **`wg.Wait()`**: 调用 `wg.Wait()` 的 goroutine 会阻塞，直到 `WaitGroup` 的内部计数器变为 0。这意味着所有通过 `wg.Add` 增加的 goroutine 都已经调用了 `wg.Done()`。

**3. 使用超时控制：**

````go
import (
	"net"
	"time"
)

// connectWithTimeout 函数尝试在指定的网络和地址上建立连接，并在指定的超时时间内完成。
// 如果在超时时间内未能建立连接，则返回错误。
func connectWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	// net.DialTimeout 函数是 net.Dial 的一个变体，它接受一个额外的 timeout 参数。
	// timeout 参数指定了等待连接建立的最长时间。
	// 如果在 timeout 时间内连接成功建立，它将返回一个 net.Conn 对象和 nil 错误。
	// 如果在 timeout 时间内连接失败或超时，它将返回 nil 连接和一个错误，该错误通常包含超时的信息。
	conn, err := net.DialTimeout(network, address, timeout)
	// 返回建立的连接和可能发生的错误。
	return conn, err
}
````

**注释解释：**

*   **`net.DialTimeout(network, address, timeout)`**: 这是 Go 标准库 `net` 包提供的用于建立网络连接的函数。
    *   `network`: 指定网络类型，例如 "tcp"、"udp"、"unix" 等。
    *   `address`: 指定要连接的地址，格式取决于 `network` 类型（例如 "127.0.0.1:8080" 对于 TCP）。
    *   `timeout`: 一个 `time.Duration` 类型的值，表示等待连接建立的最大时长。如果在该时间内连接未建立成功，`DialTimeout` 将返回一个错误。
*   该函数直接返回 `net.DialTimeout` 的结果，即一个 `net.Conn` 接口类型的连接（如果成功）和一个 `error` 类型的错误（如果失败或超时）。

**4. 添加错误处理和重试机制：**

````go
import (
	"fmt"
	"net"
	"time"
)

// connectWithRetry 函数尝试在指定的网络和地址上建立连接，如果失败，则会进行最多 maxRetries 次重试。
// 每次重试之间会有递增的延迟。
func connectWithRetry(network, address string, maxRetries int) (net.Conn, error) {
	// 声明一个 error 类型的变量 err，用于存储连接过程中可能发生的错误。
	var err error

	// 使用 for 循环进行重试，循环次数最多为 maxRetries。
	for i := 0; i < maxRetries; i++ {
		// 尝试建立网络连接。
		conn, err := net.Dial(network, address)
		// 检查连接是否成功建立（err 是否为 nil）。
		if err == nil {
			// 如果连接成功，立即返回连接和 nil 错误。
			fmt.Printf("成功连接到 %s:%s (尝试次数: %d)\n", network, address, i+1)
			return conn, nil
		}

		// 如果连接失败，打印错误信息（在实际应用中可能需要更详细的日志记录）。
		fmt.Printf("连接 %s:%s 失败 (尝试次数: %d): %v\n", network, address, i+1, err)

		// 如果不是最后一次尝试，则等待一段时间后重试。
		if i < maxRetries-1 {
			// 计算等待时间，每次重试的等待时间都会增加（1秒 * (i+1)）。
			waitDuration := time.Second * time.Duration(i+1)
			fmt.Printf("等待 %v 后重试...\n", waitDuration)
			// 使用 time.Sleep 暂停当前 goroutine 指定的时间。
			time.Sleep(waitDuration)
		}
	}

	// 如果所有重试都失败了，返回 nil 连接和最后一次遇到的错误。
	fmt.Printf("所有 %d 次重试均失败，无法连接到 %s:%s\n", maxRetries, network, address)
	return nil, err
}
````

**注释解释：**

*   **`var err error`**: 声明一个变量来存储连接过程中可能出现的错误。
*   **`for i := 0; i < maxRetries; i++`**: 使用一个 `for` 循环来控制重试的次数。
*   **`conn, err := net.Dial(network, address)`**: 尝试建立网络连接。`net.Dial` 会返回一个 `net.Conn` 对象（如果成功）和一个 `error` 对象（如果失败）。
*   **`if err == nil { ... }`**: 检查 `err` 是否为 `nil`。如果为 `nil`，表示连接成功，函数立即返回建立的连接和 `nil` 错误。
*   **`fmt.Printf(...)`**: 用于打印连接尝试的信息和错误（在实际应用中，应该使用更完善的日志记录机制）。
*   **`if i < maxRetries-1 { ... }`**: 只有在不是最后一次尝试时才进行等待。
*   **`waitDuration := time.Second * time.Duration(i+1)`**: 计算重试之间的等待时间。这里使用了简单的指数退避策略，每次失败后等待的时间都会增加。
*   **`time.Sleep(waitDuration)`**: 使当前 goroutine 暂停执行指定的时间。
*   如果循环结束后仍然没有成功建立连接，函数将返回 `nil` 连接和最后一次遇到的错误。

这些详细的注释应该能帮助你理解每个代码片段的功能和实现细节。在实际应用中，你可能需要根据具体的需求进行更完善的错误处理、日志记录和配置管理。
