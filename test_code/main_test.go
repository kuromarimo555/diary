package main

import (
	"fmt"
)

func main() {
	test_string := "test"
	err := pre(test_string)

	if err == 1 {
		fmt.Println("errer")
	} else {
		fmt.Println("成功")
	}
}
func pre(code string) (err int) {

	//code := "test"
	if code == "sum" {
		return 1
	} else {
		return 0
	}
	return 3

	/*(int, err) {
	    code := "test"
	  	if code == "hoge" {
	  		return 1, nil
	  	}
	  	return 0, errors.New("code must be hoge")
	*/
}
