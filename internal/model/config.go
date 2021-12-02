package model

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Logger     *logrus.Logger
	DistFolder string
}
