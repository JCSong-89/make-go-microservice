package data

var data = []Kitten{
	{
		Id:     "1",
		Name:   "Felix",
		Weight: 12.3,
	},
	{
		Id:     "2",
		Name:   "Fat Freddy's Cat",
		Weight: 20.0,
	},
	{
		Id:     "3",
		Name:   "Garfield",
		Weight: 35.0,
	},
}

type MemoryStore struct {
}

//name 매개 변수와 동일한 목 데이터들이 존재하면 append하여 리턴.
func (m *MemoryStore) Search(name string) []Kitten {
	var kittens []Kitten

	for _, k := range data {
		if k.Name == name {
			kittens = append(kittens, k)
		}
	}

	return kittens
}
