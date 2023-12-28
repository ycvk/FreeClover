package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		return
	}
	logrus.SetLevel(level)

	Log.SetOutput(os.Stdout)

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
