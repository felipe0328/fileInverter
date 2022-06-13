package main

type FileReadedLine struct {
	Data   chan [2]interface{}
	Error  chan error
	IsOpen bool
}

func NewFileReadedLine() *FileReadedLine {
	return &FileReadedLine{
		Data:   make(chan [2]interface{}),
		Error:  make(chan error),
		IsOpen: true,
	}
}

func (fileReadedLine *FileReadedLine) WriteDataToMap(fileMap *FileMap) {
	for fileReadedLine.IsOpen {
		select {
		case data := <-fileReadedLine.Data:
			fileMap.Lock()
			mapIndex, _ := data[0].(int)
			mapData, _ := data[1].(string)
			fileMap.data[mapIndex] = mapData
			fileMap.Unlock()

		case err := <-fileReadedLine.Error:
			check(err)
		}
	}
}

func (fileReadedLine *FileReadedLine) Close() {
	close(fileReadedLine.Data)
	close(fileReadedLine.Error)
	fileReadedLine.IsOpen = false
}
