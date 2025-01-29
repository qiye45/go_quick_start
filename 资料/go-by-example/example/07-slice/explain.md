在 Go 语言中，除了 append 外，还有一些用于处理切片/数组的常用内置函数和包：

1. 内置函数：
```go
// copy - 复制切片
copy(dst, src []Type) int

// len - 获取长度
len(v Type) int

// cap - 获取容量
cap(v Type) int

// make - 创建切片
make([]Type, length, capacity)

// delete - 删除map中的元素(不是切片)
delete(m map[Type]Type1, key Type)
```

2. sort包提供的函数：
```go
import "sort"

// 排序
sort.Slice(slice interface{}, less func(i, j int) bool)
sort.Strings(slice []string)  // 字符串切片排序
sort.Ints(slice []int)       // 整数切片排序
sort.Float64s(slice []float64) // 浮点数切片排序

// 搜索
sort.SearchInts(slice []int, key int) // 二分查找
```

3. 使用示例：
```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 切片排序
    numbers := []int{3, 1, 4, 1, 5, 9}
    sort.Ints(numbers)
    fmt.Println(numbers) // [1 1 3 4 5 9]

    // 自定义排序
    people := []struct {
        Name string
        Age  int
    }{
        {"Alice", 25},
        {"Bob", 30},
        {"Charlie", 20},
    }
    
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })
    fmt.Println(people) // 按年龄升序排序

    // 合并切片
    slice1 := []int{1, 2, 3}
    slice2 := []int{4, 5, 6}
    merged := append(slice1, slice2...)
    fmt.Println(merged) // [1 2 3 4 5 6]
}
```

4. slices包（Go 1.21新增）：
```go
import "slices"

// 比较切片
slices.Equal(slice1, slice2)

// 克隆切片
slices.Clone(slice)

// 包含元素
slices.Contains(slice, element)

// 插入元素
slices.Insert(slice, index, elements...)

// 删除元素
slices.Delete(slice, start, end)
```

5. 自定义辅助函数：
```go
// 删除切片中的某个元素
func remove(slice []int, i int) []int {
    return append(slice[:i], slice[i+1:]...)
}

// 在指定位置插入元素
func insert(slice []int, index int, value int) []int {
    slice = append(slice, 0)      // 扩展一个空间
    copy(slice[index+1:], slice[index:])
    slice[index] = value
    return slice
}
```

这些函数和包提供了丰富的切片操作功能。选择使用哪个取决于具体需求：
- append 适合添加元素
- copy 适合复制数据
- sort包适合排序和搜索
- slices包提供了更多现代的切片操作
- 对于特殊需求，可以编写自定义的辅助函数