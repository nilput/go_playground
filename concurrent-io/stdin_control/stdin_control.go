package stdin_control

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"concurrent-io/http_serve"
)

func inputCommandError(errMsg string) {
	fmt.Errorf("Error: %s", errMsg)
}

func updateHttpContent(content string) {
	http_serve.
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
		content := strings.Join(part[1:])
		updateHttpContent(content)
	}
	else if first == "quit" || first == "exit" {
		return true
	}
}

func stdinControl(ch chan bool) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		proccessLine(line)
		s := fmt.Scan()
	}
	ch <- true
}
