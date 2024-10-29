# 指针(1)：Pointer

## &和*
* **&：地址操作符，用以得到变量的内存地址**
* ***：解引用，提供内存地址指向的值**
```go
answer := 42
fmt.Println(&answer) //output: 0xc0000a00b0

address := &answer
fmt.Println(*address) //42

fmt.Printf("address is a %T\n", address) //address is a *int
```
**%T**
**英文全称: Type**

上面的address变量实际上是一个`*int`类型的指针，它可以指向类型为`int的其他变量。  

* 指针类型也是一种类型，也可以用在变量声明，函数形参，返回值类型，结构字段类型等。
```go
china := "China"
var home *string
fmt.Printf("home is a %T\n", home) //home is a *string

home = &china
fmt.Println(*home) //China
```
但是如果上面的**`home`变量不可以指向除了`string`类型之外的其他类型，这使得Go相对于C来说更加安全。**
```go
home = &answer //error:cannot use &answer (type *int) as type *string in assignment
```

## 指针的基本操作
看示例代码，一目了然
```go
//声明一个string指针
var administrator *string

//指针指向第一个人
scolese := "Christopher J. Scolese"
administrator = &scolese
fmt.Println(*administrator) //Christopher J. Scolese

//指针指向第二个人
bolden := "Charles F. Bolden"
administrator = &bolden
fmt.Println(*administrator) //Charles F. Bolden

//修改bolden变量，使用指针访问可以看到变量的更改
bolden = "Charles Frank Bolden Jr."
fmt.Println(*administrator) //Charles Frank Bolden Jr.

//也可以通过“解引用”间接改变变量
*administrator = "Maj. Gen. Charles Frank Bolden Jr."
fmt.Println(bolden) //Maj. Gen. Charles Frank Bolden Jr.

//把指针赋值给变量，将会产生一个指向相同变量的指针。
major := administrator
*major = "Maj. General Charles Frank Bolden Jr."
fmt.Println(bolden) //Maj. General Charles Frank Bolden Jr.

fmt.Println(administrator == major) //true

//但是解引用将指针指向的变量赋值给另一个变量将产生一个副本
charles := *major
*major = "Charles Bolden"
fmt.Println(charles) //Maj. General Charles Frank Bolden Jr.
fmt.Println(bolden)  //Charles Bolden

//就算两个string变量指向不同的地址，但是只要他们的字符串值相同，那么判等时就是ture
charles = "Charles Bolden"
fmt.Println(bolden == charles)   //true
fmt.Println(&bolden == &charles) //false
```

### 指向结构的指针
对于指向结构的指针，Go的设计者为其做了优化。比如在复合字面量的前面可以用`&`，但是在访问字段时，前面可以不用加`*`，因为Go会自动实施指针解引用。
```go
type person struct {
	name, superpower string
	age              int
}

timmy := &person{
    name: "Timothy",
    age:  10,
}

timmy.superpower = "fly"
fmt.Printf("%+v\n", timmy) //&{name:Timothy superpower:fly age:10}
```
**注意：** 字符串字面量和整数浮点数字面量之前不能放置`&`。

```
字面量(Literals)的本质：
字面量是直接在代码中写出的常量值
它们不存储在固定的内存地址中
它们是临时的、即时的值

// 以下代码都是错误的
ptr1 := &"hello"    // 错误：不能取字符串字面量的地址
ptr2 := &42         // 错误：不能取整数字面量的地址
ptr3 := &3.14       // 错误：不能取浮点数字面量的地址

// 正确方式：先将字面量赋值给变量，再取地址
str := "hello"
ptr1 := &str        // 正确

num := 42
ptr2 := &num        // 正确

f := 3.14
ptr3 := &f          // 正确
```



### 指向数组的指针
和指向结构的指针类似，对于数组的复合字面量，Go会自动实施指针解引用。
```go
superpowers := &[3]string{"flight", "invisibility", "super-strength"}
fmt.Println(superpowers[1])   //invisibility
fmt.Println(superpowers[1:3]) //[invisibility super-strength]
```

# 指针(2)：Pointer

## 实现修改
**通过指针可以实现跨越函数和方法边界的修改**

### 将指针用作形参
通过前面的学习我们知道函数是以传值的方式传递形参的。
```go
type person struct {
	name string
	age  int
}
//入参是 person 类型
func birthday(p person) {
	p.age++
}

func main() {
	jack := person{
		name: "Jack",
		age:  12,
	}

    //这里会传入一个jack的副本
	birthday(jack)
    //原jack的字段值不会改变
	fmt.Println(jack.age) //12    
}
```
但当**指针被被传递至函数时，函数将接收到传入内存地址的副本**，在此之后函数就可以**通过解引用内存地址来修改指针指向的值**。
```go
//入参改为 person 指正类型
func birthday(p *person) {
	p.age++
}

jack := person {
    name: "Jack",
    age:  12,
}

birthday(&jack)
fmt.Println(jack.age) //13
```
### 指针接收者
与形参的写法类似，**将指针用作方法接收者（Receiver）时，便可以实现对接收者字段的修改**，看示例。

**方法接收者用指针，为这个类型添加方法**

```go
func (p *person) birthday() {
	p.age++
}

func main() {
	jack := &person{
		name: "Jack",
		age:  12,
	}

	jack.birthday()
	fmt.Println(jack.age) //13
}
```
其实，就算在声明struct时不写`&`，仍然可以正常运行。因为Go在变量通过点标记调用方法是会自动使用`&`取得变量的内存地址。
```go
func (p *person) birthday() {
	p.age++
}

