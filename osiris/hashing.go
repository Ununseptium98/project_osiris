package osiris

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gosimple/hashdir"
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

func DirHash(path string) (string, error) {
	//TODO calculate hash of files into the directory, then calculate a hash of the directory
	//based on the file hashes ?

	dirHash, err := hashdir.Make(path, "md5")

	if err != nil {
		return "Error - failed to hash the directory, verify the path ?", err
	}

	fmt.Println("hash => ", dirHash)
	return dirHash, err

}
