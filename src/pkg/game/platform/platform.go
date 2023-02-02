package platform

// TODO: Validationメソッドの追加

// PlatformID
type ID uint64

// プラットフォーム名
type Name string

// 1 ≦ name.length ≦ 256
func (n Name) Valid() bool {
	return len(n) > 0 && len(n) < 256
}

// プラットフォーム説明
type Description string

// 1 ≦ Description.length ≦ 2049
func (d Description) Valid() bool {
	return len(d) > 0 && len(d) < 2049
}

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
