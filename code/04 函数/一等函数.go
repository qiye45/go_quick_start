package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Kelvin float64

// 返回模拟温度的传感器
func fakeSensor() Kelvin {
	return Kelvin(rand.Intn(151) + 150)
}

// 返回真实温度的传感器
func realSensor() Kelvin {
	return 0
}

// 【case2】将函数传递给其他函数
// 测量温度，使用传入的传感器测量samples次温度
func measureTemperature(samples int, sensor func() Kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		//time.Sleep(2 * time.Second)  // 秒（时间单位）
		time.Sleep(time.Second)
	}
}

// 匿名函数
var f = func() {
	fmt.Println("Dress up for the masquerade")
}

// sensor函数类型
type sensor func() Kelvin

// 声明并返回一个匿名函数
func calibrate(s sensor, offset Kelvin) sensor {
	return func() Kelvin {
		return s() + offset
	}
}

func main() {
	fmt.Println("lesson9 First-class functions")

	//【case1】将函数赋值给变量
	sensor1 := fakeSensor
	fmt.Println(sensor1()) //156(随机的)

	sensor2 := realSensor
	fmt.Println(sensor2()) //0

	//fmt.Println(sensor) //报错： 实参 'sensor' 不是函数调用

	measureTemperature(3, sensor1) //0° K (连续打印三次)
	measureTemperature(3, sensor2) //0° K (连续打印三次)

	//调用匿名函数
	f() //Dress up for the masquerade

	//将匿名函数赋值给函数中的变量
	ff := func(message string) {
		fmt.Println(message)
	}
	ff("Go to the party") //Go to the party

	//将匿名函数的声明和执行放在一起写
	func() {
		fmt.Println("function anonymous")
	}() //function anonymous

	newSensor := calibrate(realSensor, 5)
	fmt.Println(newSensor()) //5
}
