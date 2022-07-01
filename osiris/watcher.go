package osiris

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreateWatchList() {
	file, err := os.Create("watchList.txt") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done

	len, err := file.WriteString(`C:\Users\Nazim\Videos\` + "\n")
	len, err = file.WriteString(`C:\Users\Nazim\Videos\World Of Warcraft` + "\n")
	len, err = file.WriteString(`C:\Users\Nazim\Videos\World Of Warcraft\bite.png` + "\n")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())

}

func AppendWatchList(pathToWatch string) {
	/*
		Appends string pathToWatch to WatchList file
	*/
	if string(pathToWatch[len(pathToWatch)-1]) != "\n" {
		pathToWatch = pathToWatch + "\n"
	}
	file, err := os.OpenFile("watchlist.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(pathToWatch); err != nil {
		log.Fatal(err)
	}
}

func ReadWatchList() []string {
	/*
		Return a table containing all the paths in the WatchList file
	*/

	var pathsTable []string
	file, err := os.Open("watchlist.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		pathsTable = append(pathsTable, scanner.Text())

	}
	err = scanner.Err()

	if err != nil {
		log.Fatal(err)
	}

	return pathsTable

}

func WriteJsonWatchList() map[string]string {
	/*
		writes in json format the hash of the directories or files
	*/

	pathsTable := ReadWatchList() //Reads file paths from watchlist

	pathHashMap := make(map[string]string)

	for _, path := range pathsTable { //iterates over the paths

		lastChar := string(path[len(path)-1]) //gets the last char

		if lastChar == `\` || lastChar == "/" { //test if the path describes a file or a directory
			//if it is a directory
			a, err := DirHash(path)

			if err != nil {
				log.Fatal(err)
			}
			pathHashMap[path] = string(a)

		} else { // else it's a file
			pathHashMap[path] = PathTomd5(path)
		}

	}

	return pathHashMap
}
