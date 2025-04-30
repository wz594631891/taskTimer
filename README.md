# TaskTimer 任务定时器

## 简介
TaskTimer 是一个用 Go 语言编写的命令行工具，用于在指定时间或延迟后执行命令，还支持定时重复执行。

## 安装
```sh
# 编译项目
go build -o bin/tasktimer.exe src/main.go
```

## 使用方法
```sh
# N 秒后执行命令
./bin/tasktimer -A 60 -C "echo Hello World"

# 在指定时间执行命令
./bin/tasktimer -T 20:00 -C "echo Hello World"

# 在指定日期和时间执行命令
./bin/tasktimer -D 10-01 -T 20:00 -C "echo Hello World"

# 每隔一段时间重复执行命令
./bin/tasktimer -C "echo Hello World" -frequency 60
```

## 命令行参数
- `-C, --cmd`: 要执行的命令
- `-T, --time`: 执行时间 (格式: HH:MM)
- `-D, --date`: 执行日期 (格式: MM-DD)
- `-A, --after`: N 秒后执行
- `-frequency`: 每隔多少秒执行一次
- `-h, --help`: 显示帮助信息

---

# TaskTimer Task Scheduler

## Introduction
TaskTimer is a command-line tool written in Go that allows you to execute commands at a specified time or after a delay. It also supports repeated execution at regular intervals.

## Installation
```sh
# Build the project
go build -o bin/tasktimer.exe src/main.go
```

## Usage
```sh
# Execute a command after N seconds
./bin/tasktimer -A 60 -C "echo Hello World"

# Execute a command at a specified time
./bin/tasktimer -T 20:00 -C "echo Hello World"

# Execute a command at a specified date and time
./bin/tasktimer -D 10-01 -T 20:00 -C "echo Hello World"

# Execute a command repeatedly at regular intervals
./bin/tasktimer -C "echo Hello World" -frequency 60
```

## Command-line Arguments
- `-C, --cmd`: Command to execute
- `-T, --time`: Execution time (format: HH:MM)
- `-D, --date`: Execution date (format: MM-DD)
- `-A, --after`: Execute after N seconds
- `-frequency`: Execute every N seconds
- `-h, --help`: Show help information