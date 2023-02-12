package errors

import "fmt"

// Badリクエスト
type InvalidRequestError interface {
	error
	ErrorInvalidRequestError() // ダミーメソッド
}

type invalidRequest struct {
	layer Layer
	info  *Information
	msg   string
}

func NewInvalidRequest(layer Layer, info *Information, msg string) InvalidRequestError {
	return &invalidRequest{layer, info, msg}
}

func (r *invalidRequest) Information() *Information {
	return r.info
}

func (r *invalidRequest) Error() string {
	return fmt.Sprintf("%s: INVALID_REQUEST: %s", r.layer, r.msg)
}

func (r *invalidRequest) ErrorInvalidRequestError() {}

// バリデートエラー
type InvalidValidateError interface {
	error
	ErrorInvalidValidateError() // ダミーメソッド
}

type invalidValidate struct {
	layer Layer
	info  *Information
	msg   string
}

func NewInvalidValidate(layer Layer, info *Information, msg string) InvalidValidateError {
	return &invalidValidate{layer, info, msg}
}

func (r *invalidValidate) Information() *Information {
	return r.info
}

func (r *invalidValidate) Error() string {
	return fmt.Sprintf("%s: INVALID_VALIDATE: %s", r.layer, r.msg)
}

func (r *invalidValidate) ErrorInvalidValidateError() {}

// Not Found
type NotFoundError interface {
	error
	ErrorNotFoundError() // ダミーメソッド
}

type notFound struct {
	layer Layer
	info  *Information
	msg   string
}

func NewNotFound(layer Layer, info *Information, msg string) NotFoundError {
	return &notFound{layer, info, msg}
}

func (r *notFound) Information() *Information {
	return r.info
}

func (r *notFound) Error() string {
	return fmt.Sprintf("%s: NOT_FOUND: %s", r.layer, r.msg)
}

func (r *notFound) ErrorNotFoundError() {}

// 未サポートメディアタイプ指定
type UnsupportedMediaTypeError interface {
	error
	ErrorUnsupportedMediaTypeError() // ダミーメソッド
}

type unsupportedMediaType struct {
	layer Layer
	info  *Information
	msg   string
}

func NewUnsupportedMediaType(layer Layer, info *Information, msg string) UnsupportedMediaTypeError {
	return &unsupportedMediaType{layer, info, msg}
}

func (r *unsupportedMediaType) Information() *Information {
	return r.info
}

func (r *unsupportedMediaType) Error() string {
	return fmt.Sprintf("%s: UNSUPPORTED_MEDIA_TYPE: %s", r.layer, r.msg)
}

func (r *unsupportedMediaType) ErrorUnsupportedMediaTypeError() {}

// 未認証
type UnauthorizedError interface {
	error
	ErrorUnauthorizedError() // ダミーメソッド
}

type unauthorized struct {
	layer Layer
	info  *Information
	msg   string
}

func NewUnauthorized(layer Layer, info *Information, msg string) UnauthorizedError {
	return &unauthorized{layer, info, msg}
}

func (r *unauthorized) Information() *Information {
	return r.info
}

func (r *unauthorized) Error() string {
	return fmt.Sprintf("%s: UNSUPPORTED_MEDIA_TYPE: %s", r.layer, r.msg)
}

func (r *unauthorized) ErrorUnauthorizedError() {}

// アクセス禁止
type ForbiddenError interface {
	error
	ErrorForbiddenError() // ダミーメソッド
}

type forbidden struct {
	layer Layer
	info  *Information
	msg   string
}

func NewForbidden(layer Layer, info *Information, msg string) ForbiddenError {
	return &forbidden{layer, info, msg}
}

func (r *forbidden) Information() *Information {
	return r.info
}

func (r *forbidden) Error() string {
	return fmt.Sprintf("%s: UNSUPPORTED_MEDIA_TYPE: %s", r.layer, r.msg)
}

func (r *forbidden) ErrorForbiddenError() {}

// アクセス禁止
type InternalServerErrorError interface {
	error
	ErrorInternalServerErrorError() // ダミーメソッド
}

type internalServerError struct {
	layer Layer
	info  *Information
	msg   string
}

func NewInternalServerError(layer Layer, info *Information, msg string) InternalServerErrorError {
	return &internalServerError{layer, info, msg}
}

func (r *internalServerError) Information() *Information {
	return r.info
}

func (r *internalServerError) Error() string {
	return fmt.Sprintf("%s: INTERNAL_SERVER_ERROR: %s", r.layer, r.msg)
}

func (r *internalServerError) ErrorInternalServerErrorError() {}
