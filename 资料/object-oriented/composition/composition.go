package composition

import (
	"fmt"
)

// Engine 发动机结构体
type Engine struct {
	Power     int
	FuelType  string
	isStarted bool
}

// Start 启动发动机
func (e *Engine) Start() {
	e.isStarted = true
	fmt.Printf("发动机启动，功率: %d hp\n", e.Power)
}

// Stop 停止发动机
func (e *Engine) Stop() {
	e.isStarted = false
	fmt.Println("发动机停止")
}

// IsRunning 检查发动机是否运行中
func (e Engine) IsRunning() bool {
	return e.isStarted
}

// Wheel 轮子结构体
type Wheel struct {
	Diameter int
	Material string
}

// Rotate 旋转轮子
func (w Wheel) Rotate() {
	fmt.Printf("直径为 %d 英寸的%s轮子旋转\n", w.Diameter, w.Material)
}

// Car 汽车结构体，通过嵌入组合Engine和Wheel
type Car struct {
	Engine          // 嵌入Engine，"继承"其字段和方法
	Wheels [4]Wheel // 包含4个Wheel
	Brand  string
	Model  string
}

// Drive 汽车驾驶方法
func (c *Car) Drive() {
	if !c.IsRunning() {
		c.Start() // 可以直接调用Engine的方法
	}
	fmt.Printf("驾驶 %s %s\n", c.Brand, c.Model)
	for i := range c.Wheels {
		c.Wheels[i].Rotate()
	}
}

// ElectricCar 电动车结构体，嵌入Car
type ElectricCar struct {
	Car          // 嵌入Car，"继承"其字段和方法
	BatteryLevel int
}

// Charge 电动车充电方法
func (e *ElectricCar) Charge(amount int) {
	prevLevel := e.BatteryLevel
	e.BatteryLevel += amount
	if e.BatteryLevel > 100 {
		e.BatteryLevel = 100
	}
	fmt.Printf("%s %s 充电: %d%% -> %d%%\n", e.Brand, e.Model, prevLevel, e.BatteryLevel)
}

// DemoComposition 演示Go语言中的组合（"继承"）
func DemoComposition() {
	// 创建一个汽车实例
	car := Car{
		Engine: Engine{
			Power:    150,
			FuelType: "汽油",
		},
		Wheels: [4]Wheel{
			{Diameter: 17, Material: "合金"},
			{Diameter: 17, Material: "合金"},
			{Diameter: 17, Material: "合金"},
			{Diameter: 17, Material: "合金"},
		},
		Brand: "丰田",
		Model: "卡罗拉",
	}

	// 调用嵌入类型的方法
	car.Drive()
	car.Stop()

	// 创建一个电动车实例
	tesla := ElectricCar{
		Car: Car{
			Engine: Engine{
				Power:    300,
				FuelType: "电力",
			},
			Wheels: [4]Wheel{
				{Diameter: 19, Material: "合金"},
				{Diameter: 19, Material: "合金"},
				{Diameter: 19, Material: "合金"},
				{Diameter: 19, Material: "合金"},
			},
			Brand: "特斯拉",
			Model: "Model 3",
		},
		BatteryLevel: 50,
	}

	// 调用嵌入类型的方法以及自己的方法
	tesla.Drive()
	tesla.Charge(30)
	tesla.Stop()
}
