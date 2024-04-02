package internal

import (
	"fmt"
)

func Search(searchValue string) *Subscriber {
	for _, sub := range Data {
		if searchValue == sub.Tel {
			return &sub
		}
	}
	return nil
}

func List() {
	for _, value := range Data {
		fmt.Println(value)
	}
}

func Insert(pS *Subscriber) error {
	_, ok := Index[(*pS).Tel]
	if ok {
		fmt.Errorf("%s already exists", pS.Tel)
	}
	Data = append(Data, *pS)

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
	Data = append(Data[:i], Data[i+1:]...)
	delete(Index, key)

	err := SaveCSVfile(PATH)
	if err != nil {
		return err
	}
	fmt.Println("Success: subscriber was deleted")
	return nil
}
