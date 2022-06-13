package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func FileInverter(filePath string) {
	file := getFileFromPath(filePath)
	defer file.Close()

	invertedFileMap := getInvertedFileMap(file)
	fmt.Println(invertedFileMap)
	writeInvertedTextToFile(filePath, invertedFileMap)
}

func getFileFromPath(filePath string) *os.File {
	file, err := os.Open(filePath)

	check(err)

	return file
}

func getInvertedFileMap(file *os.File) map[int]string {
	invertedFileMap := newFileMap()

	fileReadedLine := NewFileReadedLine()
	defer fileReadedLine.Close()

	fileScanner := bufio.NewScanner(file)

	fileIndex := 0

	var wg sync.WaitGroup
	wg.Add(1)
	go invertedFileMap.readFileScannerAndAddsToMap(fileScanner, fileIndex, fileReadedLine, &wg)
	go fileReadedLine.WriteDataToMap(invertedFileMap)
	wg.Wait()

	return invertedFileMap.data
}

func writeInvertedTextToFile(originalFilePath string, invertedTextMap map[int]string) {
	newFile, err := os.Create("inverted_" + originalFilePath)
	check(err)

	defer newFile.Close()

	for mapIndex := len(invertedTextMap) - 1; mapIndex >= 0; mapIndex-- {
		newFile.WriteString(fmt.Sprintf("%v\n", invertedTextMap[mapIndex]))
	}
	newFile.Sync()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
