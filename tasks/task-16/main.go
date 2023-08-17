package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Slice []string
type Map map[string]int

type alias interface {
	ToString()
}

func main() {
	var s Slice = []string{"Hello", "World"}
	var m Map = make(map[string]int)
	m["bangalore-1"] = 560086
	m["bangalore-2"] = 560096
	var IAlias []alias = []alias{s, m}
	for _, a := range IAlias {
		a.ToString()
	}
}

func (s Slice) ToString() {
	output := ""
	for _, v := range s {
		output += v + ","
	}
	output = strings.TrimSuffix(output, ",")
	fmt.Println(output)
}

func (s Map) ToString() {
	output := ""
	for k, v := range s {
		output += "Key:" + k + "," + "Value:" + strconv.Itoa(v) + ";"
	}
	output = strings.TrimSuffix(output, ";")
	fmt.Println(output)
}
