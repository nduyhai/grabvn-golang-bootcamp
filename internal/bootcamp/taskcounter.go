package bootcamp

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type TaskCounter struct {
	FilePath string
}

func (t *TaskCounter) Execute() map[string]int {
	var finalMap = make(map[string]int)

	for e := range counterByFile(t.FilePath) {
		for k, v := range e {
			size, ok := finalMap[k]
			if ok {
				finalMap[k] = size + v
			} else {
				finalMap[k] = v
			}
		}
	}
	return finalMap
}

func counterByFile(filePath string) chan map[string]int {
	result := make(chan map[string]int, 100)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	defer close(result)

	in := bufio.NewScanner(file)

	for in.Scan() {
		func(line string) {
			stats := make(map[string]int)

			for _, work := range strings.Fields(line) {
				size, ok := stats[work]
				if ok {
					stats[work] = size + 1
				} else {
					stats[work] = 1
				}
			}
			result <- stats

		}(in.Text())

	}

	if err := in.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
