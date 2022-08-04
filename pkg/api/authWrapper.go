package api

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func authJWTWrapper(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("X-Authentication-Token")
		if len(jwtToken) < 1 || jwtToken[:2] != "ey" {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString("jwtSecret")), nil
			})

			if err == nil {
				_ = token
				for key, val := range claims {
					fmt.Printf("Key: %v, value: %v\n", key, val)
				}
				next(w, r)
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))

			}

		}
	}
}

func authorizeRequest(r *http.Request) *http.Request {
	ctx := r.Context()
	return r.WithContext(ctx)
}
