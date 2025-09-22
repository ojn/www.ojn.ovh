package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	matched, err := filepath.Match("index.html", fi.Name())

	if err != nil {
		panic(err)
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(read))
		fmt.Println(path)

		newContents := strings.Replace(string(read), "<!-- Book generated using mdBook -->", 
		"<link rel=\"alternate\" type=\"application/rss+xml\" title=\"ojn website feed\" href=\"https://www.ojn.ovh/feed\">", -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}

func main() {
	err := filepath.Walk("./docs/", visit)
	if err != nil {
		panic(err)
	}
}
