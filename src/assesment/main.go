package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	count := 100
	for i := 1; i <= count; i++ {
		wg.Add(1)
		// fileName := "sample.pdf"
		fileName := fmt.Sprintf("sample%v.pdf", i)
		URL := "http://www.africau.edu/images/default/sample.pdf"
		go downloadFile(URL, fileName, &wg)
	}
	wg.Wait()
}

func downloadFile(URL string, fileName string, wg *sync.WaitGroup) {
	fmt.Println("download-", fileName)
	defer wg.Done()
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error on get response", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create("files" + "/" + fileName)
	if err != nil {
		fmt.Println("error on create file", err)
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("error on write file", err)
	}
	fmt.Println("Successfully downloaded")
}
