package main

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
	"time"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	path, err := filepath.Abs("template.html")
	if err != nil {
		log.Fatal(err)
	}
	path = "file://" + path

	ln, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port 7777")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("socket error: " + err.Error())
		}
		go handleConnection(ctx, path, conn)
	}
}

func handleConnection(ctx context.Context, templatePath string, conn net.Conn) {
	start := time.Now()
	var buff []byte
	if err := chromedp.Run(ctx, fullScreenshot(templatePath, &buff)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./screenshot.pdf", buff, 0644); err != nil {
		log.Fatal(err)
	}
	_, err := conn.Write(buff)
	if err != nil {
		log.Println("error writing data to socket: " + err.Error())
	}
	end := time.Now()
	log.Printf("serving pdf. took %dms\n", end.Sub(start).Milliseconds())
	_ = conn.Close()
}

func fullScreenshot(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().Do(ctx)
			*res = buf
			return err
		}),
	}
}