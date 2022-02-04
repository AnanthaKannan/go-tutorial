package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	C1 := make(chan error)
	var wg sync.WaitGroup
	count := 100
	for i := 1; i <= count; i++ {
		wg.Add(1)
		// fileName := "sample.pdf"
		fileName := fmt.Sprintf("sample%v.pdf", i)
		URL := "http://www.africau.edu/images/default/sample.pdf"
		go downloadFile(URL, fileName, &wg, C1)
	}

	successCount := 0
	failedCount := 0
	for i := 1; i <= count; i++ {

		select {
		case error_ := <-C1:
			// fmt.Println("data received", opt1)
			if error_ != nil {
				fmt.Println(error_)
				fmt.Println("err", error_)
				failedCount++
			} else {
				successCount++
			}
		}
	}
	fmt.Println("successCount", successCount, "failedCount", failedCount)

	wg.Wait()
}

func downloadFile(URL string, fileName string, wg *sync.WaitGroup, channel1 chan error) {
	fmt.Println("download-", fileName)
	defer wg.Done()
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error on get response", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		channel1 <- errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create("files" + "/" + fileName)
	if err != nil {
		channel1 <- err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		channel1 <- err
	}
	channel1 <- nil
}
