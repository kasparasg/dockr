package queue

var WorkQueue = make(chan Job)

type Job struct {
	Number int
}
