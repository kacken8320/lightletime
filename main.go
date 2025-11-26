package main

import "fmt"

func main() {
	service := &CategoryService{File: "categories.json"}

	service.Add(Category{
		Name:			"root",
		Multiplier:		1.0,
		TimeSpentInMinutes:	0,
		MinimalTimeInSeconds:	0,
		ParentID:		0,
		IsActivity:		false,
	})
	service.Add(Category{
		Name:			"C",
		Multiplier:		3.0,
		TimeSpentInMinutes:	10,
		MinimalTimeInSeconds:	600,
		ParentID:		1,
		IsActivity:		true,
	})
	service.Add(Category{
		Name:			"git",
		Multiplier:		1.5,
		TimeSpentInMinutes:	0,
		MinimalTimeInSeconds:	1800,
		ParentID:		1,
		IsActivity:		true,
	})

	cats := service.List()
	for _, c := range cats {
		fmt.Printf("%d: %s (x%.1f)\n", c.ID, c.Name, c.Multiplier)
	}
}
