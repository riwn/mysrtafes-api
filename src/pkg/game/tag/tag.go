package tag

// TODO: Validationメソッドの追加

// TagID
type ID uint64

// タグ名
type Name string

// 1 ≦ name.length ≦ 256
func (n Name) Valid() bool {
	return len(n) > 0 && len(n) < 256
}

// タグ説明
type Description string

// 1 ≦ Description.length ≦ 2049
func (d Description) Valid() bool {
	return len(d) > 0 && len(d) < 2049
}

// タグ
type Tag struct {
	ID          ID
	Name        Name
	Description Description
}
