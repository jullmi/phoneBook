package internal

var Index map[string]int

func CreateIndex() error {
	Index = make(map[string]int)

	for i, k := range Data {
		key := k.Tel
		Index[key] = i
	}

	return nil
}
