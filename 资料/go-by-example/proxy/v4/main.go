package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

// SOCKS5协议相关常量
const (
	socks5Ver = 0x05 // SOCKS5版本号
	cmdBind   = 0x01 // CONNECT请求
	atypeIPV4 = 0x01 // IPv4类型
	atypeHOST = 0x03 // 域名类型
	atypeIPV6 = 0x04 // IPv6类型
)

func main() {
	// 创建TCP服务器监听1080端口
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}

	// 循环接受客户端连接
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("接受连接失败: %v", err)
			continue
		}
		// 为每个客户端创建一个goroutine处理请求
		go process(client)
	}
}

// 处理客户端连接
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// 进行SOCKS5认证
	if err := auth(reader, conn); err != nil {
		log.Printf("客户端 %v 认证失败:%v", conn.RemoteAddr(), err)
		return
	}

	// 处理CONNECT请求
	if err := connect(reader, conn); err != nil {
		log.Printf("客户端 %v 连接失败:%v", conn.RemoteAddr(), err)
		return
	}
}

// SOCKS5认证处理
func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	// 读取版本号
	ver, err := reader.ReadByte() // 5
	if err != nil {
		return fmt.Errorf("读取版本号失败:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("不支持的版本号:%v", ver)
	}

	// 读取认证方法数量
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("读取认证方法数量失败:%w", err)
	}

	// 读取认证方法列表
	method := make([]byte, methodSize)
	if _, err = io.ReadFull(reader, method); err != nil {
		return fmt.Errorf("读取认证方法失败:%w", err)
	}

	// 返回选择的认证方法（这里选择无需认证）
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("发送认证响应失败:%w", err)
	}
	return nil
}

// 处理CONNECT请求
func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 版本号，socks5的值为0x05
	// CMD 0x01表示CONNECT请求
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名
	// DST.ADDR 一个可变长度的值
	// DST.PORT 目标端口，固定2个字节

	// 读取请求头
	buf := make([]byte, 4)
	if _, err = io.ReadFull(reader, buf); err != nil {
		return fmt.Errorf("读取请求头失败:%w", err)
	}

	ver, cmd, atyp := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("不支持的版本号:%v", ver)
	}
	if cmd != cmdBind {
		return fmt.Errorf("不支持的命令:%v", cmd)
	}

	// 解析目标地址
	var addr string
	switch atyp {
	case atypeIPV4:
		// 处理IPv4地址
		if _, err = io.ReadFull(reader, buf); err != nil {
			return fmt.Errorf("读取IPv4地址失败:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])

	case atypeHOST:
		// 处理域名地址
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("读取域名长度失败:%w", err)
		}
		host := make([]byte, hostSize)
		if _, err = io.ReadFull(reader, host); err != nil {
			return fmt.Errorf("读取域名失败:%w", err)
		}
		addr = string(host)

	case atypeIPV6:
		return errors.New("暂不支持IPv6")

	default:
		return errors.New("无效的地址类型")
	}

	// 读取端口号
	if _, err = io.ReadFull(reader, buf[:2]); err != nil {
		return fmt.Errorf("读取端口失败:%w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2])

	// 连接目标服务器
	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		return fmt.Errorf("连接目标服务器失败:%w", err)
	}
	defer dest.Close()
	log.Println("连接到", addr, port)

	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// ATYPE 地址类型
	// BND.ADDR 服务绑定的地址
	// BND.PORT 服务绑定的端口DST.PORT
	// 发送连接成功响应

	if _, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0}); err != nil {
		return fmt.Errorf("发送响应失败: %w", err)
	}

	// 使用context控制goroutine生命周期
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 双向转发数据
	go func() {
		_, _ = io.Copy(dest, reader)
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, dest)
		cancel()
	}()

	<-ctx.Done()
	return nil
}
