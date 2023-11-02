package repository

import (
	"time"

	"github.com/ca-lee-b/go-rest-boilerplate/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func newSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (s *SessionRepository) Create() *models.Session {

	now := time.Now()
	var session = &models.Session{
		Id:     uuid.NewString(),
		Issued: now,
		Expiry: now.AddDate(0, 1, 0),
	}

	result := s.db.Create(&session)
	if result.Error != nil {
		return nil
	}

	return session
}
