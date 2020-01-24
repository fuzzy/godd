package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/gosuri/uiprogress"
)

type GoddOpts struct {
	// these should be set at init, and are only used
	// internally.
	Quiet     bool
	Progress  bool
	BlockSize int
	Count     int
	Skip      int
	Seek      int
	Size      int64
	StartTime time.Time
	EndTime   time.Time
}

func mover(wg *sync.WaitGroup, st int, bs int, f_in GoddReader, f_out GoddWriter) {
	defer wg.Done()
	buffer := make([]byte, bs)
	fmt.Printf("st: %d\nbs: %d\nbuff: %d", st, bs, len(buffer))
	bar := uiprogress.AddBar(st).AppendCompleted().PrependElapsed()

	for bar.Incr() {
		_, err := f_in.Read(buffer)
		if err != nil {
			return
		}

		_, err = f_out.Write(buffer)
		if err != nil {
			return
		}
	}

}

func main() {
	opts, input, output := ArgParser()
	s1 := int(opts.Size / int64(opts.BlockSize))
	s2 := int(math.Remainder(float64(opts.Size), float64(opts.BlockSize)))

	uiprogress.Start()
	var wg sync.WaitGroup

	wg.Add(1)
	go mover(&wg, (s1 + s2), opts.BlockSize, input, output)

	time.Sleep(time.Second * 1)
	wg.Wait()
	uiprogress.Stop()
}
