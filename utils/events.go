package utils

import "fmt"

// EventNotFound when event is not found and need to retry
var EventNotFound = fmt.Errorf("event not found")

// EventIterator contains functions that make events listening more easier
type EventIterator interface {
	Next() bool
	Error() error
	Close() error
}

// LookForEvent checks if the iterator is ready with the Event
// if no event is found, returns EventNotFound
// returns iter.Error when iterator errored out
func LookForEvent(iter EventIterator) (err error) {
	defer iter.Close()
	if iter.Next() {
		return nil
	}

	if iter.Error() != nil {
		return iter.Error()
	}

	return EventNotFound
}
