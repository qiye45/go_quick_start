# goroutine和并发

在Go中如何实现并发，答案就是使用**goroutine**。在Go语言中，一个独立运行的任务就被称为一个 **goroutine**。`goroutine`的创建效率非常高，并且Go也能够很简洁地协同多个并发操作。

**goroutine** 通常直接音译为“**协程**”  (routine 例程 coroutine 协程）

synchronization 同步 asynchronization 异步

## 启动goroutine
在执行的操作之前加上一个`go`关键字即可，就是这么简单。看一个简单直接的例子。
```go
import (
	"fmt"
	"time"
)

func sleepyGopher() {
	time.Sleep(time.Second * 3)	//睡3秒，模拟处理任务
	fmt.Println("...snore...")
}

func main() {
	//通过关键字go，启动goroutine
	go sleepyGopher()
	fmt.Println("this is main func")
	time.Sleep(time.Second * 4)
}
```
执行结果是，控制台会输出`.this is main func`，接着3秒之后，控制台会输出`...snore...`。但是注意，因为在main函数返回时，该程序运行的所有`goroutine`都会被回收，这就是为什么例子中的main函数需要一个比`sleepyGopher`函数长的等待时间。当然处理这个问题有很多方法，比如使用sync包中的`sync.WaitGroup`。

## 启动多个goroutine
每次使用`go`关键字都会创建一个新的`goroutine`。
```go
import (
	"fmt"
	"time"
)

func sleepyGopher() {
	time.Sleep(time.Second * 3)
	fmt.Println("...snore...")
}

func main() {
	// 循环启动5个goroutine
	for i := 0; i < 5; i++ {
		go sleepyGopher()
	}
	time.Sleep(time.Second * 4)
}
```
带参数的函数，一样可以简单的使用`go`关键字启动`goroutine`。为了标记每个`goroutine`，接下来为函数传入一个id。
```go
func sleepyGopher(id int) {
	time.Sleep(time.Second * 3)
	fmt.Println("...snore...", id)
}

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(time.Second * 4)
}
```
输出：
```
...snore... 0
...snore... 3
...snore... 2
...snore... 1
...snore... 4

```
**其实每次的输出都可能不一样，可以看出，goroutine的执行顺序不是我们可以控制的。**   
看完上面的例子后，抛出两个问题

1. main函数不得不`Sleep4`一定的时间来确保所有的`goroutine`全部执行完毕。那么如果`goroutine`中执行的不是上面的这种可知具体耗时的操作（比如数据库操作，网络访问等），那么**如何确定goroutine什么时候结束呢**。
2. **不同的`goroutine`之间如何传递数据**   
Go中提供的**通道**即可解决这两个问题。

## 通道
* 通道（channel）可以在**多个`goroutine`之间安全地传递值**。可以**类比**想象**成**我们平时用的**消息队列**，可以向通道中写入值，可以从通道中取出值。   
* 跟Go中的其他类型一样，可以将通道作为变量，传递至函数，结构中的字段。
* 创建通道的方法：使用内置的`make`函数。并且还要指定相应的类型。
```go
// 无缓冲的整型通道
unbuffered := make(chan int)

// 有缓冲的字符串通道，缓冲容量为10
buffered := make(chan string, 10)
```
`unbuffered`这个通道只能传递`int`类型，同理`buffered`通道只能传递`string`类型。  
对通道中读写值使用左箭头操作符`<-`:

1. 向通道中写入值：`unbuffered <- 7`
2. 从通道中读取值：`i := <- unbuffered`

另外，通道是否带有缓冲，其行为会有一些不同。

### 无缓冲通道
无缓冲的通道（ unbuffered channel） 是指在接收前没有能力保存任何值的通道。**不论是向通道中写入值或者读取值，都会形成阻塞。**比如发送操作会等待直到有另一个`goroutine`尝试对相同的通道执行读取操作为止。相同的，读取操作会等待直到有另一个`goroutine`尝试对相同的通道执行写入操作。
```go
import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("...snore...", id)
	c <- id
}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher", gopherID, "has finished sleeping")
	}
}
```
上面的5个goroutine都向通道`c`中写入了ID值，`main`函数会等待到5个goroutine全部执行结束，即是向通道中写入值之后，才会返回，这样我们就可以不需再像之前一样让main函数也Sleep一段时间来确保goroutine的执行了。

### 有缓冲的通道
有缓冲的通道（ buffered channel）是一种在被接收前能存储一个或者多个值的通道。这种类型的通道并不强制要求`goroutine`之间必须同时完成发送和接收。只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，发送动作才会阻塞。   
有缓冲的通道和无缓冲的通道之间的一个很大的不同：无缓冲的通道保证进行发送和接收的`goroutine`会在同一时间进行数据交换。但有缓冲的通道没有这种保证。

## 使用select处理多个通道
然而在很多时候，我们不希望程序一直阻塞在等待通道处。根据以往的经验，我们可以想到的是为这些等待设置超时时间。   
Go提供了`time.After`函数来设置超时时间，`time.After`函数会返回一个**通道**，这个通道会在特定的时间后接收到一个值（由Go运行时发送）。如果我们不想程序一直等待所有的`goroutine`完成而设置一个超时时间，一个思路是让程序同时等待由`time.After`函数返回的计时通道和其他通道，如果计时通道的值返回了就不再去等待其他通道了。
**为了实现这个功能，Go提供的`select`语句很方便，其语法与`switch`很相似，某个case的准备就绪就会执行相应的操作。这样一来，就可以同时监视多个通道了。**

```go
timeout := time.After(2 * time.Second)
for i := 0; i < 5; i++ {
	select {
	case gopherID := <-c:
		fmt.Println("gopher", gopherID, "has finished sleeping")
	case <-timeout: //等待直到超时
		fmt.Println("my patience ran out")
		return
	}
}
```
**即使程序已经停止等待`goroutine`，但只要`main`函数没有返回，仍在运行的goroutine就仍在占用内存。所以我们应该主动去关闭无用的`goroutine`。可以使用通道的`close()`方法来关闭通道。   
【注意】 如果`select`语句不包含任何分支，将会永远等待下去。** 

#### nil通道
如果不使用`make`函数初始化通道变量，那么和映射和切片一样，通道的值是`nil`。对其进行发送和接收操作不会引起`panic`，但是会导致永久阻塞。但是如果对nil通道执行`close()`方法则会引发`panic`。
* nil通道的应用场景   
`nil`通道的存在并不是一无是处的，比如，一个包含`select`语句的循环，**如果我们不希望循环每次循环都等待所有`case`，那么可以先将某些通道设置为`nil`，等到待发送的值准备好之后再将这些`nil`通道初始化，再去执行发送操作。**

#### 阻塞和死锁
`goroutine`在对通道进行等待和发送时会引起阻塞，等待时程序会一直监视通道啥时候来值。但是这种阻塞和那些空转死循环不一样，除了`goroutine`本身所占的少量资源外，`goroutine`并不消耗任何其他资源。当一个或多个`goroutine`因为某些永远无法发生的事情而被阻塞时，这就发生了**死锁**。

```go
func main(){
	c := make(chan, int)
	<- c
}
```

## 实践一下：流水线
接下来做一个实例综合运用本节学的几个语法。用3个`goroutine`来形成一个流水线作业。三个`goroutine`分为上游，中游和下游，上游产生几个`string`文本，传递到中游进行单词过滤。下游在获取中游过滤后的文本。
```go
//上游
func sourceGopher(downStream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downStream <- v
	}
	downStream <- ""
}

//中游
func filterGopher(upStream, downStream chan string) {
	for {
		item := <-upStream
		if item == "" {
			downStream <- ""
			return
		}
		if !strings.Contains(item, "bad") {
			downStream <- item
		}
	}
}

//下游
func printGopher(upStream chan string) {
	for {
		v := <-upStream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}

//执行
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceGopher(c1)
	go filterGopher(c1, c2)
	printGopher(c2)
}
```
上面的示例中我们是使用空字符串作为一个结束标志，但是这不是很稳定的做法。如果上游的字符串数组中包含一个空字符串，那么流程会被意外关闭。实际上更好的做法是**使用`close`函数来关闭通道**。
```go
close(c1)
```
**如果向已关闭的通道执行写入会引发`panic`，读取已关闭的通道会得到相应的零值。**   
**【注意】** 如果在循环中读取一个已关闭的通道，并且没有检查该通道是否已关闭，那么这个循环会一直运转下去并消耗大量的性能。所以务必对可能关闭的通道检查是否关闭。   
**检查通道是否关闭的写法：**

```go
v, ok := <-c
```
如果`ok`是`false`，那么说明通道`c`已关闭。   
那么上游和中游的代码可以做如下优化
```go
//上游
func sourceGopher(downStream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downStream <- v
	}
	//downStream <- ""
	close(downStream)
}

//中游
func filterGopher(upStream, downStream chan string) {
	for {
		//item := <-upStream
		item, ok := <-upStream
		if !ok {
			//downStream <- ""
			close(downStream)
			return
		}
		if !strings.Contains(item, "bad") {
			downStream <- item
		}
	}
}
```
另外，由于**“从通道中读取值，直到通道被关闭为止”这个操作很常用，所以Go提供了快捷的方法。使用`range`来读取通道，程序会在通道关闭前一直去获取通道的值。**   
这样一来，中游的代码可以再次优化。

```go
func filterGopher(upStream, downStream chan string) {
	//使用range来读取通道的值
	for item := range upStream {
		if !strings.Contains(item, "bad") {
			downStream <- item
		}
	}
	close(downStream)
}
```
下游的代码同理
```go
func printGopher(upStream chan string) {
	//使用range来读取通道的值
	for v := range upStream {
		fmt.Println(v)
	}
}
```

## 总结
goroutine是Go语言中十分重要的知识点，涉及的内容不少，做个总结
1. 使用go关键字启动goroutine
2. 通道用于多个goroutine之间传值
3. 使用make函数创建通道
4. 使用`<-`读写通道
5. 使用select语句同时等待多个通道
6. 使用time.After函数实现超时
7. **使用close函数关闭通道**
8. **使用range语句读取通道的值**


# 并发状态

在使用goroutine后不得不面对一个问题，就是当多个`goroutine`同时去操作同一个共享值就会发生并发问题，这种情况被称作竞争状态（race condition）。在其他语言中我们往往会通过上锁来解决这类问题，那么在Golang中该如何解决。

## 互斥锁
Go中提供了互斥锁`Mutex`(mutual exclusive 相互排斥)，存在于包sync中。从名字大概就可以理解其意思。   
互斥锁中有`Lock`和`Unlock`两个方法。`Lock`就是上锁，`Unlock`就是解锁。如果有`goroutine`尝试在互斥锁已经锁定的情况下再次调用`Lock`方法，那么它将等待直到解除锁定之后才能再次上锁。为了防止锁来锁去引发不可预知的错误，通常情况下只会在包的内部使用。

```go
import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	mu.Lock()
	//通常为了防止有多个return方法而忘记解锁，Unlock通常都用defer来写
	defer mu.Unlock()  
}
```
将`sync.Mutex`用作结构成员的做法是一种常见的模式。下面这个示例，现在有多个`goroutine`来执行网络爬虫，需要有个结构来存储所有被爬的网页的次数。如果使用映射来存储，在多个`goroutine`尝试更新映射时，就会产生竟态条件。那么这时就需要一个互斥锁来保护。
```go
//Visited 用于记录网页是否被访问过
type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

//VisitLink 记录本次针对网址url的访问，更新对url的访问总次数
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
    // 	count := v.visited[url]
	// 	count++
	// 	v.visited[url] = count
	// 	return count
	v.visited[url]++
	return v.visited[url]
}
```
* 使用互斥锁时要小心陷入**死锁**。

## 原子函数
原子函数能够以很底层的加锁机制来同步访问整型变量和指针。`atmoic`包中两个有用的原子函数是 `LoadInt64` 和 `StoreInt64`。这两个函数提供了一种**安全地读和写一个整型值**的方式。常用的使用的方式如下示例。
```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown 是通知正在执行的 goroutine 停止工作的标志
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 创建两个 goroutine
	go doWork("A")
	go doWork("B")

	// 给上面两个goroutine一些执行时间
	time.Sleep(1 * time.Second)

	// 该停止工作了，安全地设置 shutdown 标志
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

// doWork 用来模拟执行工作的 goroutine，检测之前的 shutdown 标志来决定是否提前终止
func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		//检测是否收到终止信号，如果终止则退出，否则继续下一次循环
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
```
atmoic包的`AddInt64`函数。这个函数会同步整型值的加法，方法是强制同一时刻只能有一个**goroutine**运行并完成这个加法操作，例`atomic.AddInt64(&counter, 1)`，类似于C#中的原子累加器`Interlocked.Increment(ref ActivityCount);`。


## 长时间运行的工作进程
我们将一直存在并且独立运行的goroutine称为“工作进程”（worker）。比如一些定时执行某些功能的工作进程，如网站的轮询器等。在C#中我们可以使用一个定时器来完成这样的需求，那么在Golang中我们该如何搭建一个较为通用的工作进程。
```go
//长时间运行的工作进程
func worker() {
	n := 0
	next := time.After(time.Second) //创建一个计时器通道
	for {
		select {
		case <-next: //等待计时器触发
			n++
			fmt.Println(n)
			next = time.After(time.Second) //为下一次循环创建新的计时器
		}
	}
}
```
**其实上面的示例完全可以不用`select`和`time.After`，直接用一个含有`time.Sleep`的循环就可以实现。这里这样写主要为了方便将这个示例拓展成等待多个通道的工作进程。**

### 综合示例
背景描述：现在有一个在火星表面行走的探测器，通过遥控发送命令可控制探测器的行走。探测器有一个工作进程来接受命令进行移动，并且定时刷新探测器的位置。   
首先我们将上面的工作进程进行改写。由于需要记录位置，image包中的`Point`结构很适合，它可以存储x轴和y轴的坐标，并且有一个Add方法可以将一个坐标点与另一个坐标点相加。
```go
func worker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current position is ", pos)
			next = time.After(time.Second)
		}
	}
}
```
上面这个探测器中的工作进程功能还比较单一，只会让探测器直线前进。为了让它能够听从遥控器发送的指令进行转方向、停止、或者加速等，需要添加一个命令通道来发送命令。当工作进程从命令通道中接收到命令后便会立即执行相应的命令。
```go
//命令类型
type command int

const (
	right = command(0) //简单的代表向右转
	left  = command(1) //简单的代表向左转
)

//RoverDriver 用来控制探测器
type RoverDriver struct {
	commandc chan command
}

//NewRoverDriver 创建通道并启动工作进程
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("new direction %v", direction)

		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

//Left 会将探测器转向左方
func (r *RoverDriver) Left() {
	r.commandc <- left
}

//Right 会将探测器转向右方
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(2 * time.Second)
	r.Right()
	time.Sleep(4 * time.Second)
}
```

