package hashing

import (
	"crypto/md5"
	"io/ioutil"
	"log"
)

func Testing(PathToFile string) string {
	//file, err := os.OpenFile(PathToFile, os.O_RDONLY, 0755)

	file, err := ioutil.ReadFile(PathToFile)

	if err != nil {
		log.Fatal(err)
	}

	hash := md5.New()

	md5 := hash.Sum([]byte(file))

	return (string(md5))

}
