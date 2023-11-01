package repository

type Repo struct {
	BookRepo BookRepo
}

func New() (*Repo, error) {
	db, err := connectToPostgres()
	if err != nil {
		return nil, err
	}

	return &Repo{
		BookRepo: *newBookRepository(db),
	}, nil
}
