package main

import (
	"fmt"
	"io/ioutil"
	"maxFreguentWords/find"
)

func main()  {
	file, err := ioutil.ReadFile("./mobydick.txt")
	if err != nil {
		fmt.Println(err)
	}
	asd, err := find.MaxFrequentWords(file, 3)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(asd)
}
