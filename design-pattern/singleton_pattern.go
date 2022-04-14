package main

import (
	"fmt"
)

func main() {
	tot := GetTotalValue(&DummyDB{}, "alex", "kim")
	fmt.Println(tot == 3)
}

type DB interface {
	GetValue(name string) int
}

func GetTotalValue(db DB, names ...string) int {
	tot := 0
	for _, name := range names {
		tot += db.GetValue(name)
	}
	return tot
}

type DummyDB struct {
	dummyData map[string]int
}

func (d *DummyDB) GetValue(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"kim":  1,
			"alex": 2,
		}
	}
	return d.dummyData[name]
}
