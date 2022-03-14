package download

import (
	"fmt"
	"log"
	"time"
)

// Donwload - Helps you download the torrent file.
func Download(url, targetPath string, totalSection int) (string, error) {
	startTime := time.Now()
	d := DownloadModel{
		url:          url,
		targetPath:   targetPath,
		totalSection: totalSection,
	}

	if err := d.do(); err != nil {
		log.Printf("An error occured while downloading file: %s\n", err.Error())
	}

	output := fmt.Sprintf("Download Completed in %v seconds\n", time.Now().Sub(startTime).Seconds())

	return output, nil
}
