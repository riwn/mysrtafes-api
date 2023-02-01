package platform

// TODO: Validationメソッドの追加

// PlatformID
type ID uint64

// プラットフォーム名
type Name string

// プラットフォーム説明
type Description string

// プラットフォーム
type Platform struct {
	ID          ID
	Name        Name
	Description Description
}

func New(id ID, name Name, description Description) *Platform {
	return &Platform{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
