package main

import (
	"context"
	"fmt"
)

func main() {
	// 创建一个根 Context
	ctx := context.Background()

	// 使用 WithValue 传递键值对
	// 注意：最好用自定义类型作为 key，避免冲突
	type ctxKey string
	const userIDKey ctxKey = "userID"

	// 给 Context 附加一个 userID
	ctxWithValue := context.WithValue(ctx, userIDKey, 42)

	// 传递给函数
	processRequest(ctxWithValue, userIDKey)
}

func processRequest(ctx context.Context, key interface{}) {
	// 从 Context 中取值
	if v := ctx.Value(key); v != nil {
		fmt.Printf("Got value from context: %v\n", v)
	} else {
		fmt.Println("No value found in context")
	}
}
