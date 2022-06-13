package main

import (
	"bufio"
	"sync"
)

type FileMap struct {
	data map[int]string
	sync.Mutex
}

func newFileMap() *FileMap {
	return &FileMap{data: make(map[int]string)}
}

// Using concurrency and goroutines to avoid nested for loops and make the process faster
func (outputMap *FileMap) readFileScannerAndAddsToMap(fileScanner *bufio.Scanner, currentIndex int, fileReadedLine *FileReadedLine, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	couldScan := fileScanner.Scan()

	if !couldScan {
		wg.Done()
		return
	}
	// While reading the next file in a goroutine this thread works just in invert the text
	go outputMap.readFileScannerAndAddsToMap(fileScanner, currentIndex+1, fileReadedLine, wg)

	fileReadedLine.Data <- [2]interface{}{currentIndex, invertText(fileScanner.Text())}
}

func invertText(inputText string) string {
	textRunes := []rune(inputText)

	outputString := ""
	for runeIndex := len(textRunes) - 1; runeIndex >= 0; runeIndex-- {
		outputString += string(textRunes[runeIndex])
	}

	return outputString
}
