package internal

import (
	"strconv"
	"time"
)

type Subscriber struct {
	Name    string
	Surname string
	Tel     string
	LastAccess string
}

func NewSubscriber(name, surname, tel string) *Subscriber {
	if tel == "" || surname == "" {
		return nil
	}

	lastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Subscriber{Name: name, Surname: surname, Tel: tel, LastAccess: lastAccess}
}
