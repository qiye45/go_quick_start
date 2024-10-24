package main

import (
	"fmt"
)

// 深拷贝
func pass_array_cut(arrays [8]string) {
	for i, array := range arrays {
		arrays[i] = array[:len(array)/2]
	}
}

// 深拷贝切片，返回一个新的切片
func pass_array_cut2(arrays []string) []string {
	newArray := make([]string, len(arrays)) // 创建新的切片
	for i, array := range arrays {
		newArray[i] = array[:len(array)/2] // 修改副本
	}
	return newArray // 返回新的切片
}
func pass_array(arrays []string) {
	for i, array := range arrays {
		arrays[i] = array + array
	}
}

//浅拷贝：pass_array 函数接收的是切片 []string，对它进行修改会直接影响原数组，因为切片是对数组的引用。
//深拷贝：pass_array_cut 函数接收一个数组 [8]string，并返回一个新的数组。这相当于对数组的每个元素进行了副本操作，而不修改原始数组。这里不需要明确声明数组长度，因为我们操作的是数组本身。

// 创建一个新数组，长度为原数组的两倍
func arrayTimesTwo(arr [8]string) [16]string {
	var doubledArr [16]string // 新数组，长度为 2 倍
	for i, v := range arr {
		doubledArr[i*2] = v   // 第一个位置放原始值
		doubledArr[i*2+1] = v // 第二个位置放重复值
	}
	return doubledArr
}

func main() {
	fmt.Println("lesson10 Array")

	//Array of planets
	var planets [8]string
	planets[0] = "Mercury" //Assigns a planet at index 0
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2] //retrieves the planet at index 2
	fmt.Println(earth)  //Earth

	//获取数组的长度，通过内建函数len()即可
	fmt.Println(len(planets)) //8

	//复合字面值 初始化数组
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	//试试其他奇怪的写法
	dwarfs2 := [4]string{"a", "b", "c"} //正常初始化
	fmt.Println(dwarfs2[3])             //输出空
	//或者
	planetCollection := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(len(planetCollection)) //8
	fmt.Println(planetCollection)      //[Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
	// 将数组转换为切片
	//数组转切片：planetCollection[:] 表示将数组 planetCollection 转换为一个切片。在 Go 中，[:] 语法将数组或切片的全部元素切片化。
	pass_array(planetCollection[:])
	fmt.Println("浅拷贝", planetCollection)
	pass_array_cut(planetCollection)
	fmt.Println("深拷贝", planetCollection)
	pass_array_cut2(planetCollection[:])
	fmt.Println("深拷贝", planetCollection)
	//使用索引来为数组赋值
	array := [5]int{1: 10, 3: 30}
	fmt.Printf("%+v\n", array) //[0 10 0 30 0]

	//for循环遍历数组
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Printf(dwarf)
	}

	//range 遍历数组
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

	planetsMarkII := planets //会将planets的完整副本赋值给planetMarkII
	planets[2] = "whoops"
	//改变了原数组planets，数组planetsMarkII是副本，所以不受影响
	fmt.Println(planets[2])       //whoops
	fmt.Println(planetsMarkII[2]) //Earth

	//二维数组
	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"
	for column := range board[1] {
		board[1][column] = "p"
	}
	fmt.Println(board) //[[r       r] [p p p p p p p p] [       ] [       ] [       ] [       ] [       ] [       ]]
}
