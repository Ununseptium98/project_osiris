package main

import (
	"fmt"
	"osiris/procexplorer"
	"osiris/regmanip"
)

func main() {

	regmanip.TryAccess(`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, 1)
	procexplorer.Testing()
	fmt.Printf("PATH : %d", procexplorer.GetProcessExePath(19784))
}
