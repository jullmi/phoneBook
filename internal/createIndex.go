package internal

var Index map[string]int

func CreateIndex() error {
	Index = make(map[string]int)

	for i, k := range DATA {
		key := k.Tel
		Index[key] = i
	}

	return nil
}
