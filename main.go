package main

import (
	"io"
	"net/http"
	"os"
)

func fetchData (url string) ([]byte,error){

	resp,err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)

	return result, err

}

func main(){
	result, err := fetchData("http://www.google.com")

    if err!= nil {
        panic(err)
    }

    outFile, err := os.Create("output.txt")

	if err != nil {
		panic(err)
	}


}