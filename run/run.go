package main

import(
	"osiris/regmanip"
)

func main(){

	regmanip.TryAccess(`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, 1)

}
