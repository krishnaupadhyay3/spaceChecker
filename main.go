package main

import (
	"log"
	"net/http"
	"spaceChecker/model"
	"spaceChecker/utils"
	"time"
)

func main() {
	dir := model.Directory{DirSize: make(map[string]float64), Delete: make(chan string)}
	ticker := time.NewTicker(10 * time.Second)

	// dir := "/home/krishna/rar"
	go func(dir *model.Directory) {
		for {
			select {
			case <-ticker.C:

				model.CheckDirs(dir)

			case path := <-dir.Delete:
				utils.RemoveLast(path)
			}
		}
	}(&dir)
	http.HandleFunc("/v1/space", dir.SpaceApi)
	http.HandleFunc("/v1/space/add", dir.AddSpace)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
