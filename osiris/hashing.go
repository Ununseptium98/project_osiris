package osiris

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/mod/sumdb/dirhash"
)

func PathTomd5(PathToFile string) string {

	file, err := os.Open(PathToFile)

	if err != nil {
		if os.IsPermission(err) {
			fmt.Printf("permission error")
		}

		log.Fatal(err)

	}

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))

}

func DirHash(dir []string) string {
	//TODO calculate hash of files into the directory, then calculate a hash of the directory
	//based on the file hashes ?

	var hash = dirhash.Hash(dir)

	hash := dirhash.HashDir(`C:\Users\Nazim\Videos`)

	return hash

}
