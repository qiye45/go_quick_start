package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// QueryFrameworkStats 根据github仓库统计信息接口查询某个仓库信息
func QueryFrameworkStats(ctx context.Context, framework string) <-chan string {
	stats := make(chan string)
	go func() {
		repos := "https://api.github.com/repos/" + framework
		req, err := http.NewRequest("GET", repos, nil)
		if err != nil {
			return
		}
		req = req.WithContext(ctx)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		stats <- string(data)
	}()

	return stats
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	framework := "gin-gonic/gin"
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case statsInfo := <-QueryFrameworkStats(ctx, framework):
		fmt.Println(framework, " fork and start info : ", statsInfo)
	}
}
