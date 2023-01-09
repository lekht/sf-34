package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filenameRead := "./tasks.txt"
	f1, err := os.OpenFile(filenameRead, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fileReader := bufio.NewReader(f1)
	filenameWrite := "./results.txt"
	_ = os.Remove(filenameWrite)
	f2, err := os.OpenFile(filenameWrite, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		// fmt.Println("NEW LINE:", string(line))
		s, ok := regexFilter(string(line))
		if !ok {
			continue
		}
		fmt.Println(s)
		_, _ = f2.WriteString(s + "\n")
	}
}

func regexFilter(s string) (string, bool) {
	re1, err := regexp.Compile(`[0-9]+[+-][0-9]+[=]`)
	if err != nil {
		fmt.Println("error is ", err)
	}
	ss := re1.FindString(s)
	if ss == "" {
		return "", false
	}
	re2, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		fmt.Println("error is ", err)
	}
	operands := re2.FindAllString(s, 2)
	var ints []int
	for _, n := range operands {
		i, _ := strconv.Atoi(n)
		ints = append(ints, i)
	}

	re3, err := regexp.Compile(`[+-]`)
	if err != nil {
		fmt.Println("error is ", err)
	}
	operator := re3.FindString(s)
	result, err := calc(operator, ints)
	if err != nil {
		fmt.Println("error is ", err)
	}

	return ss + result, true
}

func calc(s string, i []int) (string, error) {
	switch s {
	case "+":
		return strconv.Itoa(i[0] + i[1]), nil
	case "-":
		return strconv.Itoa(i[0] - i[1]), nil
	}
	return "", fmt.Errorf("unknown operator")
}
