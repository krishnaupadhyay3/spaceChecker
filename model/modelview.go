package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"spaceChecker/utils"
	"strconv"
)

func (dir *Directory) SpaceApi(w http.ResponseWriter, r *http.Request) {
	QueryDict := r.URL.Query()
	dir_path := QueryDict.Get("dir")
	fmt.Println(os.UserHomeDir())
	value, err := utils.CheckDir(dir_path)

	w.Header().Set("Content-Type", "application/json")
	errResp := ""
	respStruct := RespStatus{}
	if err != nil {
		errResp = err.Error()
		respStruct.Error = errResp
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp := fmt.Sprintf("%.6f GB", value)
		respStruct.Size = resp
	}

	fmt.Println(json.NewEncoder(w).Encode(respStruct))
}

func (dir *Directory) AddSpace(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println(os.UserHomeDir())
		body, err := ioutil.ReadAll(r.Body) // check for errors

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(body)
		keyVal := JsonInput{}
		err = json.Unmarshal(body, &keyVal)
		fmt.Println(keyVal)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		directory := keyVal.Path
		size := keyVal.Size
		fmt.Println(size)
		fmt.Println(directory)
		sizeFloat, err := strconv.ParseFloat(size, 64)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(w.Write([]byte("incorrect size format")))
			return
		}
		homeDir, _ := os.UserHomeDir()
		file_path := homeDir + string(os.PathSeparator) + directory
		fmt.Println(file_path)
		dir.Add(file_path, sizeFloat)
		fmt.Fprintf(w, "added successfully")
	} else {
		log.Println(w.Write([]byte("incorrect size format")))
	}
}
