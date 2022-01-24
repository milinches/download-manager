package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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

	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Can't process, response is %v\n", resp.StatusCode))
	}

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		fmt.Printf("Size is %v bytes\n", size)
	}

	sections := make([][2]int, d.totalSection)
	eachSize := size / d.totalSection
	fmt.Printf("Each Size is %v bytes\n", eachSize)
	
	for i := range sections {
		if i == 0 {
			sections[i][0] = 0
		} else {
			sections[i][0] = sections[i-1][1] + 1
		}

		if i < d.totalSection-1 {
			sections[i][1] = sections[i][0] + eachSize
		} else {
			sections[i][1] = size - 1
		}
	}
	fmt.Println(sections)

	for i, s := range sections {
		err = d.DownloadSession(i, s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Download) DownloadSession(i int, c[2]int) error {
	r, err := d.GetNewRequest("GET")
	if err != nil {
		return err
	}
	r.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", c[0], c[1]))
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Can't process, response is %v", resp.StatusCode))
	}

	fmt.Printf("Downloaded %v bytes for section %v\n", resp.Header.Get("Content-Length"), i)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("section-%v.tmp", i), b, os.ModePerm)
	if err != nil {
		return err
	}

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
		url: "https://github.com/dappuniversity/eth-todo-list/archive/refs/heads/master.zip",
		targetPath: "master.zip",
		totalSection: 10,
	}

	if err := d.Do(); err != nil {
		log.Printf("An error occured while downloading file: %s\n", err.Error())
	}

	fmt.Printf("Download Completed in %v seconds\n", time.Now().Sub(startTime).Seconds())
}
