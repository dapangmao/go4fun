

```go
package main

import (
	"time"
	"fmt"
	"os/exec"
	"flag"
)

var HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK int

const INTERVAL_PERIOD = 24 * time.Hour

type jobTicker struct {
	t *time.Timer
}

func getNextTickDuration() time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), HOUR_TO_TICK, MINUTE_TO_TICK,
		SECOND_TO_TICK, 0, time.UTC)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	return nextTick.Sub(time.Now())
}

func NewJobTicker() jobTicker {
	fmt.Println("New tick has been set")
	return jobTicker{time.NewTimer(getNextTickDuration())}
}

func (self jobTicker) updateJobTicker() {
	fmt.Println("New tick has been set")
	self.t.Reset(getNextTickDuration())
}

func main() {
	flag.IntVar(&HOUR_TO_TICK, "hour", 0, "the UTC hour to run")
	flag.IntVar(&MINUTE_TO_TICK, "minute", 0, "the minute to run")
	flag.IntVar(&SECOND_TO_TICK, "second", 0, "the second to run")
	var project string
	flag.StringVar(&project, "project", "cts", "the project to run")

	flag.Parse()

	jt := NewJobTicker()
	for {
		<-jt.t.C
		go func() {
			path, err := exec.LookPath("python")
			if err != nil {
				fmt.Println(err)
				fmt.Println("Cannot find the Python")
				return
			}
			fmt.Println(path)
			cmd := exec.Command(path, "pipeline.py", "run",
				"--settings=settings.qa_docker", "--projects=" + project)
			err2 := cmd.Run()
			if err2 != nil {
				fmt.Println(err2)
			} else {
				fmt.Println("The pipeline is finished within a Go routine")
			}
		}()
		fmt.Println(time.Now(), "Just fired up an ETL pipeline")
		jt.updateJobTicker()
	}
}

```
