package osiris

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func CreateWatchList() {

	if _, err := os.Stat("WatchList.txt"); err == nil { //file exists
		return
	} else {
		file, err := os.Create("watchList.txt") // Truncates if file already exists, be careful!
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer file.Close() // Make sure to close the file when you're done

	}
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

func WatcherMap() map[string]string {
	/*
		writes in json format the hash of the directories or files
	*/

	pathsTable := ReadWatchList() //Reads file paths from watchlist
	pathHashMap := make(map[string]string)

	for _, path := range pathsTable { //iterates over the pathsZ

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

func WriteWatcherJSON() {

	date := strings.Split(time.Now().String(), " ")[0] //Gets date

	watcherReport := WatcherMap()
	json, _ := json.Marshal(watcherReport)
	ioutil.WriteFile(date+"_WatcherReport.json", json, os.ModePerm)

	fmt.Println("OUTPUT : " + date + "_WatcherReport.json")
}
