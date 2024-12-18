package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/feiandxs/port-releaser/i18n"
	"golang.org/x/text/language"
)

// detectLanguage 检测系统语言
func detectLanguage() string {
	langEnv := os.Getenv("LANG")
	langTag, err := language.Parse(langEnv)
	if err != nil {
		return "en" // 默认为英文
	}
	base, _ := langTag.Base()
	langCode := base.String()

	// 检查语言是否支持，如果不支持则返回英文
	if _, ok := i18n.Messages[langCode]; !ok {
		return "en"
	}
	return langCode
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
		fmt.Printf(i18n.Messages[lang]["process"], process)

		fields := strings.Fields(process)
		var pid string
		if runtime.GOOS == "windows" {
			pid = fields[len(fields)-1]
		} else {
			pid = fields[1]
		}

		fmt.Print(i18n.Messages[lang]["terminateProcess"])
		response, _ := reader.ReadString('\n')
		if strings.TrimSpace(response) == "y" {
			err := killProcess(pid)
			if err != nil {
				fmt.Printf(i18n.Messages[lang]["unableToTerminate"], pid, err)
			} else {
				fmt.Printf(i18n.Messages[lang]["processTerminated"], pid)
			}
		}
	}
}

func main() {
	lang := detectLanguage()

	flag.Parse()
	port := flag.Arg(0)
	if port == "" {
		fmt.Println(i18n.Messages[lang]["specifyPort"])
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
		fmt.Printf(i18n.Messages[lang]["unsupportedOS"], runtime.GOOS)
		return
	}

	if err != nil {
		fmt.Printf(i18n.Messages[lang]["errorFindingPort"], err)
		return
	}

	if len(processes) == 0 {
		fmt.Printf(i18n.Messages[lang]["noProcessFound"], port)
		return
	}

	confirmAndKill(processes, lang)
}
