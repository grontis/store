package storage

import (
	"errors"
	"fmt"
	"grontis/store/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Memory struct {
	items []models.Item
	users []models.User
}

func NewMemory() *Memory {
	m := new(Memory)
	m.items = buildMockItems()
	m.users = buildMockUsers()
	return m
}

func (m Memory) GetItems() ([]models.Item, error) {
	return m.items, nil
}

func (m Memory) GetItem(id uuid.UUID) (models.Item, error) {
	item := findFirst[models.Item](m.items, id)

	if item != nil {
		return *item, nil
	} else {
		return models.Item{}, errors.New("Item not found")
	}
}

func (m *Memory) CreateItem(item models.Item) (models.Item, error) {
	if item.Id == uuid.Nil {
		item.Id = uuid.New()
	}
	m.items = append(m.items, item)
	return item, nil
}

func (m *Memory) UpdateItem(item models.Item) (models.Item, error) {
	for i, v := range m.items {
		if v.Id == item.Id {
			m.items[i] = item
			return item, nil
		}
	}
	return models.Item{}, errors.New("Item not found")
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

func (m Memory) GetUsers() ([]models.User, error) {
	return m.users, nil
}

func (m Memory) GetUser(id uuid.UUID) (models.User, error) {
	user := findFirst[models.User](m.users, id)

	if user != nil {
		return *user, nil
	} else {
		return models.User{}, errors.New("User not found")
	}
}

func (m *Memory) CreateUser(user models.User) (models.User, error) {
	if user.Id == uuid.Nil {
		user.Id = uuid.New()
	}
	m.users = append(m.users, user)
	return user, nil
}

func (m *Memory) UpdateUser(user models.User) (models.User, error) {
	for i, v := range m.users {
		if v.Id == user.Id {
			m.users[i] = user
			return user, nil
		}
	}
	return models.User{}, errors.New("User not found")
}

func (m *Memory) DeleteUser(id uuid.UUID) error {
	for i, v := range m.users {
		if v.Id == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return nil
		}
	}
	return errors.New("User not found")
}

func buildMockUsers() []models.User {
	users := []models.User{}
	for i := 0; i < 3; i++ {
		id := uuid.New()
		username := fmt.Sprintf("user%d", i)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("pass%d", i)), bcrypt.DefaultCost)
		users = append(users, models.User{Id: id, Username: username, Password: string(hashedPassword)})
	}
	return users
}

func buildMockItems() []models.Item {
	items := []models.Item{
		{Id: uuid.New(), Name: "Raspberry Pi", Price: 11.99, Description: "RP4"},
		{Id: uuid.New(), Name: "Keyboard", Price: 5.99, Description: "KB"},
		{Id: uuid.New(), Name: "Mouse", Price: 4.99, Description: "M"},
	}
	return items
}

func findFirst[T models.IdGetter](items []T, id uuid.UUID) *T {
	for _, item := range items {
		if item.GetId() == id {
			return &item
		}
	}
	return nil
}
