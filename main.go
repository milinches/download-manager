package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Download struct {
	url          string
	targetPath   string
	totalSection int
}

func (d *Download) Do() error {
	fmt.Println("Making Connection...")
	req, err := d.GetNewRequest("HEAD")
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("Got: %v\n", resp.StatusCode)

	return nil
}

func (d *Download) GetNewRequest(method string) (*http.Request, error) {
	req, err := http.NewRequest(method, d.url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Download Manager")
	return req, nil
}

func main() {
	startTime := time.Now()
	d := Download{
		url: "file:///C:/Users/HP/Documents/Languages/Rust/rust.pdf",
		targetPath: "rust.pdf",
		totalSection: 10,
	}

	if err := d.Do(); err != nil {
		log.Printf("An error occured while downloading file: %s\n", err.Error())
	}

	fmt.Printf("Download Completed in %v seconds\n", time.Now().Sub(startTime).Seconds())
}
