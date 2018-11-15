package goruntine

import (
	"os"
	"time"
	"errors"
	"os/signal"
	"log"
)

type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []func(int)
}


var ErrorTimeout = errors.New("received Timeout")

var ErrorInterrupt = errors.New("received Interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks =append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id,task := range r.tasks {
		if r.gotInterrupt() {
			return ErrorInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
		case <- r.interrupt:
			signal.Stop(r.interrupt)
		return true

	default:
		return false
	}
}


func (r *Runner) Start() error{
	signal.Notify(r.interrupt,os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	 	case err := <- r.complete:
	 		return err
	 	case <- r.timeout:
	 		return ErrorTimeout

	}
}

var timeout = 4 * time.Second

func RunnerMain() {
	log.Println("String work")

	r := New(timeout)

	r.Add(creatTask(),creatTask(),creatTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrorTimeout:
			log.Println("timeout")
			os.Exit(1)
		case ErrorInterrupt:
			log.Println("interrupt")
			os.Exit(2)
		}

	}

	log.Println("Process ended")


}

func creatTask() func(int){
	return func(id int) {
		log.Printf("Processor - Task #%d.",id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}