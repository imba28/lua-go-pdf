package main

import "C"
import (
	"context"
	"errors"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"os"
	"path/filepath"
	"unsafe"
)

func main() {

}

//export render
func render(templatePath string) (unsafe.Pointer, C.size_t) {
	buf, err := renderPdf(templatePath)
	if err != nil {
		println(err.Error())
		return nil, 0
	}

	cdata := C.malloc(C.size_t(len(buf)))
	copy((*[1<<24]byte)(cdata)[0:len(buf)], buf)

	return cdata, C.size_t(len(buf))
}

func renderPdf(template string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	path, err := filepath.Abs(template)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(path)
	if os.IsNotExist(err) || info.IsDir() {
		return nil, errors.New("file not found")
	}

	path = "file://" + path

	var buff []byte
	if err := chromedp.Run(ctx, fullScreenshot(path, &buff)); err != nil {
		return nil, err
	}
	//if err := ioutil.WriteFile("./screenshot.pdf", buff, 0644); err != nil {
	//	return nil, err
	//}

	return buff, nil
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
