package internal

import (
	"fmt"
)

func Search(subs []Subscriber, searchValue string) *Subscriber {
	for _, sub := range subs {
		if searchValue == sub.Tel {
			return &sub
		}
	}
	return nil
}

func List(subs []Subscriber) {
	for _, value := range subs {
		fmt.Println(value)
	}
}

