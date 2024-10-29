# NIL

可以将其与C#中的`null`类比。在Go中如果一个指针没有明确的指向，那么它的值就是nil。
```go
var i int
var s string
var p *string

fmt.Printf("%v\n", i) //0
fmt.Printf("%v\n", s) // (空字符串)
fmt.Printf("%v\n", p) // <nil>
```

## nil可能会引发Panic
如果对一个nil指针进行解引用会引发panic（引发Go程序崩溃的错误）。
```go
var p *string
fmt.Printf("%v\n", p) // <nil>
fmt.Printf("%v\n", *p) //panic: runtime error: invalid memory address or nil pointer dereference
```
避免这种情况的方法可以在解引用之前先判断指针是否是nil
```go
var nowhere *int

if nowhere != nil {
    fmt.Println(nowhere)
}
```
**以往的编程经验告诉我们，在方法中如果入参或者接收者是指针类型，那么最好都要进行下空判断来确保安全。**
```go
func (p *person) birthday{
    if p == nil{
        return
    }
    p.age++
}
```

## 默认值是nil的情况

### 函数类型
当变量被声明为**函数类型**，在没有被赋值的情况下，其就为nil值。
```go
var fn func(a, b int) int
fmt.Println(fn == nil) //true
```

### 切片
同理，切片在声明之后没有使用复合字面量或者make函数赋值，其值便为nil。
```go
var soup []string
fmt.Println(soup == nil) //true
```
但是一些内置函数和关键字都可以很好的解决nil切片的问题，比如`len`,`append`,`cap`和`range`。
```go
//range可以处理nil
for _, ingredient := range soup {
    fmt.Println(ingredient)
}
//len、append也可以处理nil
fmt.Println(len(soup)) //0
soup = append(soup, "onion", "carrot")
fmt.Println(soup) //[onion carrot]
```
### 映射
同理，映射在声明之后没有使用复合字面量或者make函数赋值，其值便为nil。对nil映射的读取操作不会引发panic，但是**写入操作则会引发panic**。
```go
var souplist map[string]int
fmt.Println(souplist == nil) //true

measurement, ok := souplist["onion"]

if ok {
    fmt.Println(measurement)
}

for ingredient, measurement := range souplist {
    fmt.Println(ingredient, measurement)
}

//souplist["onion"] = 1 //panic: assignment to entry in nil map
```
解释

```
var souplist map[string]int 只是声明了一个 map 变量，但没有实际初始化它。此时 souplist 的值为 nil。向 nil map 写入数据会导致 panic。

要正确使用 map，必须先初始化。有以下几种正确的初始化方式:
// 方式1: 使用 make 函数
souplist = make(map[string]int)

// 方式2: 使用 map 字面量
souplist := map[string]int{}

// 方式3: 初始化时直接指定元素
souplist := map[string]int{
    "onion": 1,
}
```



### 接口

接口类型的变量在未被赋值时的零值是nil，并且它的接口类型和值都是nil。
```go
var v interface{}
fmt.Printf("%T %v %v\n", v, v, v == nil) //<nil> <nil> true
```
值得注意的是，当接口类型的变量被赋值之后，接口就会在内部指向该变量的类型和值。先看下面的示例。
```go
var v interface{}

var po *int
v = po
fmt.Printf("%T %v %v\n", v, v, v == nil) //*int <nil> false
```
**在将`po`赋值给`v`之后，`v`的类型就变成了`*int`，虽然值仍然是`nil`，但是Go认定接口类型的变量只有在类型和值都为nil时才等于`nil`。所以`v == nil`的结果是`false`。**

```go
//格式化变量 %#v 可以同时打印出变量的类型和值
fmt.Printf("%#v", v)    //(*int)(nil)
```

# 错误处理

## 处理错误
在Go语言中，`error`类型是专门为错误而设的一种内置类型，有点类型C#中的`Exception`类型（但是`error`不捕获也不会使程序崩溃）。由于Go允许函数有多个返回值，所以在Go语言中一种较为常见的写法来传递发生的错误的信息，就是**将错误信息写在返回值（一般为最后一个返回值）**。举个栗子
```go
//第二参数为错误(error)类型
files, err := os.ReadDir(".")

//如果其不为空，则是发生了异常
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}

for _, file := range files {
    fmt.Println(file.Name())
}
```

