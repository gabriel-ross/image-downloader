package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	var URL string
	var fileName string

	fmt.Print("Enter URL to download from: ")
	fmt.Scanln(&URL)
	fmt.Print("Enter destination file name: ")
	fmt.Scanln(&fileName)

	err := downloadFile(URL, fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %s successfully downloaded to local directory", fileName)
}

func downloadFile(URL string, fileName string) error {
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// check status code for successful download
	if res.StatusCode != 200 {
		return errors.New("Received non-200 status code")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	bytesCopied, err := io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%d bytes copied to new file successfully", bytesCopied)
	return nil

}
