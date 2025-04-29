package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 定义命令行参数
	cmd := flag.String("C", "", "要执行的命令")
	cmdAlias := flag.String("cmd", "", "要执行的命令 (别名)")
	timeStr := flag.String("T", "", "执行时间 (格式: HH:MM)")
	timeAlias := flag.String("time", "", "执行时间 (别名)")
	dateStr := flag.String("D", "", "执行日期 (格式: MM-DD)")
	dateAlias := flag.String("date", "", "执行日期 (别名)")
	after := flag.Int("A", 0, "N秒后执行")
	afterAlias := flag.Int("after", 0, "N秒后执行 (别名)")
	frequency := flag.Int("frequency", 0, "每隔多少秒执行一次")
	help := flag.Bool("h", false, "显示帮助信息")
	helpAlias := flag.Bool("help", false, "显示帮助信息 (别名)")

	flag.Parse()

	// 帮助信息
	if *help || *helpAlias {
		flag.Usage()
		return
	}

	// 解析命令
	command := *cmd
	if command == "" {
		command = *cmdAlias
	}
	if command == "" {
		fmt.Println("错误: 必须指定要执行的命令 (-C/--cmd)")
		os.Exit(1)
	}

	// 解析时间
	execTime := time.Now()
	if *timeStr == "" {
		*timeStr = *timeAlias
	}
	if *dateStr == "" {
		*dateStr = *dateAlias
	}
	if *after == 0 {
		*after = *afterAlias
	}

	if *timeStr != "" && *after > 0 {
		fmt.Println("错误: -T/--time 和 -A/--after 参数不能同时使用")
		os.Exit(1)
	}

	if *after > 0 {
		execTime = execTime.Add(time.Duration(*after) * time.Second)
	} else {
		if *timeStr == "" {
			*timeStr = "18:00"
		}
		hourMin := strings.Split(*timeStr, ":")
		if len(hourMin) != 2 {
			fmt.Println("错误: 时间格式错误，应为 HH:MM")
			os.Exit(1)
		}
		hour, err1 := strconv.Atoi(hourMin[0])
		min, err2 := strconv.Atoi(hourMin[1])
		if err1 != nil || err2 != nil {
			fmt.Println("错误: 时间格式错误，应为 HH:MM")
			os.Exit(1)
		}
		execTime = time.Date(execTime.Year(), execTime.Month(), execTime.Day(), hour, min, 0, 0, execTime.Location())
	}

	if *dateStr != "" {
		monthDay := strings.Split(*dateStr, "-")
		if len(monthDay) != 2 {
			fmt.Println("错误: 日期格式错误，应为 MM-DD")
			os.Exit(1)
		}
		month, err1 := strconv.Atoi(monthDay[0])
		day, err2 := strconv.Atoi(monthDay[1])
		if err1 != nil || err2 != nil {
			fmt.Println("错误: 日期格式错误，应为 MM-DD")
			os.Exit(1)
		}
		execTime = time.Date(execTime.Year(), time.Month(month), day, execTime.Hour(), execTime.Minute(), 0, 0, execTime.Location())
	}

	// 等待执行时间
	fmt.Printf("任务将在 %s 执行: %s\n", execTime.Format("2006-01-02 15:04:05"), command)
	time.Sleep(time.Until(execTime))

	// 执行命令
	for {
		fmt.Printf("正在执行命令: %s\n", command)
		cmd := exec.Command("cmd", "/c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("命令执行失败: %v\n", err)
		}

		if *frequency <= 0 {
			break
		}
		time.Sleep(time.Duration(*frequency) * time.Second)
	}
}
