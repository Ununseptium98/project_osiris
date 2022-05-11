package main

import (
	"osiris/procexplorer"
	"osiris/regmanip"
)

func main() {

	regmanip.TryAccess(`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, 1)
	procexplorer.Testing()
}
