package data

var Users = []string{"user1", "user2", "user3", "user4"}

var Followers = map[string][]string{
	"user1": {"user2", "user3"},
	"user2": {"user1", "user4"},
	"user3": {"user2"},
	"user4": {"user1", "user3"},
}