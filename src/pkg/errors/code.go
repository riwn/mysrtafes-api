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
	ID_JsonDecodeError
	ID_DBCreateError
	ID_DBReadError
	ID_DBUpdateError
	ID_DBDeleteError
	ID_DBDataFormatError
)

func (id ID) Detail() (Code, Message) {
	switch id {
	case ID_InvalidParams:
		return "E00001", "invalid params error"
	case ID_JsonDecodeError:
		return "E00002", "json decode error"
	case ID_DBCreateError:
		return "E10001", "database create error"
	case ID_DBReadError:
		return "E10002", "database read error"
	case ID_DBUpdateError:
		return "E10003", "database update error"
	case ID_DBDeleteError:
		return "E10004", "database delete error"
	case ID_DBDataFormatError:
		return "E10005", "database data format error"
	default:
		return "E99999", "unknown error"
	}
}
