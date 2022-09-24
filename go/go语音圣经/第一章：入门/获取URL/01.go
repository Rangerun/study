package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch err is %v\n", err)
			os.Exit(0)
		}
		io.Copy(os.Stdout, resp.Body)
		//b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			os.Exit(0)
		}
		//fmt.Printf("%s", b)
	}
}