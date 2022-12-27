package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func FindReplaceFile(src, dst, old, new string) (occ int, lines []int, err error) {
	create, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}

	writer := bufio.NewWriter(create)
	defer writer.Flush()

	file, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()
	i := 0
	for scanner.Scan() {
		i++
		t := scanner.Text()
		found, res, occL := ProcessLine(t, old, new)
		_, err := fmt.Fprintln(writer, res)
		if err != nil {
			return occ, lines, err
		}

		if !found {
			continue
		}
		occ += occL
		lines = append(lines, i)
	}
	return occ, lines, nil
}

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	found = strings.Contains(line, old)
	occ = strings.Count(line, old)
	res = strings.Replace(line, old, new, occ)
	return found, res, occ
}

func main() {
	occ, lines, err := FindReplaceFile("wikigo.txt", "wikigo-new.txt",
		"Go", "python")
	if err != nil {
		log.Fatalf("message: %s", err.Error())
	}
	println("===== Summary =====")
	fmt.Printf("Number of occurrences of Go : %d\n", occ)
	fmt.Printf("Number of lines: %d\n", len(lines))
	fmt.Printf("Lines: %d\n", lines)
	println("===== End of Summary =====")
}
