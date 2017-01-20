package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"reflect"

	"github.com/as27/copyfolder"
	"github.com/as27/copyfolder/fs"
)

const (
	version  = "0.7.0"
	confFile = "copyfolder.yaml"
)

func main() {
	vFlag := flag.Bool("version", false, "Prints out the version")
	flag.Parse()
	if *vFlag {
		fmt.Printf("Version: %s", version)
		os.Exit(0)
	}
	// Check if there is a conf file, when not an default/Example
	// file is created.
	if !fileExists(confFile) {
		err := touchConf(confFile)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Default/Example conf file created.")
		os.Exit(0)
	}

	conf := loadConf(confFile)
	// Check if conf is still the conf file
	if reflect.DeepEqual(*conf, defaultOptions) {
		fmt.Printf("Please change the file %s\nThere are still the example values inside that file.", confFile)
		os.Exit(1)
	}
	for _, sd := range conf.Folders {
		folderCopier := fs.NewFolderCopier(sd.SrcFolder, sd.DstFolder)
		err := copyfolder.Copy(folderCopier)
		if err != nil {
			log.Println(err)
		}
	}
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
