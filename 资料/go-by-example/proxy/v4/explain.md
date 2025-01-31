1. 使用channel进行数据转发：

```go
func forward(src, dst net.Conn) {
    done := make(chan bool)
    go func() {
        io.Copy(dst, src)
        done <- true
    }()
    go func() {
        io.Copy(src, dst)
        done <- true
    }()
    <-done
}
```

2. 使用sync.WaitGroup等待goroutine完成：

```go
func forward(src, dst net.Conn) {
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        io.Copy(dst, src)
    }()
    go func() {
        defer wg.Done()
        io.Copy(src, dst)
    }()
    wg.Wait()
}
```

3. 使用超时控制：

```go
func connectWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
    return net.DialTimeout(network, address, timeout)
}
```

4. 添加错误处理和重试机制：

```go
func connectWithRetry(network, address string, maxRetries int) (net.Conn, error) {
    var err error
    for i := 0; i < maxRetries; i++ {
        conn, err := net.Dial(network, address)
        if err == nil {
            return conn, nil
        }
        time.Sleep(time.Second * time.Duration(i+1))
    }
    return nil, err
}
```
