package dns1

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

// dscacheutil -q host -a name localhost

// 在macOS上使用tcpdump抓包，以下是常用方法：

// 基本DNS查询抓包：
// sudo tcpdump -i any port 53
// 更详细的DNS包（包含完整内容）：
// sudo tcpdump -i any -vvv -s0 port 53
// 保存抓包结果到文件（可以用Wireshark打开分析）：
// sudo tcpdump -i any -w dns_capture.pcap port 53
// 针对性抓包localhost相关：
// sudo tcpdump -i lo0 -n port 6379
// 实用参数说明：

// -i any: 监听所有网络接口
// -i lo0: 只监听本地回环接口
// -n: 不解析主机名和端口号
// -vvv: 最详细的输出
// -s0: 抓取完整的包
// -w: 写入文件
// port 53: 只抓DNS包
// host localhost: 只抓与localhost相关的包

// networksetup -listallnetworkservices
// networksetup -getinfo "Wi-Fi"

func TestFoo(t *testing.T) {
	// resolver := net.DefaultResolver
	resolver := net.Resolver{
		// PreferGo: true,
		// PreferGo: false,
	}

	// spew.Dump("resolver", resolver)

	t1 := time.Now()

	ctx := context.Background()
	re, err := resolver.LookupHost(ctx, "localhost")
	boom(err)

	fmt.Println("LookupHost:", re)
	fmt.Println("spent:", time.Since(t1))

	// re2, err := net.LookupNS("localhost")
	// re2, err := resolver.LookupNS(context.Background(), "localhost")
	// boom(err)

	// spew.Dump("LookupNS:", re2)
}

func boom(err error) {
	if err != nil {
		panic(err)
	}
}
