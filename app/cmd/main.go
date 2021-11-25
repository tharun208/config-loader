package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tharun208/config-loader/pkg/config"
)

var initConfig *config.Config

func main() {
	initConfig = config.NewConfig()
	printRepl()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}
		processString(text)
	}
}

func processString(args string) {
	val := strings.Split(args, " ")
	switch val[0] {
	case "load":
		if len(val) <= 1 {
			fmt.Fprintf(os.Stderr, "%v\n", "Error: please provide a file to load")
			break
		}
		files := val[1:]
		if err := initConfig.LoadFiles(files...); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		fmt.Fprintf(os.Stdin, "%v\n", "successfully loaded files")
	case "get":
		if len(val) <= 1 {
			fmt.Fprintf(os.Stderr, "%v\n", "Error: please provide a key")
			break
		}
		value := initConfig.GetConfigValueAsString(val[1])
		fmt.Fprintf(os.Stdout, "%v\n", value)
	case "clear":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		printRepl()
	default:
		fmt.Print("Invalid action")
	}
}

func printRepl() {
	fmt.Println("config-loader> This are the Avaliable commands: ")
	fmt.Println("config-loader> load   - load configuration files [args: [] -> list of file(s)]")
	fmt.Println("config-loader> clear  - Clear the Terminal Screen ")
	fmt.Println("config-loader> exit   - Exits the config loader REPL ")
	fmt.Println("config-loader> get    - get the value for given key name [args: [0] -> key]")
	fmt.Println("config-loader> ")
}
