package repository

import (
	"os"
	"strconv"
)

type Repo struct {
	BookRepo    *BookRepository
	UserRepo    *UserRepository
	SessionRepo *SessionRepository
}

func New() (*Repo, error) {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic("Port must be an integer")
	}

	db, err := connectToPostgres(host, user, pass, name, portInt)
	if err != nil {
		return nil, err
	}

	return &Repo{
		BookRepo:    newBookRepository(db),
		UserRepo:    newUserRepository(db),
		SessionRepo: newSessionRepository(db),
	}, nil
}
