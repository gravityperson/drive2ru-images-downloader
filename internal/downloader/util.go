package downloader

import (
	"fmt"
	"github.com/gravityperson/parser-drive2ru-images-downloader/internal/message"
	"github.com/gravityperson/parser-drive2ru-images-downloader/internal/model"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func prepareDistFolder(config *model.Config) {
	config.Logger.Debugf(message.MsgGeneric, "preparing dist folder")
	entry, err := os.Stat(config.DistFolder)

	if !os.IsNotExist(err) && entry.IsDir() {
		config.Logger.Debugf(message.MsgGeneric, "dist folder exists; clearing")
		if err := os.RemoveAll(config.DistFolder); err != nil {
			config.Logger.Fatalf(message.ErrGeneric, err)
		}

		if err := os.Mkdir(config.DistFolder, os.ModeDir); err != nil {
			config.Logger.Errorf(message.ErrGeneric, err)
		}
	}

	if os.IsNotExist(err) {
		config.Logger.Debugf(message.MsgGeneric, "dist folder doesn't exist; creating")
		if err := os.Mkdir(config.DistFolder, os.ModeDir); err != nil {
			config.Logger.Errorf(message.ErrGeneric, err)
		}
	}
}

func printConfig(config *model.Config, postUrl string) {
	config.Logger.WithFields(logrus.Fields{
		"log-level": config.Logger.Level,
		"post-url":  postUrl,
	}).Debugf(message.MsgGeneric, "application starting")
}

func buildFilePath(distFolder string, fileName string) string {
	return fmt.Sprintf("%s/%s", distFolder, fileName)
}

func extractFileName(fileUrl string) string {
	parts := strings.Split(fileUrl, "/")
	return parts[len(parts)-1]
}
