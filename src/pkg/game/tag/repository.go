package tag

import "mysrtafes-backend/pkg/errors"

type Repository interface {
	TagCreate(*Tag) (*Tag, error)
	TagRead(ID) (*Tag, error)
	TagUpdate(*Tag) (*Tag, error)
	TagDelete(ID) error
}

type Server interface {
	Create(*Tag) (*Tag, error)
	Read(ID) (*Tag, error)
	Update(*Tag) (*Tag, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

// GameTagの作成
func (s *server) Create(t *Tag) (*Tag, error) {
	// 名前のValidate
	if !t.Name.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("name", t.Name),
				},
			),
			"Name Valid error",
		)
	}
	// DescriptionのValidate
	if !t.Description.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("description", t.Name),
				},
			),
			"Description Valid error",
		)
	}
	return s.repository.TagCreate(t)
}

// GameTagの読み込み
func (s *server) Read(id ID) (*Tag, error) {
	// IDのValidate
	if !id.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
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
	return s.repository.TagRead(id)
}

// GameTagの更新
func (s *server) Update(t *Tag) (*Tag, error) {
	// IDのValidate
	if !t.ID.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("id", t.ID),
				},
			),
			"ID Valid error",
		)
	}
	// 名前のValidate
	if !t.Name.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("name", t.Name),
				},
			),
			"Name Valid error",
		)
	}
	// DescriptionのValidate
	if !t.Description.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Request,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("description", t.Name),
				},
			),
			"Description Valid error",
		)
	}
	return s.repository.TagUpdate(t)
}

// GameTagの削除
func (s *server) Delete(id ID) error {
	if !id.Valid() {
		return errors.NewInvalidRequest(
			errors.Layer_Request,
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
	return s.repository.TagDelete(id)
}
