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
	#fmt.Println("HTML: ", output)
	return string(output), nil
}

func processFile(path string, info fs.FileInfo, err error) error {
	parts := strings.Split(path, "/")

    // Verifica se há pelo menos 3 partes (0-indexed)
    if len(parts) >= 3 {
        // Obtém o elemento na posição 2
        platform := parts[2]
        product  := parts[3]
        fmt.Println("Segunda parte do caminho:", platform)
    } else {
        fmt.Println("O caminho não contém pelo menos 3 partes.")
    }

	if strings.HasSuffix(path, ".md") {
		input, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		htmlContent, err := convertMarkdownToHTML(string(input))
		if err != nil {
			return err
		}

		newPath := strings.TrimSuffix(path, ".md") + "-" + platform + "-" + product + ".html"
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

