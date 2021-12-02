package main

import (
	"flag"
	"github.com/gravityperson/parser-drive2ru-images-downloader/internal/downloader"
	"github.com/gravityperson/parser-drive2ru-images-downloader/internal/message"
	"github.com/gravityperson/parser-drive2ru-images-downloader/internal/model"
	"github.com/sirupsen/logrus"
	"log"
)

var (
	distFolder string
	logLevel   string
	postUrl    string
)

func init() {
	flag.StringVar(&distFolder, "dist-folder", "dist", "the results folder")
	flag.StringVar(&logLevel, "log-level", "info", "an application log-level")
	flag.StringVar(&postUrl, "post-url", "", "a post of drive2.ru post")
}

func main() {
	flag.Parse()

	logger := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf(message.ErrLogLevel, logLevel)
	}

	textFmt := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logger.SetLevel(level)
	logger.SetFormatter(textFmt)

	config := &model.Config{
		Logger:     logger,
		DistFolder: distFolder,
	}

	if postUrl == "" {
		logger.Fatal(message.ErrPostUrl)
	}

	if err := downloader.Start(config, postUrl); err != nil {
		logger.Fatalf(message.ErrGeneric, err)
	}
}
