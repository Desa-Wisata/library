package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/desa-wisata/library/logger"

	"github.com/desa-wisata/library/response"
	"github.com/desa-wisata/library/utils"
)

//JWTValidation ...
func JWTValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		header := r.Header.Get("Authorization")
		rAuth := strings.Split(header, " ")

		if len(rAuth) != 2 {
			response.Errors(ctx, http.StatusUnauthorized, "Invalid Bearer").Default(w)
			return
		}

		claim, err := utils.ParseJWT(rAuth[1])
		if err == utils.ErrInvalidToken {
			response.Errors(ctx, http.StatusUnauthorized, "Unathorized").Default(w)
			return
		}
		if err != nil {
			response.Errors(ctx, http.StatusUnauthorized, "Invalid Tokens").Default(w)
			return
		}

		jtwCtx := claim.SetValueToContext(ctx)
		value := utils.GetValueFromContext(jtwCtx)

		// Add userCode to log context
		v := jtwCtx.Value(logger.LogKey).(*logger.Data)
		v.ID = fmt.Sprintf("%d %s", value.UserID, value.Role)
		newCtx := context.WithValue(jtwCtx, logger.LogKey, v)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

// IsStudent ...
func IsStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := utils.GetValueFromContext(ctx)
		if v.Role != "students" {
			response.Errors(ctx, http.StatusForbidden, "Forbidden").Default(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// IsTeacher ...
func IsTeacher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := utils.GetValueFromContext(ctx)
		if v.Role != "teachers" {
			response.Errors(ctx, http.StatusForbidden, "Forbidden").Default(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// IsWaliKelas ...
func IsWaliKelas(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := utils.GetValueFromContext(ctx)
		if v.Role != "wali_kelas" {
			response.Errors(ctx, http.StatusForbidden, "Forbidden").Default(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// IsParent ...
func IsParent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := utils.GetValueFromContext(ctx)
		if v.Role != "parents" {
			response.Errors(ctx, http.StatusForbidden, "Forbidden").Default(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
