package downloader

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gravityperson/parser-drive2ru-images-downloader/internal/message"
	"github.com/gravityperson/parser-drive2ru-images-downloader/internal/model"
	"io"
	"net/http"
	"os"
	"sync"
)

func Start(config *model.Config, postUrl string) error {
	// printing config in case it's "debug" log-level
	printConfig(config, postUrl)

	// start parsing
	document, err := htmlquery.LoadURL(postUrl)
	if err != nil {
		config.Logger.Fatalf(message.ErrGeneric, err)
	}

	images := htmlquery.Find(document, "//a[contains(@class, 'c-pic-zoom')]/@href")
	imagesCount := len(images)
	sources := make([]string, 0, imagesCount)
	config.Logger.Infof(message.MsgGeneric, fmt.Sprintf("%d images found", imagesCount))

	for _, img := range images {
		href := htmlquery.SelectAttr(img, "href")
		sources = append(sources, href)
	}

	prepareDistFolder(config)
	downloadImages(config, &sources)

	config.Logger.Infof(message.MsgGeneric, "images downloaded successfully")
	// end parsing

	return nil
}

func downloadImages(config *model.Config, images *[]string) {
	config.Logger.Debugf(message.MsgGeneric, "starting to download images")
	wg := &sync.WaitGroup{}

	for _, image := range *images {
		wg.Add(1)

		go func(image string, wg *sync.WaitGroup) {
			defer wg.Done()

			response, err := http.Get(image)
			defer response.Body.Close()
			if err != nil {
				config.Logger.Errorf(message.ErrGeneric, err)
			}

			if response.StatusCode != 200 {
				config.Logger.Fatalf(message.ErrGeneric, "file downloading error")
			}

			fileName := extractFileName(image)
			filePath := buildFilePath(config.DistFolder, fileName)
			file, err := os.Create(filePath)
			defer file.Close()
			if err != nil {
				config.Logger.Fatalf(message.ErrGeneric, err)
			}

			if _, err := io.Copy(file, response.Body); err != nil {
				config.Logger.Fatalf(message.ErrGeneric, err)
			}
		}(image, wg)

		wg.Wait()
	}
}
