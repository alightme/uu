package main

import (
	"flag"
	"fmt"
	"os"
	"uu/src/ip"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// 定义命令行参数
	ipCmd := flag.NewFlagSet("ip", flag.ExitOnError)
	// 解析命令行参数
	if len(os.Args) < 2 {
		fmt.Println("使用方法: uu <command>")
		fmt.Println("可用命令: ip")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "ip":
		ipCmd.Parse(os.Args[2:])
		if ipCmd.Parsed() {
			ipAddr, err := ip.GetLocalIP()
			if err != nil {
				fmt.Println("错误:", err)
				os.Exit(1)
			}
			fmt.Println("本地IP地址:", ipAddr)
		}
	default:
		fmt.Println("未知命令:", os.Args[1])
		os.Exit(1)
	}
}
