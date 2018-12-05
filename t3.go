package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/hoisie/mustache"
)

var (
	dataPath     = flag.String("d", "", "json data file path")
	templatePath = flag.String("t", "", "mustache template file path")
)

func main() {
	var data interface{}
	var jsonFile io.ReadCloser
	var err error
	var rawJSON []byte
	var templateFile io.ReadCloser
	var template []byte

	// templatePath := "template2.mustache"
	// dataPath := "data2.json"

	flag.Parse()

	if templateFile, err = os.Open(*templatePath); err != nil {
		fmt.Println("Could not open template file")
		os.Exit(1)
	}
	defer templateFile.Close()
	if template, err = ioutil.ReadAll(templateFile); err != nil {
		fmt.Println("Could not read template file")
		os.Exit(1)
	}

	if jsonFile, err = os.Open(*dataPath); err != nil {
		fmt.Println("Could not open data file")
		os.Exit(1)
	}
	defer jsonFile.Close()
	if rawJSON, err = ioutil.ReadAll(jsonFile); err != nil {
		fmt.Println("Could not read data file")
		os.Exit(1)
	}
	json.Unmarshal([]byte(rawJSON), &data)

	rendered := mustache.Render(string(template), data)
	fmt.Printf(rendered)
}
