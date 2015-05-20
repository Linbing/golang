package main

import (
	"flag"
	"fmt"
	"strings"
)

func start_flag() {
	//第一个参数是“参数名”，第二个是“默认值”，第三个是“说明”，返回得是指针
	//host := flag.String("host", "192.168.40.1,192.168.40.2", "a host name")
	host := flag.String("host", "192.168.40.1", "a host name")
	port := flag.Int("port", 80, "a port number")
	debug := flag.Bool("d", false, "enable/disable debug mode")
	var Servers []string
	Servers = strings.Split(*host, ",")
	fmt.Println("len of serveres is = %d ....", len(Servers))
	fmt.Println("servers o  = %s ....", Servers[0])
	//fmt.Println("servers o  = %s ....", Servers[1])
	//上面是定义规则，这里开始正式parse命令行参数
	flag.Parse()
	//命令行如下：go run ./src/main.go ./src/flagtest.go -host=www.baidu.com -port=333 -d=true
	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
	fmt.Println("debug:", *debug)
}

func GetInputFile() *string {
	configfile := flag.String("c", "config", "config file")
	flag.Parse()
	if configfile != nil {
		return configfile
	}
	return nil
}
