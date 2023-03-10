package errors

// TODO: エラーの設計微妙。要検討
// エラー詳細情報
// DBでエラー出た時とか、decodeでエラー出た時とかのmsgなど、APIユーザーに見せたくない情報をdetailに配置
type Information struct {
	Code    Coder       // エラーコード
	Detail  string      // errorメッセージ
	Problem interface{} // 問題発生詳細データ
}

func NewInformation(code Coder, detail string, problem interface{}) *Information {
	return &Information{code, detail, problem}
}

// エラー情報取得Interface
type Informator interface {
	error
	Information() *Information
}

// Problem Sample
type InvalidParams struct {
	Name  string
	Param interface{}
}

func NewInvalidParams(name string, param interface{}) InvalidParams {
	return InvalidParams{name, param}
}
