package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/text/language"
)

// 消息本地化
var msg map[string]map[string]string

// 初始化消息
func init() {
	msg = make(map[string]map[string]string)
	msg["en"] = map[string]string{
		"specifyPort":       "Please specify a port number.",
		"unsupportedOS":     "Unsupported operating system: %s\n",
		"errorFindingPort":  "Error finding port: %v\n",
		"noProcessFound":    "No process found occupying port %s.\n",
		"process":           "Process: %s\n",
		"terminateProcess":  "Do you want to terminate this process? [y/N]: ",
		"unableToTerminate": "Unable to terminate process %s: %v\n",
		"processTerminated": "Process %s has been terminated.\n",
	}
	msg["zh"] = map[string]string{
		"specifyPort":       "请指定一个端口号。",
		"unsupportedOS":     "不支持的操作系统: %s\n",
		"errorFindingPort":  "查找端口时发生错误: %v\n",
		"noProcessFound":    "没有找到占用端口 %s 的进程。\n",
		"process":           "进程: %s\n",
		"terminateProcess":  "是否终止此进程? [y/N]: ",
		"unableToTerminate": "无法终止进程 %s: %v\n",
		"processTerminated": "进程 %s 已终止。\n",
	}
}

// detectLanguage 检测系统语言
func detectLanguage() string {
	langEnv := os.Getenv("LANG")
	langTag, err := language.Parse(langEnv)
	if err != nil {
		return "en" // 默认为英文
	}
	base, _ := langTag.Base()
	return base.String()
}

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
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
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
func confirmAndKill(processes []string, lang string) {
	reader := bufio.NewReader(os.Stdin)
	for _, process := range processes {
		fmt.Printf(msg[lang]["process"], process)

		fields := strings.Fields(process)
		var pid string
		if runtime.GOOS == "windows" {
			pid = fields[len(fields)-1]
		} else {
			pid = fields[1]
		}

		fmt.Print(msg[lang]["terminateProcess"])
		response, _ := reader.ReadString('\n')
		if strings.TrimSpace(response) == "y" {
			err := killProcess(pid)
			if err != nil {
				fmt.Printf(msg[lang]["unableToTerminate"], pid, err)
			} else {
				fmt.Printf(msg[lang]["processTerminated"], pid)
			}
		}
	}
}

func main() {
	lang := detectLanguage()

	flag.Parse()
	port := flag.Arg(0)
	if port == "" {
		fmt.Println(msg[lang]["specifyPort"])
		return
	}

	var processes []string
	var err error

	switch runtime.GOOS {
	case "windows":
		processes, err = findProcessWindows(port)
	case "darwin", "linux":
		processes, err = findProcessUnix(port)
	default:
		fmt.Printf(msg[lang]["unsupportedOS"], runtime.GOOS)
		return
	}

	if err != nil {
		fmt.Printf(msg[lang]["errorFindingPort"], err)
		return
	}

	if len(processes) == 0 {
		fmt.Printf(msg[lang]["noProcessFound"], port)
		return
	}

	confirmAndKill(processes, lang)
}
