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

/*
Run:
go run fetchall.go https://www.Google.com https://www.Youtube.com https://www.Facebook.com https://www.Baidu.com https://www.Wikipedia.org https://www.Yahoo.com https://www.Google.co.in https://www.Amazon.com https://www.Qq.com https://www.Google.co.jp https://www.Taobao.com https://www.Live.com https://www.Vk.com https://www.Twitter.com https://www.Instagram.com

Result:
0.49s   10408 https://www.Google.co.jp
0.60s   12397 https://www.Google.co.in
0.71s  129806 https://www.Facebook.com
0.82s   10501 https://www.Google.com
1.17s       0 https://www.Instagram.com
1.35s   81473 https://www.Wikipedia.org
1.45s  413701 https://www.Yahoo.com
1.89s  564616 https://www.Youtube.com
1.90s  303271 https://www.Twitter.com
2.15s     173 https://www.Baidu.com
2.20s   15921 https://www.Live.com
2.67s  226112 https://www.Amazon.com
8.62s   84485 https://www.Taobao.com
11.17s    6538 https://www.Vk.com
Get https://www.Qq.com: dial tcp 103.7.30.123:443: i/o timeout
30.01s elapsed
*/
