package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var input string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please drag osu Songs folder on this program")
		Exit()
	}
	songDict := os.Args[1]
	files, err := ioutil.ReadDir(songDict)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	for _, f := range files {
		if f.IsDir() && f.Name() != "Failed" {
			buf.WriteString(f.Name() + "\n")
		}
	}
	ex, err := os.Executable()
	savePath := filepath.Dir(ex)
	err = ioutil.WriteFile(savePath+"./song.txt", buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("save failed", err)
	}
	fmt.Println("save to ", savePath+"/song.txt")
	Exit()
}

func Exit() {
	fmt.Println("press enter to exit...")
	fmt.Scanln(&input)
	os.Exit(0)
}
