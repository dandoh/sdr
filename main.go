package main

import (
	_"github.com/jinzhu/gorm/dialects/postgres"
	"sdr/model"
	"github.com/graphql-go/handler"
	"net/http"
	_"fmt"
	_"github.com/rs/cors"
	"github.com/graphql-go/graphql"
	_"encoding/json"
	 _"github.com/dgrijalva/jwt-go"
	_"github.com/graphql-go/graphql"
	_"github.com/graphql-go/handler"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"regexp"
	"time"
	"context"

)








/*
func main() {
	// initialize database
	model.Init();
	model.InitType();

	h := handler.New(&handler.Config{
		Schema: &model.SchemaQL,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	c := cors.Default().Handler(h);
	http.ListenAndServe(":8080", c)
	defer model.Close();
}
*/


func main(){
	setupServer()
}

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()

	// graphql Handler
	graphqlHandler := http.HandlerFunc(graphqlHandlerFunc)

	// login Handler
	mux.HandleFunc("/login", loginFunc)

	// add in addContext middlware
	mux.Handle("/graphql", requireAuth(graphqlHandler))

	return mux
}

func setupServer() {
	http.ListenAndServe(":8080", setupMux())
}
func graphqlHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// get query
	opts := handler.NewRequestOptions(r)

	// execute graphql query
	params := graphql.Params{
		Schema:         model.SchemaQL, // defined in another file
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		Context:        r.Context(), // pass http.Request.Context() to our graphql object
	}
	result := graphql.Do(params)

	// output JSON
	var buff []byte
	w.WriteHeader(http.StatusOK)
	/*
	if prettyPrintGraphQL {
		buff, _ = json.MarshalIndent(result, "", "\t")
	} else {
		buff, _ = json.Marshal(result)
	}
	*/
	buff, _ = json.Marshal(result)
	w.Write(buff)
}

type Claims struct {
	username string `json: "username"`
	password string   `json:"password"` ////???
	jwt.StandardClaims
}

type requestBody  struct {
username string `json:"username"`
password string `json:"password"`
}

// secret string for signing requests
var jwtSecret = []byte("secret") // make sure you change this to something secure

// key type is not exported to prevent collisions with context keys defined in
// other packages.
type key int

// userAuthKey is the context key for our added struct.  Its value of zero is
// arbitrary.  If this package defined other context keys, they would have
// different integer values.
const userAuthKey key = 0


func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		// load userID & isAdmin into context
		authContext := struct {
			username  string `json:"username"`
			password string   `json:"password"`
		}{
			claims.username,
			claims.password,
		}
		ctx := context.WithValue(r.Context(), userAuthKey, authContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func confirmLogin(requestbody requestBody) (bool, error){
	username := requestbody.username
	password := requestbody.password
	return model.IsUserValid(username, password ) , nil
}
func loginFunc(w http.ResponseWriter, req *http.Request) {
	// get username & password
	decoder := json.NewDecoder(req.Body)
	requestBody := struct {
		username string `json:"username"`
		password string `json:"password"`
	}{}
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	// confirmLogin is up to you to define
	valid, err := confirmLogin(requestBody)
	if valid==false || err != nil {
		http.Error(w, "invalid login", http.StatusUnauthorized)
		return
	}

	//generate token
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	claims := Claims{
		requestBody.username,
		requestBody.password,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:8080",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(jwtSecret)

	//output token
	tokenResponse := struct {
		Token string `json:"token"`
	}{signedToken}
	json.NewEncoder(w).Encode(tokenResponse)
}
