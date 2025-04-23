package gql

import (
	"github.com/CuriousHet/Notify/data"
)

// Resolver contains methods to handle GraphQL queries
type Resolver struct{}

// Notifications resolves the notifications for a specific user
func (r *Resolver) Notifications(args struct{ UserID string }) []string {
	return data.Get(args.UserID)
}
