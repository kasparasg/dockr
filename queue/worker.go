package queue

import log "github.com/Sirupsen/logrus"

type Worker struct {
	Id          int
	Job         chan Job
	WorkerQueue chan chan Job
	QuitChan    chan bool
}

func NewWorker(id int, workerQueue chan chan Job) Worker {
	return Worker{
		Id:          id,
		Job:         make(chan Job),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Job

			select {
			case job := <-w.Job:
				log.Info("Received work ", w.Id, job.Number)
			case <-w.QuitChan:
				log.Info("Worker stopping")
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
