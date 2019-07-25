package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	pwd, _ := os.Getwd()
	path := pwd + "/assets/"
	files, _ := ioutil.ReadDir(path)

	final := counterAllFile(files, path)

	var finalMap = make(map[string]int)
	for e := range final {
		for k, v := range e {
			size, ok := finalMap[k]
			if ok {
				finalMap[k] = size + v
			} else {
				finalMap[k] = v
			}
		}
	}

	fmt.Println(finalMap)
}

func counterAllFile(files []os.FileInfo, path string) chan map[string]int {
	out := make(chan map[string]int, 100)
	var wg sync.WaitGroup

	output := func(file string) {
		for n := range counterByFile(file) {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(files))

	for _, f := range files {
		go output(filepath.Join(path, f.Name()))
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
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
