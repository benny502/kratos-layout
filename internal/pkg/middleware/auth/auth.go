package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

const EXPTTL = 24 * time.Hour

var currentUserKey struct{}

type CurrentUser struct {
	UserId int64
}

type MyCustomClaims struct {
	Id int64 `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int64, secret string, issuer string) string {
	claims := MyCustomClaims{
		id,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(EXPTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		//panic(err)
		tokenString = ""
	}
	return tokenString
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tokenString := GetCustomCookie(ctx, "stoken")
			if tokenString == "" {
				return nil, errors.New("jwt token missing")
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secret), nil
			})

			if err != nil {
				return nil, err
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if u, ok := claims["id"]; ok {
					ctx = WithContext(ctx, &CurrentUser{UserId: int64(u.(float64))})
				}
				// put CurrentUser into ctx
				//SetCustomCookie(ctx, "stoken="+GenerateToken(int64(claims["id"].(float64)), ""))
			} else {
				return nil, errors.New("Token Invalid")
			}
			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) *CurrentUser {
	return ctx.Value(currentUserKey).(*CurrentUser)
}

func WithContext(ctx context.Context, user *CurrentUser) context.Context {
	return context.WithValue(ctx, currentUserKey, user)
}

func SetCustomCookie(ctx context.Context, s string) {
	if tr, ok := transport.FromServerContext(ctx); ok {
		cookie := tr.RequestHeader().Get("Cookie")
		cookies := strings.Split(cookie, "; ")
		cookies = append(cookies, s)
		cookieMap := make(map[string]string)
		for i := range cookies {
			tmp := strings.Split(cookies[i], "=")
			if len(tmp) == 2 {
				cookieMap[tmp[0]] = tmp[1]
			}
		}
		cookie = ""
		for i, v := range cookieMap {
			cookie += "; " + i + "=" + v
		}
		cookie = cookie[1:]
		tr.ReplyHeader().Set("Set-Cookie", cookie)
	}
}

func GetCustomCookie(ctx context.Context, key string) string {
	if tr, ok := transport.FromServerContext(ctx); ok {
		cookie := tr.RequestHeader().Get("Cookie")
		cookies := strings.Split(cookie, "; ")
		cookieMap := make(map[string]string)
		for i := range cookies {
			tmp := strings.Split(cookies[i], "=")
			if len(tmp) == 2 {
				cookieMap[tmp[0]] = tmp[1]
			}
		}
		if c, o := cookieMap[key]; o {
			return c
		}
	}
	return ""
}
