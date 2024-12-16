package main

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed input.txt
var input string

type file struct {
	index  int
	fileId int
	length int
}
type gap struct {
	index  int
	length int
}

func readInput(input string) []int {
	length := len(input)
	intList := make([]int, length)
	for i, char := range input {
		intList[i] = int(char) - int('0')
	}
	return intList
}

func checksum(intList []int) int {
	length := len(intList)
	sum := 0
	expIndex := 0

	j := length - 1
	if j%2 != 0 {
		j -= 1
	}

	for i := 0; i <= j; i++ {
		if i%2 == 0 {
			fileId := i / 2
			val := intList[i]
			indexSum := val*expIndex + (val*(val-1))/2

			sum += fileId * indexSum
			expIndex += val
		} else {
			block := intList[i]
			for k := 0; k < block; k++ {
				for intList[j] <= 0 {
					j -= 2
				}

				intList[j] -= 1
				intList[i] -= 1
				sum += expIndex * (j / 2)
				expIndex += 1
			}
		}
	}

	return sum
}

func parseFileSystem(blocks []int) ([]file, []gap) {
	files := make([]file, 0)
	gaps := make([]gap, 0)
	length := len(blocks)
	expandedIndex := 0

	for i := 0; i < length; i++ {
		if i%2 == 0 {
			files = append(files, file{expandedIndex, i / 2, blocks[i]})
			expandedIndex += blocks[i]
		} else {
			gaps = append(gaps, gap{expandedIndex, blocks[i]})
			expandedIndex += blocks[i]
		}
	}
	return files, gaps
}

func moveFiles(files []file, gaps []gap) ([]file, []gap) {
	for j := len(files) - 1; j > 0; j-- {
		for i, gap := range gaps {
			if gap.index > files[j].index {
				break
			}
			fileSize := files[j].length
			if fileSize <= gap.length {
				files[j].index = gap.index
				gaps[i].index += fileSize
				gaps[i].length -= fileSize

				if gaps[i].length == 0 {
					gaps = append(gaps[:i], gaps[i+1:]...)
				}
				break
			}
		}
	}

	return files, gaps
}

func filesChecksum(files []file) int {
	sum := 0
	for _, file := range files {
		sum += file.fileId * (file.index*file.length + (file.length*(file.length-1))/2)
	}
	return sum
}

func part1(input string) int {
	blocks := readInput(input)
	return checksum(blocks)
}

func part2(input string) int {
	blocks := readInput(input)
	files, gaps := parseFileSystem(blocks)
	files, gaps = moveFiles(files, gaps)
	return filesChecksum(files)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part tu run? 1 or 2")
	flag.Parse()

	res := 0
	if part == 1 {
		res = part1(input)
	} else {
		res = part2(input)
	}

	fmt.Println("Result: ", res)
}
