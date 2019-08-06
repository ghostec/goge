package event

import "github.com/ghostec/goge/store"

type Key string

type Event struct {
	key Key
	*store.Store
}

func New(key Key) *Event {
	return &Event{key: key, Store: store.New()}
}

func (e Event) Key() Key {
	return e.key
}
