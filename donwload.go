package download

import (
	"log"
	"time"
)

func Download(url, targetPath string, totalSection int) {
	startTime := time.Now()
	d := DownloadModel{
		url:          url,
		targetPath:   targetPath,
		totalSection: totalSection,
	}

	if err := d.do(); err != nil {
		log.Printf("An error occured while downloading file: %s\n", err.Error())
	}

	log.Printf("Download Completed in %v seconds\n", time.Now().Sub(startTime).Seconds())
}