## 优雅的错误处理
通过之前的说法，`error`都会在返回值中返回。接下来有一个需求是写一个函数，函数中创建一个文件并向其中写入文本。根据以往经验告诉我们，在执行这些功能时随时都有可能发生异常，文件创建时名称不合法，权限不足，目录不存在等等，在写入文本时也会遇到各种异常，这就需要我们针对所有可能发生的异常进行相应的处理。这个函数一般的写法可以这样写。
```go
func proverbs(name string) error {
	//创建文件
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	//写文本信息
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close()
		return err
	}
	//写文本信息
	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	if err != nil {
		f.Close()
		return err
	}
	//写文本信息
	_, err = fmt.Fprintln(f, "Don't Panic.")
	f.Close()
	return err
}
```
功能可以正常实现，但是容易发现这其中存在两个明显的问题。
1. **在每次出现错误后都需要显式的调用`f.Close()`**
2. **每次写一行文本信息都要检测异常，语法显得很臃肿**   

### 关键字defer
为了保证文件能够正确被关闭（`f.Close()`），可以使用`defer`关键字。`defer`是延迟的意思，`defer`关键字的功能就是延迟执行被它标记的操作。**被`defer`标记的操作，Go语言会在函数返回之前触发**。有点像在C#的`try...cathc...finally`中，我们**将这些操作写在finally块中的道理一样**。   
使用defer关键字之后的代码

```go
func proverbsWithDefer(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	//使用defer关键字，表示在函数退出之前，执行f.Close()
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Don't Panic.")
	return err
}
```

### 错误处理
我们可以声明一个新的类型`safeWriter`,在写入文件的过程中发生了错误，那么它将错误存储起来而不是直接返回它，之后当writerln尝试在此写入相同的文件时，如果发现之前已有错误，那么将不会再执行后续的操作。
```go
type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbsGracefully(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't Panic.")
	return sw.err
}
```
> **这种写法背后的思想比写法本身重要的多。**

```
这个 safeWriter 的设计理念在于延迟错误处理，即不直接处理每个单独的写入操作中出现的错误，而是记录第一个错误并在整个操作完成后再进行统一处理。这种写法有几个重要的思想：

简化错误处理：通常情况下，每次调用一个会返回错误的函数时都要检查并处理错误。这种写法避免了在多次写入操作中反复进行错误检查的繁琐步骤，只需要在所有写入操作完成后检查一次 sw.err 即可。

“错误即值”：在 Go 中，错误被视为普通值，因此可以像处理其他值一样存储和传递。在 safeWriter 中，err 就是一个普通字段。通过检查 sw.err 是否为 nil 来决定是否执行操作，这种方式让代码更加整洁，也符合 Go 的设计哲学。

延迟操作的鲁棒性：在多个操作组成的流程中，只要有一个操作失败，之后的操作通常可能会因前一个失败状态导致未定义的行为。这里的设计确保了一旦某次写入失败，后续操作将被跳过，避免了资源浪费或产生新的错误。

集中处理：所有写入操作完成后统一返回错误，使调用方只需在一个地方进行错误检查，使代码逻辑更加清晰。
```

## 新的错误
在出现错误时，我们可以通过创建并返回新的错误值来通知调用者出现了什么问题。在C#中我们是可以通过继承`Exception`作为基类来创建自定义的异常类型，那在Go中，`error`包包含了一个构造函数，它接受一个代表错误消息的字符串作为参数。通过这个构造函数可以创建并返回自定义的错误。   
接下来用一个数独的例子来举例。

```go
const rows, columns = 9, 9

//模拟一个9*9的数独网格
type Grid [rows][columns]int8

func inBound(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if row < 0 || row >= columns {
		return false
	}
	return true
}

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBound(row, column) {
		return errors.New("out of bound")
	}
	g[row][column] = digit
	return nil
}

func main() {
	var g Grid
	myErr := g.Set(10, 0, 5)
	if myErr != nil {
		fmt.Printf("An error occurred: %v\n", myErr)  //An error occurred: out of bound
		os.Exit(1)
	}
}
```
### 按需返回错误
在Go的很多包中，都会声明并导出一些变量用来表示他们可能会返回的错误。继续延续之前的数独的例子，可以声明两个错误变量。
```go
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)
```
**【注意】** 按照惯例，Go的错误类型都用Err打头。   

