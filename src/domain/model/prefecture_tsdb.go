package model

type PrefectureTSDB struct {
	Name string
}

func NewPrefectureTSDB(name string) *PrefectureTSDB {
	return &PrefectureTSDB{
		Name: name,
	}
}
