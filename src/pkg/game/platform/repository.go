package platform

import "mysrtafes-backend/pkg/errors"

type Repository interface {
	PlatformCreate(*Platform) (*Platform, error)
	PlatformRead(ID) (*Platform, error)
	PlatformFind(*FindOption) ([]*Platform, error)
	PlatformUpdate(*Platform) (*Platform, error)
	PlatformDelete(ID) error
}

type Server interface {
	Create(*Platform) (*Platform, error)
	Read(ID) (*Platform, error)
	Find(*FindOption) ([]*Platform, error)
	Update(*Platform) (*Platform, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(p *Platform) (*Platform, error) {
	// 名前のValidate
	if !p.Name.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("name", p.Name),
				},
			),
			"Name Valid error",
		)
	}
	// DescriptionのValidate
	if !p.Description.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("description", p.Name),
				},
			),
			"Description Valid error",
		)
	}
	return s.repository.PlatformCreate(p)
}

func (s *server) Read(id ID) (*Platform, error) {
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
	return s.repository.PlatformRead(id)
}

func (s *server) Find(findOption *FindOption) ([]*Platform, error) {
	return s.repository.PlatformFind(findOption)
}

func (s *server) Update(p *Platform) (*Platform, error) {
	// IDのValidate
	if !p.ID.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("id", p.ID),
				},
			),
			"ID Valid error",
		)
	}
	// 名前のValidate
	if !p.Name.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("name", p.Name),
				},
			),
			"Name Valid error",
		)
	}
	// DescriptionのValidate
	if !p.Description.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("description", p.Name),
				},
			),
			"Description Valid error",
		)
	}
	return s.repository.PlatformUpdate(p)
}

func (s *server) Delete(id ID) error {
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
	return s.repository.PlatformDelete(id)
}
