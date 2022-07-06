package osiris

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/satheshshiva/go-banner-printer/banner"
)

func PromptMenu() {

	prompt := promptui.Select{
		Label: " Please select YAFT running mode",
		Items: []string{"Interactive", "Agent"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

	if result == "Interactive" {
		InteractiveMode()
	} else {
		AgentMode()
		PromptMenu()
	}

}

func InteractiveMode() {
	prompt := promptui.Select{
		Label: "What would you like to do ?",
		Items: []string{"Print process exe with PID", "Print process exe with hash",
			"Get Registry Key Values", "Get Registry key's SubKeys' value",
			"Get Watcher report", "Append path to WatchList",
			"Enroll agent", "Send file hash"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

	switch result {
	case "Print process exe with PID": //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

		exemap := GetProcessExePathMap()
		for pid, exePath := range exemap {
			fmt.Println("PID = ", pid, "\n Exe Path =", exePath, "\n----")
		}

		InteractiveMode()
	case "Print process exe with hash": //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

		exeHashMap := GetProcessExeHashMap()
		for path, hash := range exeHashMap {
			fmt.Println("Path = ", path, "\n Hash >", hash, "\n----")
		}

		InteractiveMode()

	case "Get Registry Key Values": //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		RegKeyPrompt("key")

		InteractiveMode()

	case "Get Registry key's SubKeys' value": //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		RegKeyPrompt("subkey")
		InteractiveMode()

	case "Get Watcher report": //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

		watcherReport := WatcherMap()
		for path, hash := range watcherReport {
			fmt.Println("Path = ", path, "\n Hash >", hash, "\n----")
		}
		InteractiveMode()
	case "Append path to WatchList": //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		WatcherAppendPrompt()
		InteractiveMode()

	case "Enroll agent": //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		fmt.Println("***Enrolling Agent***")

		err := EnrollAgent("192.168.30.11")
		if err != nil {
			fmt.Println("Agent enrollement failed. Server side issue")
		}
		fmt.Println("Agent successfully enrolled or already been enrolled .")
		InteractiveMode()

	case "Send file hash": //>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

		SendHashPrompt()
		InteractiveMode()

	}

}

func AgentMode() {

	fmt.Println("**AGENT MODE RUNNING**")
	ReadTaskJson("actions.json")
	time.Sleep(3 * time.Second)

	fmt.Println("***Reading task JSON file***")

	time.Sleep(3 * time.Second)

	fmt.Println("***Executing tasks in the JSON file***")

	ExecuteTaskJson("actions.json")

	fmt.Println("***Find json results in the application folder ***")

}

func RegKeyPrompt(keytype string) {

	prompt := promptui.Select{
		Label: "Select a registry key or select custom key ",
		Items: []string{"SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", "Custom"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

	if keytype == "key" {
		if result == "Custom" {
			CustomKeyPrompt("key")
		} else {
			GetAllKeyValuesJSON(result, 1, false, "")
		}
	} else {
		if result == "Custom" {
			CustomKeyPrompt("subkey")
		} else {
			GetSubKeyValues(result, 1, -1)
		}

	}

}

func CustomKeyPrompt(keytype string) {

	validate := func(input string) error {
		err := TestRequestKey(input)
		if err != nil {
			return errors.New("Invalid Key Format")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Registry key",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You requested %q\n", result)

	if keytype == "key" {
		GetAllKeyValuesJSON(result, 1, false, "")

	} else {
		GetSubKeyValues(result, 1, -1)
	}

}

func WatcherAppendPrompt() {

	/*
		validate := func(input string) error {
			_, err := os.Stat(input)
			if err != nil {
				return errors.New("Path doesn't match any file nor directory")
			}
			return nil
		}*/

	prompt := promptui.Prompt{
		Label: "Enter file or directory path. Directories end with a / or a \\",
	}

Prompt:
	flag := 1

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You added this to Watchile file : %q\n", result)

	_, err = os.Stat(result)
	if err != nil {
		fmt.Println("/!\\ /!\\ Path doesn't match any file nor directory. Retry :")
		flag = 0
	}

	if flag != 1 {
		goto Prompt
	}

	AppendWatchList(result)

	fmt.Println("Path added to Watcher's watchlist")

}

func SendHashPrompt() {

	prompt := promptui.Prompt{
		Label: "Enter file path with its file extension. ",
	}

Prompt:
	flag := 1

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You request this file : %q\n", result)

	_, err = os.Stat(result)
	if err != nil {
		fmt.Println("/!\\ /!\\ Path doesn't match any file. Retry :")
		flag = 0
	}

	if flag != 1 {
		goto Prompt
	}

	SendHash("192.168.30.11", result)

}
func PrintBanner() {
	_ = banner.Print(nil)

}
