package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

type Template struct {
	Folders []string
	Files   []string
}

func CreateProject(goTemplates embed.FS, language, projectName, projectType string) {
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Println("Project Folder Failed: ", err)
		return
	}
	if language == "go" {
		goTemplate := Template{
			Folders: []string{
				"utils",
				"models",
				"src",
				"tests",
			},
			Files: []string{
				"main.go",
			},
		}
		//Create GO Outline
		for _, folder := range goTemplate.Folders {
			path := filepath.Join(projectName, folder)
			err := os.MkdirAll(path, 0755)
			if err != nil {
				fmt.Println("Folder Creation Failed: ", err)
				return
			}
		}

		template, err := goTemplates.ReadFile("templates/go/main.go.tpl")
		if err != nil {
			fmt.Println("Reading Template Failed:", err)
			return
		}

		for _, file := range goTemplate.Files {
			filepath := filepath.Join(projectName, file)
			err := os.WriteFile(filepath, template, 0644)
			if err != nil {
				fmt.Println("File Failed:", err)
				return
			}
		}
	}

	/*if language == "python" {
		pyTemplate := Template{
			Folders: []string{
				"Classes",
				"Src",
				"Tests",
			},
			Files: []string{
				"main.py",
			},
		}
		//Create Python Template
	}*/

	/*if language == "java" {
		javaTemplate := Template{
			Folders: []string{
				"Src",
				"Tests",
			},
			Files: []string{
				"Main.java",
			},
		}
		//Create Java Template
	}*/

}
