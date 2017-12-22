package cmd

import (
	"os"
	"path/filepath"

	"github.com/Unknwon/com"
	framework "github.com/rai-project/dlframework/framework/cmd"
)

func getSrcPath(importPath string) (appPath string) {
	paths := com.GetGOPATHs()
	for _, p := range paths {
		d := filepath.Join(p, "src", importPath)
		if com.IsExist(d) {
			appPath = d
			break
		}
	}

	if len(appPath) == 0 {
		appPath = filepath.Join(goPath, "src", importPath)
	}

	return appPath
}

func forallmodels(run func() error) error {

	if modelName != "all" {
		return run()
	}

	outputDirectory := outputFileName
	if !com.IsDir(outputDirectory) {
		os.MkdirAll(outputDirectory, os.ModePerm)
	}
	for _, model := range framework.DefaultEvaulationModels {
		modelName, modelVersion = framework.ParseModelName(model)
		outputFileName = filepath.Join(outputDirectory, model+"."+outputFileExtension)
		run()
	}
	return nil
}
