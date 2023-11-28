package storage

import (
	"errors"

	"github.com/google/uuid"
)

type Memory struct {
	items []Item
}

func NewMemory() *Memory {
	m := new(Memory)
	m.items = buildMockItems()
	return m
}

func (m Memory) GetItems() ([]Item, error) {
	return m.items, nil //TODO error?
}

func (m Memory) GetItem(id uuid.UUID) (Item, error) {
	item := findFirst(m.items, func(item Item) bool {
		return item.Id == id
	})

	if item != nil {
		return *item, nil
	} else {
		return Item{}, errors.New("Item not found")
	}
}

func (m *Memory) CreateItem(item Item) (Item, error) {
	if item.Id == uuid.Nil {
		item.Id = uuid.New()
	}
	m.items = append(m.items, item)
	return item, nil
}

func (m *Memory) UpdateItem(item Item) (Item, error) {
	for i, v := range m.items {
		if v.Id == item.Id {
			m.items[i] = item
			return item, nil
		}
	}

	return Item{}, errors.New("Item not found")
}

func (m *Memory) DeleteItem(id uuid.UUID) error {
	for i, v := range m.items {
		if v.Id == id {
			m.items = append(m.items[:i], m.items[i+1:]...)
			return nil
		}
	}

	return errors.New("Item not found")
}

func findFirst(items []Item, predicate func(Item) bool) *Item {
	for _, item := range items {
		if predicate(item) {
			return &item
		}
	}
	return nil
}

func buildMockItems() []Item {
	items := []Item{
		{Id: uuid.New(), Name: "Raspberry Pi", Price: 11.99, Description: "RP4"},
		{Id: uuid.New(), Name: "Keyboard", Price: 5.99, Description: "KB"},
		{Id: uuid.New(), Name: "Mouse", Price: 4.99, Description: "M"},
	}

	return items
}
