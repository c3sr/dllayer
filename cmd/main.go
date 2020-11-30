// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/c3sr/caffe"
	"github.com/c3sr/config"
	"github.com/c3sr/dlframework/framework/cmd"
	dllayercmd "github.com/c3sr/dllayer/cmd"
	"github.com/sirupsen/logrus"

	_ "github.com/c3sr/logger/hooks"
	_ "github.com/c3sr/tracer/all"
)

var (
	log *logrus.Entry = logrus.New().WithField("pkg", "dlframework/framework/cmd/evaluate")
)

func main() {
	config.AfterInit(func() {
		log = logrus.New().WithField("pkg", "dlframework/framework/cmd/evaluate")
	})
	cmd.Init()

	caffe.Register()
	dllayercmd.Framework = caffe.FrameworkManifest

	if err := dllayercmd.FlopsInfoCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
