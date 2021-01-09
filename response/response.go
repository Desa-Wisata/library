package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/desa-wisata/library/logger"
)

const (
	//ErrNotFound ...
	ErrNotFound = `Data tidak ditemukan`
	//ErrCommonServer ...
	ErrCommonServer = `Terjadi kesalahan pada server. Silahkan coba beberapa saat lagi`
	//ErrLogin ...
	ErrLogin = `Invalid Login`
	//SuccessRes ...
	SuccessRes = `SUCCESS`
)

//Response is default response
type rs struct {
	Code         int             `json:"code"`
	Status       bool            `json:"status"`
	ErrorMessage string          `json:"error_message,omitempty"`
	Data         interface{}     `json:"data,omitempty"`
	Ctx          context.Context `json:"-"`
}

//Message ...
type Message interface {
	Default(w http.ResponseWriter)
	DefaultHTML(w http.ResponseWriter)
	IsError() bool
}

//Default is ...
func (r *rs) Default(w http.ResponseWriter) {
	logger.EndRecord(r.Ctx, r.Code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}

//DefaultHTML ...
func (r *rs) DefaultHTML(w http.ResponseWriter) {
	logger.EndRecord(r.Ctx, r.Code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	encoder.Encode(r)
}

func (r *rs) IsError() bool {
	return !r.Status
}

//NotFound ...
func NotFound(ctx context.Context) Message {
	return &rs{
		ErrorMessage: ErrNotFound,
		Code:         http.StatusNotFound,
		Ctx:          ctx,
	}
}

//InternalError ...
func InternalError(ctx context.Context) Message {
	return &rs{
		ErrorMessage: ErrCommonServer,
		Code:         http.StatusInternalServerError,
		Ctx:          ctx,
	}
}

//Success ...
func Success(ctx context.Context, data interface{}) Message {
	return &rs{
		Code:   http.StatusOK,
		Status: true,
		Data:   data,
		Ctx:    ctx,
	}
}

//Errors ...
func Errors(ctx context.Context, code int, msg string) Message {
	return &rs{
		Code:         code,
		ErrorMessage: msg,
		Ctx:          ctx,
	}
}

//SuccessWithCode ...
func SuccessWithCode(ctx context.Context, code int, data interface{}) Message {
	return &rs{
		Code:   code,
		Status: true,
		Data:   data,
		Ctx:    ctx,
	}
}
