package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// userInfo 定义了用户信息结构体
type userInfo struct {
	Name  string   // 用户名
	Age   int      `json:"age"` // 年龄，使用json标签指定序列化后的字段名为"age"
	Hobby []string // 爱好列表
}

func main() {

	// 创建一个userInfo实例
	a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript", "哈哈哈"}}

	// 将结构体序列化为JSON字节切片
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf) // 打印JSON字节切片
	// 十六进制打印
	fmt.Printf("%x\n", buf)
	// 写buf到文件
	// 获取当前文件所在目录
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("无法获取当前文件路径")
	}
	dir := filepath.Dir(currentFile)
	fmt.Println("源文件所在目录:", dir)
	filePath := filepath.Join(dir, "json.txt")
	fmt.Println("文件将保存在:", filePath)
	// 写入文件
	err = os.WriteFile(filePath, buf, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf)) // 将JSON字节切片转换为字符串并打印

	// 使用MarshalIndent生成格式化的JSON字符串
	// 第二个参数是每行的前缀，第三个参数是缩进字符
	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf)) // 打印格式化后的JSON字符串

	// 定义一个新的userInfo变量用于接收解析结果
	var b userInfo
	// 将JSON数据解析到结构体b中
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	// 使用%#v格式化输出结构体的详细信息
	fmt.Printf("%#v\n", b)
}
