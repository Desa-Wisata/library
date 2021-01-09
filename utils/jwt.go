package utils

import (
	"context"
	"errors"
	"os"
	"regexp"

	"github.com/dgrijalva/jwt-go"
)

type key int

const (
	jwtclaim = key(47)
	jwtform  = `^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`
)

var (
	jKeys = []byte(os.Getenv("JWT_SALT"))
	//ErrInvalidToken ...
	ErrInvalidToken = errors.New(`Invalid tokens`)
)

//Receiver is ....
type Receiver struct {
	UserID       int    `json:"user_id"`
	Code         string `json:"code"`
	Role         string `json:"role"`
	IsWaliKelas  bool   `json:"is_wali_kelas"`
	IsGuruMatpel bool   `json:"is_guru_matpel"`
}

//JWTClaim is part of inside jwt
type JWTClaim interface {
	SetJWT() (string, error)
	SetValueToContext(ctx context.Context) context.Context
}

//claim is struct for claim
type claim struct {
	Receiver
	jwt.StandardClaims
}

//SetJWT ...
func (c *claim) SetJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(jKeys)
}

//SetValueToContext ...
func (c *claim) SetValueToContext(ctx context.Context) context.Context {
	ct := context.WithValue(ctx, jwtclaim, c.Receiver)
	return ct
}

//GetValueFromContext ...
func GetValueFromContext(ctx context.Context) *Receiver {

	rs, ok := ctx.Value(jwtclaim).(Receiver)
	if !ok {
		return &Receiver{}
	}

	return &rs
}

//NewClaim ...
func NewClaim(userid int, code, role string, IsWaliKelas, IsGuruMatpel bool) JWTClaim {
	return &claim{
		Receiver: Receiver{userid, code, role, IsWaliKelas, IsGuruMatpel},
	}
}

//ParseJWT ...
func ParseJWT(tn string) (JWTClaim, error) {
	c := new(claim)

	if match, _ := regexp.MatchString(jwtform, tn); !match {
		return nil, errors.New("not jwt format")
	}

	tk, err := jwt.ParseWithClaims(tn, c,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jKeys, nil
		})

	if err == jwt.ErrSignatureInvalid || !tk.Valid {
		return nil, ErrInvalidToken
	}

	if err != nil {
		return nil, err
	}
	return c, nil
}
