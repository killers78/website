package main

import (
	"io/fs"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"strings"
	"github.com/russross/blackfriday/v2"
)

func convertMarkdownToHTML(input string) (string, error) {
	output := blackfriday.Run([]byte(input))
	fmt.Println("HTML: ", output)
	return string(output), nil
}

func processFile(path string, info fs.FileInfo, err error) error {
	if strings.HasSuffix(path, ".md") {
		input, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		htmlContent, err := convertMarkdownToHTML(string(input))
		if err != nil {
			return err
		}

		newPath := strings.TrimSuffix(path, ".md") + ".html"
		err = ioutil.WriteFile(newPath, []byte(htmlContent), info.Mode())
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	root := "./platforms" // Defina o diretório raiz conforme necessário
	err := filepath.Walk(root, processFile)
	if err != nil {
		panic(err)
	}
}

