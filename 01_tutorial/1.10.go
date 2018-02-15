// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strconv"
)

func main() {
	start := time.Now()
	logFileTime := time.Now().UTC().UnixNano()
	logfile,err := os.Create(strconv.FormatInt(logFileTime, 10)+".log")
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create: %v\n", err)
	}
	
	ch := make(chan string)
	
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		//fmt.Println(<-ch) // receive from channel ch
		logfile.WriteString(<-ch + "\n")
	}
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Fprintf(logfile, "%.2fs elapsed\n", time.Since(start).Seconds())
	
	logfile.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
