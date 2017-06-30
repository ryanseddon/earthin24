package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	frames := make([]int, 24*6)
	now := time.Now()
	// Change the date to midday AEST in UTC time
	t := time.Date(now.Year(), now.Month(), now.Day(), 2, 00, 00, 0, time.UTC)
	first := t.Add(-24 * time.Hour)
	setupLogs()

	for i, _ := range frames {
		wg.Add(1)
		go getImage(framePath(first), i, &wg)
		first = first.Add(time.Duration(10) * time.Minute)
	}

	wg.Wait()
}

func framePath(time time.Time) string {
	url := "http://himawari8-dl.nict.go.jp/himawari8/img/D531106/1d/550"
	minute := time.Minute() - time.Minute()%10

	return fmt.Sprintf("%s/%d/%02d/%02d/%02d%02d00_0_0.png", url, time.Year(), time.Month(), time.Day(), time.Hour(), minute)
}

func getImage(url string, name int, wg *sync.WaitGroup) {
	defer wg.Done()
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		//open a file for writing
		file, err := os.Create(fmt.Sprintf("%04d.png", name))
		if err != nil {
			log.Fatal(err)
		}

		// Generating md5 hashes for frames
		hash := md5.New()
		var writers []io.Writer
		writers = append(writers, hash, file)

		defer file.Close()
		dest := io.MultiWriter(writers...)

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(dest, response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// If an image matches the error file hash remove it
		// TODO: Can I abort the file write instead?
		errorHash := "b697574875d3b8eb5dd80e9b2bc9c749"
		imgHash := hex.EncodeToString(hash.Sum(nil))

		if imgHash == errorHash {
			os.Remove(fmt.Sprintf("%04d.png", name))
		}
	}
}

//func compareHashes(hash *hash.Hash, name int) {
//// If an image matches the error file hash remove it
//// TODO: Can I abort the file write instead?
//errorHash := "b697574875d3b8eb5dd80e9b2bc9c749"
//imgHash := hex.EncodeToString(hash.Sum(nil))

//if imgHash == errorHash {
//os.Remove(fmt.Sprintf("%04d.jpg", name))
//}
//}

func setupLogs() {
	//log url requested for later debugging
	logs, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logs.Close()
	log.SetOutput(logs)
}