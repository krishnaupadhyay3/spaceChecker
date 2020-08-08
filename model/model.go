package model

import (
	"log"
	"spaceChecker/utils"
)

type JsonInput struct {
	Path string `json:"path"`
	Size string `json:"size"`
}
type Directory struct {
	Dirs    []string
	DirSize map[string]float64
	Delete  chan string
}

type RespStatus struct {
	Size  string `json:"size"`
	Error string `json:"error"`
}

// method on struct Directory , checks the limit on the given  parameter directory,
//return a Boolean whether given directory exceed limits or not
func (dir *Directory) CheckLimit(dirPath string) bool {
	size := dir.DirSize[dirPath]

	currentSize, err := utils.CheckDir(dirPath)
	log.Println("checking limit")
	if err != nil {
		log.Println(err)
	}
	if currentSize > 0.8*size {

		log.Printf("exceed limit. directory %s send for cleaning", dirPath)
		return true
	}
	return false
}

// Takes Directory Struct as input and Checks the Limit on the paths
// send path on Delete channel if exceed limit
func CheckDirs(dirs *Directory) {
	for _, dir := range dirs.Dirs {
		IsFull := dirs.CheckLimit(dir)
		if IsFull {

			go func() {
				dirs.Delete <- dir
			}()

		}
	}
}

// Method on Struct Directory , add a path for monitoring for space.
// takes a string path and Float64 size as input

func (dir *Directory) Add(dirName string, size float64) {
	_, ok := dir.DirSize[dirName]
	if !ok {
		dir.DirSize[dirName] = size
		dir.Dirs = append(dir.Dirs, dirName)
	}
}

// Method on Struct Directory  remove a path from monitoring.
// takes a string path as input
func (dir *Directory) Remove(dirName string) {
	for i, val := range dir.Dirs {
		if dirName == val {
			dir.Dirs = append(dir.Dirs[:i], dir.Dirs[i+1:]...)
		}
	}
}
