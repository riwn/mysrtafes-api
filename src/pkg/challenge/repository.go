package challenge

import "mysrtafes-backend/pkg/errors"

type Repository interface {
	ChallengeCreate(*Challenge) (*Challenge, error)
	ChallengeRead(ID) (*Challenge, error)
	ChallengeUpdate(*Challenge) (*Challenge, error)
	ChallengeDelete(ID) error
}

type Server interface {
	Create(*Challenge) (*Challenge, error)
	Read(ID) (*Challenge, error)
	Update(*Challenge) (*Challenge, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(c *Challenge) (*Challenge, error) {
	if err := c.ValidCreate(); err != nil {
		return nil, err
	}

	return s.repository.ChallengeCreate(c)
}

func (s *server) Read(id ID) (*Challenge, error) {
	// IDのValidate
	if !id.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("id", id),
				},
			),
			"ID Valid error",
		)
	}
	return s.repository.ChallengeRead(id)
}

func (s *server) Update(c *Challenge) (*Challenge, error) {
	if err := c.ValidUpdate(); err != nil {
		return nil, err
	}

	return s.repository.ChallengeUpdate(c)
}

func (s *server) Delete(id ID) error {
	// IDのValidate
	if !id.Valid() {
		return errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("id", id),
				},
			),
			"ID Valid error",
		)
	}
	return s.repository.ChallengeDelete(id)
}
