package auth

import (
	"regexp"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"github.com/dandoh/sdr/util"
	"github.com/dandoh/sdr/model"
	"context"
	"os"
)

type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}


func RequireAuth(next http.Handler) http.Handler {
	jwtSecret := []byte(os.Getenv("SECRET_KEY"));
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		util.PrintBody(r)
		// extract jwt
		authorizationHeader := r.Header.Get("Authorization")
		authRegex, _ := regexp.Compile("(?:Bearer *)([^ ]+)(?: *)")
		authRegexMatches := authRegex.FindStringSubmatch(authorizationHeader)
		if len(authRegexMatches) != 2 {
			// didn't match valid Authorization header pattern
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		jwtToken := authRegexMatches[1]

		// parse tokentoken
		token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return jwtSecret, nil
		})
		if err != nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// extract claims
		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		// load userID
		authContext := model.AuthorContext{
			AuthorID: claims.UserID,
		}
		ctx := context.WithValue(r.Context(), "authorContext", authContext)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
