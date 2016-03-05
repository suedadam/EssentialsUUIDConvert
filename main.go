package main

import (
	"fmt"
	"github.com/smallfish/simpleyaml"
	"path/filepath"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
    searchDir := os.Args[1]

    fileList := []string{}
    _ = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        fileList = append(fileList, path)
        return nil
    })

    for _, file := range fileList {
	    if strings.Contains(file, "yml") {  	
	    	parse(file)
	    }
    }
}

func changename(oldname string, UUID string) {
	newname := fmt.Sprintf("%s.yml", UUID)
	err := os.Rename(oldname, newname)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func parse(file string) {
	filename := file
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	yaml, err := simpleyaml.NewYaml(source)
	if err != nil {
		panic(err)
	}
	uuid, err := yaml.Get("uuid").String()
	if err != nil {
		panic(err)
	}
	changename(filename, uuid)
}