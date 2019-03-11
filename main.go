package main

import (
	"fmt"
	"bufio"
	"os"
	"path/filepath"

	. "github.com/kkdai/youtube"
)

func read() string {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	y := NewYoutube(true)
	
	fmt.Print("Enter a youtube url: ")
	url := read()
	fmt.Print("Enter the export location of the video: ")
	loc := read()
	fmt.Print("Enter the name of the export: ")
	name := read()

	full_name := name + ".mp3"
	
	if err := y.DecodeURL(url); err != nil {
		fmt.Println("err:", err)
	}
	if err := y.StartDownload(filepath.Join(loc, full_name)); err != nil {
		fmt.Println("err:", err)
	}
}