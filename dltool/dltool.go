package main

import (
	"encoding/json"
	//"bufio"
	"fmt"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"net/http"
	"path"
	//"strings"
)

type FileInfo struct {
	Url string `json:"url"`
	Name string `json:"name"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func downloadFile(url string, savePath string) {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	// makesure directory
	//body, err := ioutil.ReadAll(resp.Body)
	
	f, err := os.Create(savePath)
	check(err)
	defer f.Close()

	io.Copy(f, resp.Body);
	
}

func main() {
	fileListPtr := flag.String("filelist", "", "path to file list json")
	outDirPtr := flag.String("outDir", "./", "directory path to store download file")
	flag.Parse()

	fmt.Println("fileList: ", *fileListPtr)
	fmt.Println("outDir :", *outDirPtr)

	if len(*fileListPtr) == 0 {
		fmt.Println("fileList is required")
		os.Exit(-1);
	}

	if _, err := os.Stat(*fileListPtr); os.IsNotExist(err) {
		fmt.Printf("file [%s]does not exist\n", *fileListPtr)
	}

	if _, err := os.Stat(*outDirPtr); os.IsNotExist(err) {
		if err := os.Mkdir(*outDirPtr, os.ModeDir); err != nil {
			check(err)
		}
	}


	// read json

	dat, err := ioutil.ReadFile(*fileListPtr)
    check(err)

	var fl []FileInfo
	json.Unmarshal(dat, &fl)
	
	for i := 0; i < len(fl); i++ {
		fmt.Printf("Start to download %s from %s\n", fl[i].Name, fl[i].Url)
		downloadFile(fl[i].Url, path.Join(*outDirPtr, fl[i].Name))
	}
}
