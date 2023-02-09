package errors

// エラー発生レイヤー
type Layer uint8

const (
	Layer_Undefined Layer = iota
	Layer_Request
	Layer_Domain
	Layer_Repository
	Layer_Model
)

func (l Layer) String() string {
	switch l {
	case Layer_Undefined:
		return "UNDEFINED"
	case Layer_Request:
		return "REQUEST"
	case Layer_Domain:
		return "DOMAIN"
	case Layer_Repository:
		return "REPOSITORY"
	case Layer_Model:
		return "MODEL"
	default:
		return "UNKNOWN"
	}
}
