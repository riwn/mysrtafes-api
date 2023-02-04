package errors

// エラーコード
type Code string

// エラーメッセージ
type Message string

type Coder interface {
	Detail() (Code, Message)
}

// エラーID
type ID uint64

const (
	_ ID = iota
	ID_InvalidParams
)

func (id ID) Detail() (Code, Message) {
	switch id {
	case ID_InvalidParams:
		return "E00001", "invalid params error"
	default:
		return "E99999", "unknown error"
	}
}
