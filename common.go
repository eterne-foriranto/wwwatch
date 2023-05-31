package main

import (
	"fmt"
	"os"
	"strconv"
)

func readIntFromFile(fileName string) int {
	content, err := os.ReadFile(fileName)

	if err != nil {
	}

	res, err2 := strconv.Atoi(string(content))

	if err2 != nil {
	}

	return res
}

func writeIntToFile(fileName string, entireInt int) {
	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	_, err = f.WriteString(strconv.Itoa(entireInt))

	if err != nil {
		fmt.Println(err)
	}
}
