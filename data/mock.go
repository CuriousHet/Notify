package data

import "sync"

var Users = []string{"user1", "user2", "user3", "user4"}

var Followers = map[string][]string{
	"user1": {"user2", "user3"},
	"user2": {"user1", "user4"},
	"user3": {"user2"},
	"user4": {"user1", "user3"},
}

var NotificationsStore = make(map[string][]string)
var mu sync.Mutex

func Save(userID, message string) {
	mu.Lock()
	defer mu.Unlock()
	NotificationsStore[userID] = append(NotificationsStore[userID], message)
}

func Get(userID string) []string {
	mu.Lock()
	defer mu.Unlock()
	return NotificationsStore[userID]
}
