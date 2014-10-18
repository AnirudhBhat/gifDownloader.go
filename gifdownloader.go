package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
)

func main() {
	response, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			err = ioutil.WriteFile(os.Args[2], contents, 0644)
			if err != nil {
				fmt.Printf("%s", err)
			}
		}
	}
}