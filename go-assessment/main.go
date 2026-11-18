package main

import (
	"errors"
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

func addItem(items *[]Item, item Item) error {
	if item.ID <= 0 || item.Name == "" || item.Quantity < 0 || item.Price < 0 {
		return errors.New("invalid item details")
	}

	for _, i := range *items {
		if i.ID == item.ID {
			return errors.New("item already exists")
		}
	}

	*items = append(*items, item)
	return nil
}

func updateItem(items []Item, id int, quantity int, price float64) error {
	for i := 0; i < len(items); i++ {
		if items[i].ID == id {
			if quantity < 0 || price < 0 {
				return errors.New("invalid quantity or price")
			}

			items[i].Quantity = quantity
			items[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

func removeItem(items *[]Item, id int) error {
	for i, item := range *items {
		if item.ID == id {
			*items = append((*items)[:i], (*items)[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

func summary(items []Item) {
	ch := make(chan float64)

	go func() {
		var total float64

		for _, item := range items {
			total += float64(item.Quantity) * item.Price
		}

		ch <- total
	}()

	total := <-ch

	fmt.Println("Total number of items:", len(items))
	fmt.Println("Total inventory value:", total)
}

func main() {
	var items []Item

	err := addItem(&items, Item{1, "Laptop", 5, 50000})
	if err != nil {
		fmt.Println(err)
	}

	err = addItem(&items, Item{2, "Mouse", 10, 500})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Inventory:")
	for _, item := range items {
		fmt.Println(item)
	}

	err = updateItem(items, 1, 8, 55000)
	if err != nil {
		fmt.Println(err)
	}

	err = removeItem(&items, 2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nUpdated Inventory:")
	for _, item := range items {
		fmt.Println(item)
	}

	fmt.Println("\nInventory Summary:")
	summary(items)
}
