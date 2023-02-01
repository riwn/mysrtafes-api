package tag

// TODO: Validationメソッドの追加

// TagID
type ID uint64

// タグ名
type Name string

// タグ説明
type Description string

// タグ
type Tag struct {
	ID          ID
	Name        Name
	Description Description
}
