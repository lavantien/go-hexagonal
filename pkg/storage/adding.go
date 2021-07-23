package storage

import (
	"log"

	"github.com/google/uuid"
	"github.com/lavantien/go-hexagonal/pkg/adding"
)

func (s *Storage) AddCandy(c adding.Candy) (string, error) {
	id := uuid.New().String()
	if err := s.db.Query(`INSERT INTO candies (candy_id, category, name, price) VALUES (?, ?, ?, ?)`,
		id, c.Category, c.Name, c.Price).Exec(); err != nil {
		log.Println("Error while trying to save to DB:", err)
		return "", err
	}

	return id, nil
}
