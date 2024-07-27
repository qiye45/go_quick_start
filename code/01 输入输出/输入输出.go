package main

import (
	"fmt"
	"math/rand"
)

const (
	maxAttempts = 5   // 最大尝试次数
	maxNumber   = 100 // 可猜测数字的最大值
)

func main() {

	var trueNum = rand.Intn(maxNumber) // 使用改进后的随机数生成，注意Intn的使用

	fmt.Println("欢迎来到数字猜谜游戏！请猜测一个0到", maxNumber-1, "之间的数字。")

	var guess int
	for attempts := 0; attempts < maxAttempts; attempts++ {
		var guessString int
		fmt.Print("请输入你的猜测: ")
		_, _ = fmt.Scanln(&guessString)

		if guess < 0 || guess >= maxNumber {
			fmt.Printf("猜测范围错误。请猜一个0到%d之间的数字。\n", maxNumber-1)
			continue
		}

		if guess == trueNum {
			fmt.Printf("恭喜你，猜对了！答案就是%d。\n", trueNum)
			return
		}

		if guess > trueNum {
			fmt.Println("猜大了")
		} else {
			fmt.Println("猜小了")
		}
	}

	fmt.Printf("很遗憾，你没有在限定的%d次尝试内猜对。答案是%d。\n", maxAttempts, trueNum)
}
