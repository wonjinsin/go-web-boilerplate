package model

import (
	"log"
	"os"
	"pikachu/util"
)

var zlog *util.Logger

func init() {
	_, err := util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[model] err[%s]", err.Error())
		os.Exit(1)
	}
}
