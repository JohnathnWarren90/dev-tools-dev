package main

import (
	"dev/generator"
	"embed"
	"fmt"
	"os"
)

//go:embed templates/go/*.tpl
var goTemplates embed.FS

func main() {
	var commands []string = []string{
		"new", "-config", "-list", "--version",
	}
	var languages []string = []string{
		"go", "python", "java",
	}
	var projectTypes []string = []string{
		"cli", "tui", "server", "website",
	}

	if len(os.Args) < 2 {
		fmt.Println("missing command")
		return
	}

	command := os.Args[1]

	if !Contains(commands, command) {
		fmt.Println("Command Unknown: -", command)
		return
	}

	switch command {
	case "new":
		if len(os.Args) < 5 {
			fmt.Println("missing arguements")
			return
		}

		language := os.Args[2]
		projectName := os.Args[3]
		projectType := os.Args[4]

		if !Contains(languages, language) {
			fmt.Println("Language Unsupported: -", language)
			return
		}
		if !HasName(projectName) {
			fmt.Println("Project Cannot be Empty: -", projectName)
			return
		}
		if !Contains(projectTypes, projectType) {
			fmt.Println("Unknown Project Type: -", projectType)
			return
		}

		generator.CreateProject(goTemplates, language, projectName, projectType)
		fmt.Println("Project Created: ", projectName)

	case "-config":
		fmt.Println("config")

	case "-list":
		fmt.Println("list")

	case "--version":
		Version()
	default:
		fmt.Println("unknown command")
	}

	fmt.Println(os.Args)
}

func Contains(list []string, value string) bool {
	for _, v := range list {
		if value == v {
			return true
		}
	}
	return false
}

func HasName(input string) bool {
	return input != ""
}

func Version() {
	fmt.Printf("dev\nVersion: 0.0.1\n")
}
