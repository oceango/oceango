package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/oceango/skeleton/model"
	"github.com/oceango/web/db"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)


func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !enable(r) {
			next.ServeHTTP(w, r)
		}else {
			dbm := db.GetDb()
			authorizationstring := r.Header.Get("Authorization")
			if len(authorizationstring) == 0 || !strings.HasPrefix(authorizationstring, "bearer ") {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			tokenString := r.Header.Get("Authorization")[7:]
			userId := getUserIdFromToken(tokenString)
			if userId <= 0 {
				log.Printf( "%s is not exist", userId)
			}

			// query user
			var user model.User
			dbm.First(&user, "id = ?", userId)

			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}

	return http.HandlerFunc(fn)
}

func enable(r *http.Request) bool {
	path := r.URL.Path

	var enablePaths []string
	enablePaths = append(append(enablePaths, "auth/info"), "welcome")


	for _, enablePath := range enablePaths  {
		re := regexp.MustCompile(enablePath)
		if re.MatchString(path) {
			return true
		}
	}
	return false
}

func getUserIdFromToken(tokenString string) int {
	type MyCustomClaims struct {
		jwt.StandardClaims
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v", claims.StandardClaims.ExpiresAt)
		userId, err := strconv.Atoi(claims.StandardClaims.Id)
		if err != nil {
			log.Print(err)
		}
		return userId
	} else {
		fmt.Println(err)
		return 0
	}

}
