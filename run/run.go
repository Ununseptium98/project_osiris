package main

import (
	"osiris/osiris"
)

func main() {
	//osiris.GetAllKeyValuesJSON(`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, 1, false, "")

	osiris.PrintBanner()
	osiris.PromptMenu()

	//osiris.EnrollAgent("192.168.30.11")
	//fmt.Printf("\n")

	//osiris.SendHash("192.168.30.11", `C:\Users\nmeza\Pictures\bite.jpg`)
}
