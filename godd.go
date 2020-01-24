package main

import (
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

func main() {
	opts, input, output := ArgParser()
	s1 := int(opts.Size / int64(opts.BlockSize))
	inBar := uiprogress.AddBar(s1).AppendCompleted().PrependElapsed()

	uiprogress.Start()

	count := 0
	for inBar.Incr() {
		count++
		buffer := make([]byte, opts.BlockSize)
		_, _ = input.Read(buffer)
		_, _ = output.Write(buffer)
		if count == 10 {
			output.Sync()
			count = 0
		}
	}

	time.Sleep(time.Second * 1)
	uiprogress.Stop()
}
