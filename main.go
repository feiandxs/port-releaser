package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// findProcessWindows 在Windows上查找占用特定端口的进程
func findProcessWindows(port string) ([]string, error) {
	cmd := exec.Command("cmd", "/c", "netstat -ano | findstr :"+port)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var processes []string
	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "LISTENING") {
			processes = append(processes, line)
		}
	}

	return processes, nil
}

// findProcessUnix 在Unix-like系统(macOS和Linux)上查找占用特定端口的进程
func findProcessUnix(port string) ([]string, error) {
	cmd := exec.Command("lsof", "-i", "tcp:"+port)
	output, err := cmd.CombinedOutput()

	// 如果错误是因为没有找到匹配的进程，则返回空切片而不是错误
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// lsof 未找到匹配的进程
			if exitErr.ExitCode() == 1 {
				return []string{}, nil
			}
		}
		return nil, err
	}

	var processes []string
	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "LISTEN") {
			processes = append(processes, line)
		}
	}

	return processes, nil
}

// killProcess 使用适当的命令终止指定PID的进程
func killProcess(pid string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("taskkill", "/f", "/pid", pid)
	} else {
		cmd = exec.Command("kill", "-9", pid)
	}

	return cmd.Run()
}

// confirmAndKill 显示进程信息，并询问用户是否终止进程
func confirmAndKill(processes []string) {
	reader := bufio.NewReader(os.Stdin)
	for _, process := range processes {
		fmt.Printf("进程: %s\n", process)

		// 提取PID，Windows和Unix-like系统的输出格式可能有差异
		fields := strings.Fields(process)
		var pid string
		if runtime.GOOS == "windows" {
			pid = fields[len(fields)-1] // Windows: PID在最后
		} else {
			pid = fields[1] // Unix-like: PID通常在第二列
		}

		fmt.Print("是否终止此进程? [y/N]: ")
		response, _ := reader.ReadString('\n')
		if strings.TrimSpace(response) == "y" {
			err := killProcess(pid)
			if err != nil {
				fmt.Printf("无法终止进程 %s: %v\n", pid, err)
			} else {
				fmt.Printf("进程 %s 已终止。\n", pid)
			}
		}
	}
}

func main() {
	// 解析命令行参数以获取端口号
	flag.Parse()
	port := flag.Arg(0)
	if port == "" {
		fmt.Println("请指定一个端口号。")
		return
	}

	var processes []string
	var err error

	// 根据操作系统调用不同的函数
	switch runtime.GOOS {
	case "windows":
		processes, err = findProcessWindows(port)
	case "darwin", "linux":
		processes, err = findProcessUnix(port)
	default:
		fmt.Printf("不支持的操作系统: %s\n", runtime.GOOS)
		return
	}

	if err != nil {
		fmt.Printf("查找端口时发生错误: %v\n", err)
		return
	}

	// 如果没有找到进程，则退出
	if len(processes) == 0 {
		fmt.Printf("没有找到占用端口 %s 的进程。\n", port)
		return
	}

	// 显示进程信息并询问用户是否终止
	confirmAndKill(processes)
}
