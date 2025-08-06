你提供的是 Go 的基准测试代码（`BenchmarkHello` 和 `BenchmarkParallel`）。要运行这些基准测试，你可以使用以下命令：

---

### ✅ 运行基准测试的命令：

```bash
go test -bench=.
```

---

### 📌 说明：

* `-bench=.` 表示运行当前包中所有以 `Benchmark` 开头的基准测试函数。
* 如果你的文件名不是以 `_test.go` 结尾（例如 `main.go`），你需要：

  1. 把基准测试代码放在一个名为 `xxx_test.go` 的文件中，例如 `main_test.go`。
  2. 并确保文件位于同一个包中。

---
### 📈 可选参数：

* 指定运行时间更长的测试（默认每个函数约运行 1 秒）：

```bash
go test -bench=. -benchtime=5s
```

* 显示内存分配信息：

```bash
go test -bench=. -benchmem
```

---

需要我帮你分析 `BenchmarkParallel` 的性能含义或 goroutine 行为吗？
