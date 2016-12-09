package queue

import (
	log "github.com/Sirupsen/logrus"
)

var WorkerQueue chan chan Job

func StartDispatcher(workers int) {
	WorkerQueue = make(chan chan Job, workers)

	for i := 0; i < workers; i++ {
		log.Info("Starting worker: ", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case job := <-WorkQueue:
				log.Info("Received a job")

				go func() {
					worker := <-WorkerQueue

					log.Info("Dispatching a job")

					worker <- job
				}()
			}
		}
	}()
}
