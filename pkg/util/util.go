package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func Readfile() string {
	file := "/Users/amarjeet/Downloads/iconic-episode-235405-01386c316000.json"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("cannot open file")
	}

	var jsonbuf = []byte(data)
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, jsonbuf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(buffer.String())
	return buffer.String()
}
