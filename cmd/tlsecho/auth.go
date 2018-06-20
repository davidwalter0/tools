package main

/*
import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
)

func RequireTokenAuthentication(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
			return VERIFICATION.PublicKey, nil
		})
		if err != nil || !token.Valid {
			log.Debug("Authentication failed " + err.Error())
			w.WriteHeader(http.StatusForbidden)
			return
		} else {
			r.Header.Set("username", token.Claims["username"].(string))
			r.Header.Set("userid", strconv.FormatFloat((token.Claims["userid"]).(float64), 'f', 0, 64))
		}
		inner.ServeHTTP(w, r)
	})
}
*/
