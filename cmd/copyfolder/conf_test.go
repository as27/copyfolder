package main

import (
	"reflect"
	"testing"
)

func TestConf(t *testing.T) {
	var exp = Options{
		Folders: []SrcDst{
			SrcDst{SrcFolder: "path/to/src", DstFolder: "path/to/dst"},
			SrcDst{SrcFolder: "another/src", DstFolder: "another/dst"},
		}}
	c := loadConf("test.yaml")
	if !reflect.DeepEqual(exp, *c) {

		t.Fail()
	}
}
