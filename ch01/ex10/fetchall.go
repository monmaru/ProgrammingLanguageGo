package main

import (
	"fmt"
	"io"
	"io/ioutil"
	mylog "log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		log(<-ch)
	}

	log(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nBytes, url)
}

func log(message string) {
	logfile, err := os.OpenFile("./fetchall.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open fetchall.log:" + err.Error())
	}
	defer logfile.Close()

	mylog.SetOutput(io.MultiWriter(logfile, os.Stdout))
	mylog.SetFlags(mylog.Ldate | mylog.Ltime)
	mylog.Println("| ", message)
}
