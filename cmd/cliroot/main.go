package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cli-crud/internal/data/session"
)

func main() {
	session := &session.Session{}
	RunCli(session)
}

func RunCli(session *session.Session) {
	fmt.Println("Welcome to CRUD CLI for SERVICES. Type `help` for commands.")
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !reader.Scan() {
			break
		}

		line := strings.TrimSpace(reader.Text())
		if line == "" {
			continue
		}

		args := strings.Fields(line)
		cmd := strings.ToUpper(args[0])
		params := args[1:]

		conf, exists := Config[cmd]
		if !exists {
			fmt.Println("Unknown command:", cmd)
			// fmt.Println("Type `help` for commands.")
			printHelp()
			continue
		}

		if len(params) < conf.minParams {
			fmt.Printf("Usage: %s %s\n", cmd, conf.params)
			continue
		}

		if conf.handler != nil {
			conf.handler(params, session)
		} else {
			switch cmd {
			case CmdHelp:
				printHelp()
			case CmdExit:
				fmt.Println("Goodbye!")
				return
			}
		}
	}
}

func printHelp() {
	fmt.Println("Available commands:")

	maxLen := 0
	for name := range Config {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	for name, conf := range Config {
		padding := strings.Repeat(" ", maxLen-len(name)+10)
		fmt.Printf("  %s%s%s\n", name, padding, conf.params)
	}
	
}