声明之后，我们就不用去临时声明`errors.New("out of bounds")`了，直接返回`ErrBound`就可以了。
```go
if !inBound(row, column) {
		//return errors.New("out of bound")
		return ErrBounds
	}
```
返回特定的错误，方法的调用者就可以根据具体的错误类型进行不同的错误处理了。

### 自定义错误类型
虽然`errors.New()`可以创建自定义的错误消息，但是有时候还是不够用。`error`类型是一个内置的接口，无论什么类型，**只要实现了一个返回字符串的Error()方法，就隐式满足了error接口**，这样就可以基于这个接口创建出新的错误类型。
```go
type error interface{
    Error() string
}
```
#### 返回多个错误
当代码执行中遇到多个错误是，比如之前的数独代码，当传入的位置越界了，值又是一个非法值，那么这时候与其每次返回一个错误，不如让方法进行多次检查一次性返回所有错误。
```go
type SudokuError []error

//Error返回一个或多个用逗号分隔的错误
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBound(row, column) {
		//return errors.New("out of bound")
		//return ErrBounds
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
}
```

#### 类型断言
上面的例子中，返回值之前会将值从`SudokuError`类型转为`error`接口类型，如果想单独访问每个错误就必须进行类型转换。



```go
var g Grid
errs := g.Set(10, 0, 15)
if errs != nil {
    if sudokuError, ok := errs.(SudokuError); ok {
        fmt.Printf("%d error(s) occurred:\n", len(sudokuError))
        for _, e := range sudokuError {
            fmt.Printf("- %v\n", e)
        }
    }
    os.Exit(1)
}
```
上面`errs.(SudokuError)`断言`err`的类型为`SudokuError`。

```
类型断言用于检查变量是否为某种类型，并返回基础接口值。类型断言仅适用于接口。
例如，在下面的代码中: 'var x interface{} = 42 t := x.(int)'，'x' 具有带基础 int 值 ('42') 的 'interface{}' 类型，'int' 是我们要检查的具体类型。如果我们打印 't'，输出将为 '42'。将具体类型更改为 'string' ('t := x.(string)') 将导致运行时宕机。类型断言可以返回两个值。例如，表达式 't, ok := x.(int)' 具有布尔值 'ok'，如果断言正确，则返回 'true'。如果 'ok' 为 'false'，则会将 't' 设置为零值，并且不会发生宕机。
空接口 'interface{}' 代表所有类型的集合。空接口类型的变量可以存储任何类型的值。
```

**检查是不是某个类型**

```
isinstance(1,int)
True
isinstance(1.0,int)
False
```



## 不要惊恐（Panic）
Go语言中没有提供异常机制，但是有名为**panic**的类似机制，前面也都有提及。如同C#中的`Exception`出现一样，Go遇到`Panic`后，程序会崩溃。在其他语言中，如果发生异常，没有人捕捉的话这个异常会一层一层的向上抛，一直抛到`main`函数之类的调用栈顶。处理这些异常会用到大量的`try...catch...finally...throw...`等等。相比之下Go语言的错误值提供了一个简单且灵活的机制来替代异常，促使开发者考虑错误，而不是像异常处理那样默认将其忽略，有助于生成更可靠的软件。  

* 如果想要引发恐慌`panic`，可以这样
```go
panic("OMG, i'm sorry")
```
**【注意】** panic在退出前会执行所有`defer`延迟的操作，而`os.Exit(1)`则不会这样，所以panic比`os.Exit(1)`还好点。当然，择情处理。  

* 当然Go也提供了“反悔”的办法，为了防止panic让程序崩溃，可以使用`recover`函数
```go
defer func() {
    if e := recover(); e != nil {
        fmt.Println(e) //OMG, i'm sorry
    }
}()

panic("OMG, i'm sorry")
```