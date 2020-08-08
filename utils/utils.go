package utils

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
)

// Takes a path as input and return as array of os.FileInfo sorted by modified time
func GetLastModifiedList(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})
	return files
}

// takes a path as input and removes the 90 percent files from the path
func RemoveLast(path string) {
	files := GetLastModifiedList(path)
	size := float64(len(files))
	val := int(math.Ceil(size * 90 / 100))
	HalfFiles := files[:val]
	for _, file := range HalfFiles {

		file_path := path + string(os.PathSeparator) + file.Name()

		log.Println("deleting file")
		log.Println(file.Name())
		log.Println(os.Remove(file_path))
	}

}

// takes a Path as input and return float64 size of the directory and error
func CheckDir(path string) (float64, error) {
	sum := 0.0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return sum, err
	}
	for _, fileName := range files {
		sum += float64(fileName.Size())
	}
	division := float64(1000 * 1000 * 1000)
	return sum / division, nil
}
