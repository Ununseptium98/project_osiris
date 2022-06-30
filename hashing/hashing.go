package hashing

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
)

func PathTomd5(PathToFile string) string {
	/*
		This funcitons takes a path to a file then returns its hash
		INPUT : Path as string
		OUTPUT : md5 Hash as string


	*/

	file, err := os.ReadFile(PathToFile)

	if err != nil {
		if os.IsPermission(err) {
			fmt.Printf("permission error")
		}

		log.Fatal(err)

	}

	hash := md5.New()

	md5 := hash.Sum(file)

	return (string(md5))

}
