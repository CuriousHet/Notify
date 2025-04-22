package notification

import (
	"log"
	"math/rand"
	"time"
)

type Dispatcher struct {
	queue   *Queue
	retries int
}

func NewDispatcher(queue *Queue, retries int) *Dispatcher {
	return &Dispatcher{
		queue:   queue,
		retries: retries,
	}
}

func (d *Dispatcher) Start(workerCount int) {
	for i := 1; i <= workerCount; i++ {
		go d.worker(i)
	}
}

func (d *Dispatcher) worker(id int) {
	for n := range d.queue.Dequeue() {
		d.processNotification(n)
	}
}

func (d *Dispatcher) processNotification(n Notification) {
	for attempt := 1; attempt <= d.retries; attempt++ {
		if sendNotification(n) {
			log.Printf("[Worker] Notification sent to %s for post %s", n.FollowerID, n.PostID)
			return
		}
		backoff := time.Duration(1<<attempt) * time.Second
		log.Printf("[Worker] Failed to send to %s. Retrying in %v...", n.FollowerID, backoff)
		time.Sleep(backoff)
	}
	log.Printf("[Worker] Giving up on sending to %s after %d attempts", n.FollowerID, d.retries)
}

func sendNotification(n Notification) bool {
	return rand.Intn(100) >= 10 // 90% success rate
}
