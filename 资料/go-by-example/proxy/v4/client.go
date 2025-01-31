package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	proxyAddr := "127.0.0.1:1080" // SOCKS5 代理服务的地址
	// 想要访问的目标服务器和端口，这个例子中使用HTTP的80端口

	// 创建与 SOCKS5 代理服务器的连接
	conn, err := net.Dial("tcp", proxyAddr)
	if err != nil {
		log.Fatalf("Failed to connect to proxy: %v\n", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// SOCKS5 认证阶段，发送不需要认证的请求
	_, err = conn.Write([]byte{0x05, 0x01, 0x00})
	if err != nil {
		log.Fatalf("Failed to write auth request: %v\n", err)
	}

	// 读取认证回复，期待的是 [0x05, 0x00]
	response := make([]byte, 2)
	_, err = io.ReadFull(reader, response)
	if err != nil {
		log.Fatalf("Failed to read auth response: %v\n", err)
	}
	if response[1] != 0x00 {
		log.Fatalf("SOCKS5 authentication failed: %v\n", response)
	}

	// 连接阶段，告诉代理要连接的目标地址
	_, err = conn.Write([]byte{
		0x05, 0x01, 0x00, 0x03, byte(len("example.com")),
	})
	if err != nil {
		log.Fatalf("Failed to write connection request: %v\n", err)
	}
	_, err = conn.Write([]byte("example.com"))
	if err != nil {
		log.Fatalf("Failed to write target hostname: %v\n", err)
	}
	_, err = conn.Write([]byte{0x00, 0x50}) // 端口号80 (0x0050)
	if err != nil {
		log.Fatalf("Failed to write target port: %v\n", err)
	}

	// 读取连接回复，期待的是成功 [0x05, 0x00]
	response = make([]byte, 10) // 固定前10字节是协议标准返回格式，忽略后面的长度
	_, err = io.ReadFull(reader, response)
	if err != nil {
		log.Fatalf("Failed to read connection response: %v\n", err)
	}
	if response[1] != 0x00 {
		log.Fatalf("Failed to establish connection through proxy: %v\n", response)
	}

	// 发送 HTTP 请求
	request := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v\n", err)
	}

	// 读取并输出响应
	responseReader := bufio.NewReader(conn)
	for {
		line, err := responseReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Failed to read response: %v\n", err)
		}
		fmt.Print(line)
	}
}
