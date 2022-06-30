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

	len, err := file.WriteString("FAK\n")
	len, err = file.WriteString("\n")

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
