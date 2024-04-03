package internal

type PhoneBook []Subscriber

var DATA = PhoneBook{}


func (entity PhoneBook) Len() int {
	return len(entity)
}

func (entity PhoneBook) Less(i, j int) bool {
	if entity[i].Surname == entity[j].Surname {
		return entity[i].Name < entity[i].Name 
	}
	return entity[i].Surname < entity[j].Surname
} 

func (entity PhoneBook) Swap(i, j int) {
	entity[i], entity[j] = entity[j], entity[i]
}