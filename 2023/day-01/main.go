package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	val, err := calculateCalibrationVal()

	if err != nil {
		panic(err)
	}

	fmt.Println("Value", val)

}

type NumberHolder struct {
	first string
	last  string
}

func (nh *NumberHolder) register(s string) {
	if len(nh.first) == 0 {
		nh.first = s
	}

	nh.last = s
}

func (nh *NumberHolder) reset() {
    nh.first = ""
    nh.last = ""
}

func calculateCalibrationVal() (int, error) {
	// f, err := os.OpenFile("mini_input.txt", os.O_RDONLY, 0644)
	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0644)

	if err != nil {
		return 0, err
	}

	defer f.Close()

	var s = bufio.NewReader(f)
	var val int
	var nh = &NumberHolder{
		first: "",
		last:  "",
	}

	for tok, err := s.ReadByte(); err == nil; tok, err = s.ReadByte() {
		switch tok {
		case '\n':
            fmt.Println(nh.first, nh.last)
			num, err := strconv.Atoi(fmt.Sprintf("%s%s", nh.first, nh.last))
			if err != nil {
				return 0, nil
			}

			val += num
            nh.reset()
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			nh.register(string(tok))
		case 'o':
			if c, err := s.Peek(2); err == nil && string(c) == "ne" {
				nh.register("1")
			}
		case 't':
			if c, err := s.Peek(2); err == nil && string(c) == "wo" {
				nh.register("2")
			} else if c, err := s.Peek(4); err == nil && string(c) == "hree" {
				nh.register("3")
			}
		case 'f':
			if c, err := s.Peek(3); err == nil {
				if string(c) == "ive" {
					nh.register("5")
				} else if string(c) == "our" {
					nh.register("4")
				}
			}
		case 's':
			if c, err := s.Peek(2); err == nil && string(c) == "ix" {
				nh.register("6")
			} else if c, err := s.Peek(4); err == nil && string(c) == "even" {
				nh.register("7")
			}
		case 'e':
			if c, err := s.Peek(4); err == nil && string(c) == "ight" {
				nh.register("8")
			}
		case 'n':
			if c, err := s.Peek(3); err == nil && string(c) == "ine" {
				nh.register("9")
			}
		case 'z':
			if c, err := s.Peek(3); err == nil && string(c) == "ero" {
				nh.register("0")
			}
		}
	}

	return val, nil
}
