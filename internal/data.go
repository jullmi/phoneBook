package internal

import "strconv"


var Data []Subscriber = []Subscriber{}

func PopulateData(len int) []Subscriber {
	for i := 0; i < len; i++ {
		name := generateString(4)
		surname := generateString(5)
		tel := strconv.Itoa(random(100, 999))
		Data = append(Data, Subscriber{name, surname, tel})
	}

	return Data
}

