package main

import (
	"osiris/osiris"
)

func main() {
	//osiris.GetAllKeyValuesJSON(`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, 1, false, "")

	osiris.PrintBanner()
	osiris.PromptMenu()

}
