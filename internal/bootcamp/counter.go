package bootcamp

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func Handle() {
	pwd, _ := os.Getwd()
	defaultPath := pwd + "/assets/"

	path := flag.String("folder", defaultPath, "The folder")
	flag.Parse()

	files, _ := ioutil.ReadDir(*path)

	pool := NewFixedThreadPool(10, 1000)

	var futures []Future
	for _, f := range files {
		future := pool.Submit(&TaskCounter{FilePath: filepath.Join(*path, f.Name())})
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

func merge(futures []Future) <-chan map[string]int {
	var wg sync.WaitGroup
	out := make(chan map[string]int)

	output := func(f Future) {
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
