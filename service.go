package main

import (
	"encoding/json"
	"os"
)

type CategoryService struct {
	File string
}

func (cs *CategoryService) load() ([]Category) {
	data, _ := os.ReadFile(cs.File)

	var categories []Category
	json.Unmarshal(data, &categories)
	return categories
}

func (cs *CategoryService) save(categories []Category) {
	data, _ := json.MarshalIndent(categories, "", " ")
	_ = os.WriteFile(cs.File, data, 0644)
}

func (cs *CategoryService) Add(cat Category) {
	categories := cs.load()

	maxID := 0
	for _, c := range categories {
		if c.ID > maxID {
			maxID = c.ID
		}
	}
	cat.ID = maxID + 1

	categories = append(categories, cat)
	cs.save(categories)
}

func (cs *CategoryService) Update(id int, updated Category) {
	categories := cs.load()

	for i, c := range categories {
		if c.ID == id {
			updated.ID = id
			categories[i] = updated
			break;
		}
	}

	cs.save(categories)
}

func (cs *CategoryService) List() []Category {
	return cs.load()
}
