package main

import (
	"fmt"
	"grabvn-golang-bootcamp/internal/bootcamp"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	pwd, _ := os.Getwd()
	path := pwd + "/assets/"
	files, _ := ioutil.ReadDir(path)

	pool := bootcamp.NewFixedThreadPool(10, 1000)

	var futures []bootcamp.Future
	for _, f := range files {
		future := pool.Submit(&bootcamp.TaskCounter{FilePath: filepath.Join(path, f.Name())})
		futures = append(futures, *future)
	}

	d := merge(futures)
	result := count(d)

	fmt.Println("Result:", result)
}

func count(in <-chan map[string]int) map[string]int {
	var finalMap = make(map[string]int)

	for e := range in {
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

func merge(futures []bootcamp.Future) <-chan map[string]int {
	var wg sync.WaitGroup
	out := make(chan map[string]int)

	output := func(f bootcamp.Future) {
		for n := range f.Data {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(futures))
	for _, f := range futures {
		go output(f)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
