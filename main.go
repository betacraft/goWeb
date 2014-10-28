package main

import (
	"fmt"
	"os"
	"os/exec"
)

func printHelp() {
	fmt.Println("Help:")
	fmt.Println("create <project_name> => To initialize a new project")
	fmt.Println("run => To run the current web app")
}

func createProject(projectName string) {
	fmt.Println("Creating project", projectName)
	fmt.Print("[1] Creating main dir => ")
	if err := os.Mkdir(projectName, 0777); err != nil {
		fmt.Println("failed\n")
		fmt.Println("Reason =>", err.Error())
		return
	}
	fmt.Print("done\n")
	fmt.Print("[2] Creating project structure => ")
	if err := os.Mkdir(projectName+"/src", 0777); err != nil {
		fmt.Println("failed\n")
		fmt.Println("Reason =>", err.Error())
		return
	}
	if err := os.Mkdir(projectName+"/bin", 0777); err != nil {
		fmt.Println("failed\n")
		fmt.Println("Reason =>", err.Error())
		return
	}
	if err := os.Mkdir(projectName+"/pkg", 0777); err != nil {
		fmt.Println("failed\n")
		fmt.Println("Reason =>", err.Error())
		return
	}
	if err := os.Mkdir(projectName+"/src/"+projectName, 0777); err != nil {
		fmt.Println("failed\n")
		fmt.Println("Reason =>", err.Error())
		return
	}
	fmt.Print("done\n")
	// fmt.Print("[3] Initializing git repo => ")
	// os.Chdir(projectName + "/src" + projectName)
	// _, err := exec.Command("git init").Output()
	// if err != nil {
	// 	fmt.Println("failed\n")
	// 	fmt.Println("Reason =>", err.Error())
	// 	return
	// }
	// fmt.Print("done\n")
	fmt.Print("[4] Adding GoDeps => ")
	_, err := exec.Command("go get github.com/tools/godep").Output()
	if err != nil {
		fmt.Println("failed\n")
		fmt.Println("Reason =>", err.Error())
		return
	}
	fmt.Print("done\n")
	fmt.Print("[5] Adding GorrillaMux => ")
	_, err = exec.Command("go get github.com/gorilla/mux").Output()
	if err != nil {
		fmt.Println("failed\n")
		fmt.Println("Reason =>", err.Error())
		return
	}
	fmt.Print("done\n")
	fmt.Print("[6] Creating kernel file => ")
	_, err = exec.Command("touch kernel.go").Output()
	if err != nil {
		fmt.Println("failed\n")
		fmt.Println("Reason =>", err.Error())
		return
	}
	fmt.Print("done\n")
}

func main() {
	if len(os.Args) == 1 {
		printHelp()
		return
	}
	switch os.Args[1] {
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Help => create <project_name>")
			fmt.Println("Error => project-name is missing")
			return
		}
		projectName := os.Args[2]
		createProject(projectName)
		break
	case "run":
		break
	default:
		fmt.Println("Illegal command =>", os.Args[1])
		printHelp()
	}
}
