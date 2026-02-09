package token

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrExpired = errors.New("token is expired")
	ErrInvalid = errors.New("token is invalid")
)

func Generate(claims jwt.Claims) (string, error) {
	now := time.Now()
	exp := jwt.NewNumericDate(now.Add(expiration))
	iat := jwt.NewNumericDate(now)

	switch c := claims.(type) {
	case jwt.MapClaims:
		c["iss"] = issuer
		c["exp"] = exp
		c["iat"] = iat
	case *jwt.RegisteredClaims:
		c.Issuer = issuer
		c.ExpiresAt = exp
		c.IssuedAt = iat
	default:
		v := reflect.ValueOf(claims)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if v.Kind() != reflect.Struct {
			return "", fmt.Errorf("unsupported claims type: must be jwt.MapClaims, *jwt.RegisteredClaims, or a struct embedding jwt.RegisteredClaims")
		}

		regClaimsField := v.FieldByName("RegisteredClaims")
		if !regClaimsField.IsValid() || !regClaimsField.CanSet() {
			return "", fmt.Errorf("invalid claims struct: must embed jwt.RegisteredClaims and be a pointer")
		}

		newRegClaims := jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: exp,
			IssuedAt:  iat,
		}
		regClaimsField.Set(reflect.ValueOf(newRegClaims))

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func Validate(tokenString string, claims jwt.Claims) error {
	if claims == nil {
		return fmt.Errorf("claims parameter cannot be nil: pass a non-nil pointer to your claims struct (e.g., &MyClaims{} or jwt.MapClaims{})")
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return ErrExpired
		}
		return fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return ErrInvalid
	}

	return nil
}
