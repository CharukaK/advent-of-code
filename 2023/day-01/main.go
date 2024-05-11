package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	val, err := calculateCalibrationVal()

	if err != nil {
		panic(err)
	}

	fmt.Println("Value", *val)

}

func calculateCalibrationVal() (*int, error) {
	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	fsInfo, err := f.Stat()

	if err != nil {
		return nil, err
	}

	data := make([]byte, fsInfo.Size())

	if _, err = f.Read(data); err != nil {
		return nil, err
	}

	var val int

	first, last := "", ""

	for _, c := range data {
		switch c {
		case '\n':
			num, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))

			if err != nil {
				return nil, err
			}

			val += num
			first = ""
			last = ""

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if len(first) == 0 {
				first = string(c)
			}

			last = string(c)
		}

	}

	return &val, nil
}
