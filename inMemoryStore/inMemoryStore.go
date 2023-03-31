package inMemoryStore

import "sync"

var InMemoryDb sync.Map

func init() {
	InMemoryDb = sync.Map{}
}
