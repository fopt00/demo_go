package goroutine

import (
	"demo_go/goroutine/util"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func ConcurrentEntry() {
	//makeThumbnailsEntry()
	waitGroupEntry()
}

func A() (string, string, string) {
	time.Sleep(1 * time.Second)
	return "X0", "Y0", "Z0"
}

func B1(x *string) {
	time.Sleep(1 * time.Second)
	*x = "X"
}

func B2(y *string) {
	time.Sleep(2 * time.Second)
	*y = "Y"
}

func B3(z *string) {
	time.Sleep(3 * time.Second)
	*z = "Z"
}

func C(x, y, z string) string {
	time.Sleep(1 * time.Second)
	return x + y + z
}

func waitGroupEntry() {
	defer elapse()()
	x, y, z := A()
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		B1(&x)
		wg.Done()
	}()
	go func() {
		B2(&y)
		wg.Done()
	}()
	go func() {
		B3(&z)
		wg.Done()
	}()

	wg.Wait()
	p := C(x, y , z)
	fmt.Printf("product: %s\n", p)
}

func makeThumbnailsEntry() {
	var filenames = make([]string, 350)
	_ = filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		filenames = append(filenames, path)
		return nil
	})

	defer elapse()()
	err := makeThumbnails(filenames)
	if err != nil {
		fmt.Println(err)
	}
}

func makeThumbnails(filenames []string) error {
	errors := make(chan error, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			_, err := util.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

func elapse() func() {
	now := time.Now()
	return func() {
		fmt.Printf("elapse:%.3f ms\n", 1000*time.Since(now).Seconds())
	}
}
