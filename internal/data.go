package internal

var Data []Subscriber = []Subscriber{}

func UseData() []Subscriber {
	sub1 := NewSubscriber("Julia", "Panova", "555-555-777")
	sub2 := NewSubscriber("Mikal", "Panov", "555-555-887")
	sub3 := NewSubscriber("Maria", "Ivanova", "555-555-897")

	subscribers := append(Data, *sub1, *sub2, *sub3)

	return subscribers
}
