package internal

import (
	"fmt"
	"sort"
)

func Search(searchValue string) *Subscriber {
	for _, sub := range DATA {
		if searchValue == sub.Tel {
			return &sub
		}
	}
	return nil
}

func List() {
	sort.Sort(DATA)
	for _, value := range DATA {
		fmt.Println(value)
	}
}

func Insert(pS *Subscriber) error {
	_, ok := Index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	DATA = append(DATA, *pS)

	_ = CreateIndex()

	err := SaveCSVfile(PATH)
	if err != nil {
		return err
	}
	fmt.Println("Success: subscriber was inserted")
	return nil
}

func DeleteSubscriber(key string) error {
	i, ok := Index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}
	DATA = append(DATA[:i], DATA[i+1:]...)
	delete(Index, key)

	err := SaveCSVfile(PATH)
	if err != nil {
		return err
	}
	fmt.Println("Success: subscriber was deleted")
	return nil
}
