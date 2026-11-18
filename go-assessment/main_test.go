package main

import "testing"

func TestAddItem(t *testing.T) {
	var items []Item

	err := addItem(&items, Item{1, "Laptop", 5, 50000})

	if err != nil {
		t.Fatal("Item was not added")
	}

	if len(items) != 1 {
		t.Fatal("Item count is incorrect")
	}
}

func TestUpdateItem(t *testing.T) {
	items := []Item{{1, "Laptop", 5, 50000}}

	err := updateItem(items, 1, 10, 55000)

	if err != nil {
		t.Fatal("Item was not updated")
	}

	if items[0].Quantity != 10 {
		t.Fatal("Quantity was not updated")
	}
}

func TestRemoveItem(t *testing.T) {
	items := []Item{
		{1, "Laptop", 5, 50000},
		{2, "Mouse", 10, 500},
	}

	err := removeItem(&items, 1)

	if err != nil {
		t.Fatal("Item was not removed")
	}

	if len(items) != 1 {
		t.Fatal("Item count is incorrect")
	}
}

func TestAddInvalidItem(t *testing.T) {
	var items []Item

	err := addItem(&items, Item{0, "Laptop", 5, 50000})

	if err == nil {
		t.Fatal("Invalid item should not be added")
	}
}
