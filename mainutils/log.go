package mainutils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)

func logAndExit(err error) {
	Logger.Output(4, err.Error())
	os.Exit(1)
}
