package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"

	"cli-crud/internal/data/session"
)

func main() {
	sess := initSession()
	RunCli(sess)
}

func initSession() *session.Session {
	return &session.Session{}
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
			printHelp()
			continue
		}

		if err := validateCommandParams(conf, params); err != nil {
			fmt.Println(err)
			continue
		}

		if conf.handler != nil {
			conf.handler(params, session)
		} else {
			handleNoHandler(cmd)
		}
	}
}

func handleNoHandler(cmd string) {
	switch cmd {
	case CmdHelp:
		printHelp()
	case CmdExit:
		fmt.Println("Goodbye!")
		os.Exit(0)
	}
}

func printHelp() {
	fmt.Println("Available commands:")

	commands := []string{}
	for name := range Config {
		commands = append(commands, name)
	}

	sort.Strings(commands)

	maxLen := 0
	for _, name := range commands {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	for _, name := range commands {
		conf := Config[name]
		padding := strings.Repeat(" ", maxLen-len(name)+10)
		fmt.Printf("  %s%s%s\n", name, padding, conf.params)
	}
	
}

