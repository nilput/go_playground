package cmd

import (
	"bufio"
	"concurrent_io/http"
	"fmt"
	"os"
	"strings"
)

func inputCommandError(errMsg string) {
	fmt.Errorf("Error: %s", errMsg)
}

func updateHttpContent(content string) {
	http.Content = content
}

func usageHelp() {
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("\tHELP:\tShow this help guide.")
	fmt.Println("\tSET:\tUpdate HTTP response content.")
	fmt.Println("\tEXIT:\tQuit.")
}

func proccessLine(line string) (quit bool) {
	line = strings.Trim(line, " \t\n")
	parts := strings.Split(line, " ")
	if len(parts) == 0 {
		return
	}
	first := parts[0]
	first = strings.ToLower(first)
	if first == "set" {
		if len(parts) < 2 {
			inputCommandError("set command takes 2+ arguments")
			return false
		}
		content := strings.Join(parts[1:], " ")
		updateHttpContent(content)
	} else if first == "help" || first == "?" {
		usageHelp()
	} else if first == "quit" || first == "exit" {
		return true
	}
	return false
}

func prompt() {
	fmt.Print("$ ")
}

func StdinControl(ch chan bool) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		prompt()
		scanner.Scan()
		line := scanner.Text()
		quit := proccessLine(line)
		if quit {
			break
		}
	}
	ch <- true
}
