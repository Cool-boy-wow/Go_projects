package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var first_name string
	var second_name string

	flag.StringVar(&first_name, "first_name", "", "set first_name")
	flag.StringVar(&second_name, "second_name", "", "set second_name")

	flag.Parse()

	result, _ := os.Create("result.txt")

	if first_name != "" && second_name != "" {
		file_1, _ := os.Create(first_name)
		file_2, _ := os.Create(second_name)
		defer file_1.Close()
		defer file_2.Close()
		file_1.WriteString("Text of first file")
		file_2.WriteString("Text of second file")
		s1, _ := ioutil.ReadFile(first_name)
		s2, _ := ioutil.ReadFile(second_name)
		s3 := []string{string(s1), string(s2)}
		s4 := strings.Join(s3, "\n")
		fmt.Println(s4)
		result.WriteString(string(s4))
	} else if second_name == "" {
		file_1, _ := os.Create(first_name)
		defer file_1.Close()
		file_1.WriteString("Text of first file")
		s1, _ := ioutil.ReadFile(first_name)
		result.WriteString(string(s1))
	}
}
