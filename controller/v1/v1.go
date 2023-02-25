package v1

import (
	"log"
	"os"
	"phantom/util"
)

var zlog *util.Logger

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[controller] err[%s]", err.Error())
		os.Exit(1)
	}
}
