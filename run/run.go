package main

import (
	"fmt"
	"osiris/osiris"
)

func main() {

	//osiris.TestKey(`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, 1)
	//osiris.TryAccess(`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, 1)
	//osiris.GetAllKeyValues(`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, 1)
	//osiris.GetSubKeyValues(`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, 1, 10)

	//fmt.Println(string(osiris.GetProcessExeHashJson()))

	osiris.CreateWatchList()

	for i, lines := range osiris.ReadWatchList() {

		fmt.Println("lines = ", lines, " i =", i)
	}

	fmt.Println(osiris.WriteJsonWatchList())

}
