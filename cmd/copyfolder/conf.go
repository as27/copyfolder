package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type SrcDst struct {
	SrcFolder string `yaml:"src"`
	DstFolder string `yaml:"dst"`
}

type Conf struct {
	Folders []SrcDst `yaml:"folders"`
}

var defaultOptions = Conf{
	[]SrcDst{
		SrcDst{SrcFolder: "path/to/src", DstFolder: "path/to/dst"},
		SrcDst{SrcFolder: "another/src", DstFolder: "another/dst"},
	},
}

func touchConf(fp string) error {
	if fileExists(fp) {
		return nil
	}
	b, err := yaml.Marshal(defaultOptions)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fp, b, 0777)
	return err
}

func loadConf(fp string) *Conf {
	b, _ := ioutil.ReadFile(fp)
	c := &Conf{}
	yaml.Unmarshal(b, c)
	return c
}
