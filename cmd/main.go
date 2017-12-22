// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/rai-project/caffe"
	"github.com/rai-project/config"
	"github.com/rai-project/dlframework/framework/cmd"
	dllayercmd "github.com/rai-project/dllayer/cmd"
	"github.com/sirupsen/logrus"

	_ "github.com/rai-project/logger/hooks"
	_ "github.com/rai-project/tracer/all"
)

var (
	log *logrus.Entry = logrus.New().WithField("pkg", "dlframework/framework/cmd/evaluate")
)

func main() {
	config.AfterInit(func() {
		log = logrus.New().WithField("pkg", "dlframework/framework/cmd/evaluate")
	})
	cmd.Init()

	dllayercmd.Framework = caffe.FrameworkManifest

	if err := dllayercmd.FlopsInfoCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
