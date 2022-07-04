package osiris

import (
	"errors"
	"fmt"

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
	}

}

func InteractiveMode() {
	prompt := promptui.Select{
		Label: "What would you like to do ?",
		Items: []string{"Print process exe with PID", "Print process exe with hash",
			"Get Registry Key Values", "Get Registry key's SubKeys' value",
			"Get Watcher report", "Append path to WatchList",
			""},
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

	}

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
func PrintBanner() {
	_ = banner.Print(nil)

}
