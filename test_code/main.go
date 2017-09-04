package main

import "os"
import "fmt"
import "bufio"

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ファイル無し")
		os.Exit(1)
	}
	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		fmt.Println(scnr.Text())
	}
	os.Exit(0)
}