func main() {
	tom := person{
		name: "Tom",
		age:  20,
	}
	tom.birthday()
	fmt.Println(tom.age) //21
}
```
可以看到就算不写`(&tom).birthday()`也可以正常运行。   
当然不是所有涉及到struct的方法都要以指针作为参数，需要视情况而定。

### 内部指针
Go提供了叫做**内部指针**的特性，来确定**struct中的指定的字段的内存地址**。
```go
type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name string
	stats
}

func main() {
	yasuo := character{name: "Yasuo"}

	levelUp(&yasuo.stats)
	fmt.Printf("%+v", yasuo) //{name:Yasuo stats:{level:1 endurance:56 health:280}}
}
```
类似于`&yasuo.stats`这样就可以提供指向struct内部的指针。

### 修改数组
虽然前面说Go中更倾向于使用切片而不是数组，但是也难免会遇到使用数组更加合理的情况。同样使用指针也可以实现对数组元素进行修改的方法。
```go
func reset(board *[8][8]rune){
    board[0][0] = 'r'
}

func main(){
    var board [8][8]rune
    reset(&board)
    fmt.Printf("%c", board[0][0])   //r
}
```

## 隐式指针
不是所有修改都需要显式的使用指针，Go有些地方会“暗中”使用指针。
* **映射也是指针**（map）   
前面我们知道，映射的传值和赋值时传递的都不是副本。因为映射实际上就是一种隐式的指针。

* **切片指向数组**   
切片在指向数组元素的时候也是使用了指针。之前提到切片其实是一个结构体类型。
```go
type slice struct {
    array unsafe.Pointer 
    len   int
    cap   int
}
```
切片内部三个字段
1. 指向数组的指针
2. 切片的长度
3. 切片的容量  

当切片被直接传递至函数或者方法的时候，切片的内部指针就可以对底层数组数据进行修改。   
**【注】**：指向切片本身的指针唯一的用处就是修改切片本身，包括长度、容量及起始位置。

## 指针和接口
先看示例A，其实和接口那节用到的示例一样
【示例A】

```go
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

//martain类型实现了接口talker
type martain struct{}

func (m martain) talk() string {
	return "neck neck"
}

func main(){
	
	shout(martain{})  //NECK NECK
	shout(&martain{}) //NECK NECK
}
```
在上面，**无论是传递`martian`变量还是传递指向`martian`变量的指针，都可以满足`talker`接口**。如果方法使用的是指针接收者，那么情况就不同了。
```go
type laser struct{}

func (l *laser) talk() string {
	return "pew pew"
}

func main(){
	//shout(laser{}) //error: laser does not implement talker
	shout(&laser{}) //PEW PEW
}

```
**如果方法使用的是指针接收者，那么只能使用指针来调用该方法。**

## 明智地使用指针
**切记：不要过度使用指针**



其他

````
在 Go 中，以下情况通常建议使用指针（传地址）：

1. **大型结构体**：
```go
type LargeStruct struct {
    Data    [1000]int
    Field1  string
    Field2  float64
    // ... 很多字段
}

func processLargeStruct(ls *LargeStruct) {
    // 使用指针避免复制大量数据
}
```

2. **需要修改的切片底层数组**：
```go
func modifyArray(arr *[]int) {
    // 需要修改切片本身（比如要改变长度或容量）时
    *arr = append(*arr, 100)
}
```

3. **maps（虽然是引用类型，但有些情况仍需要指针）**：
```go
func modifyMap(m *map[string]int) {
    // 如果需要修改 map 本身（比如重新分配）
    *m = make(map[string]int)
}
```

4. **接口实现**：
```go
type Interface interface {
    Method()
}

type MyType struct {
    data string
}

// 使用指针接收者实现接口
func (m *MyType) Method() {
    m.data = "changed"
}
```

5. **同步原语**：
```go
type Counter struct {
    sync.Mutex
    count int
}

func (c *Counter) Increment() {
    c.Lock()
    defer c.Unlock()
    c.count++
}
```

6. **需要修改的通道**：
```go
func modifyChannel(ch *chan int) {
    // 需要修改通道本身时
    *ch = make(chan int, 10)
}
```

7. **需要判断 nil 的情况**：
```go
func processValue(v *string) {
    if v == nil {
        // 处理空值情况
        return
    }
    // 处理非空值
}
```

8. **嵌入式类型需要修改时**：
```go
type Embedded struct {
    Value int
}

type Container struct {
    *Embedded  // 嵌入指针
}
```

9. **性能关键的场景**：
```go
// 频繁调用的方法，避免值复制
func (p *Point) Distance(q *Point) float64 {
    return math.Sqrt(Square(p.X - q.X) + Square(p.Y - q.Y))
}
```

10. **循环中的变量**：
```go
var ptrs []*int
for i := 0; i < 10; i++ {
    num := i  // 创建新变量
    ptrs = append(ptrs, &num)  // 存储指针
}
```

**不推荐使用指针的场景**：

1. **小的基本类型**：
```go
// 不需要用指针
func addOne(x int) int {
    return x + 1
}
```

2. **只读的小结构体**：
```go
type Point struct {
    X, Y float64
}

// 不需要修改时使用值传递
func (p Point) Distance() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
```

3. **切片和映射的一般操作**：
```go
// 切片和映射本身就是引用类型
func processSlice(s []int) {
    s[0] = 100  // 直接修改，不需要指针
}
```

使用指针的关键考虑因素：
1. 是否需要修改值
2. 数据结构的大小
3. 性能要求
4. 并发安全性
5. 接口实现需求
6. nil 检查需求

建议：在不确定是否使用指针时，可以先考虑：
- 数据结构大小
- 是否需要修改
- 性能需求
这三个方面来做决定。
````

