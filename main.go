package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	util "github.com/dandoh/sdr/util"

	model "github.com/dandoh/sdr/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/rs/cors"
)

/*
func main() {
	// initialize database


	h := handler.New(&handler.Config{
		Schema: &model.SchemaQL,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)


	http.ListenAndServe(":8080", c)
	defer model.Close();
}
*/

func main() {
	model.Init()
	model.InitType()
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
	rootMux := setupMux();
	c := cors.AllowAll().Handler(rootMux);
	http.ListenAndServe(":8080", c)
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
	fmt.Printf("%+v", result)

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
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// secret string for signing requests
var jwtSecret = []byte("So hello world") // make sure you change this to something secure

func requireAuth(next http.Handler) http.Handler {
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

		fmt.Printf("lau ra: %+v\n", claims)
		fmt.Printf("authenticating: id is %d", claims.UserID)
		// load userID
		authContext := model.AuthorContext{
			AuthorID: claims.UserID,
		}
		ctx := context.WithValue(r.Context(), "authorContext", authContext)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func confirmLogin(requestbody LoginRequest) (uint, bool) {
	username := requestbody.Username
	password := requestbody.Password
	return model.GetUserID(username, password)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginFunc(w http.ResponseWriter, req *http.Request) {
	// get username & password
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	requestBody := LoginRequest{}
	util.PrintBody(req)
	err := json.Unmarshal(bodyBytes, &requestBody)
	fmt.Printf("%+v", requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// confirmLogin is up to you to define
	userID, success := confirmLogin(requestBody)
	if !success {
		http.Error(w, "invalid login", http.StatusUnauthorized)
		return
	}

	//generate token
	expireToken := time.Now().Add(time.Hour * 48).Unix()

	fmt.Printf("user id is %d, completed login", userID)
	claims := Claims{
		userID,
		requestBody.Username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:8080",
			Id:        string(userID),
		},
	}
	fmt.Printf("them vao: %+v\n", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(jwtSecret)

	//output token
	tokenResponse := struct {
		Token string `json:"token"`
	}{signedToken}
	json.NewEncoder(w).Encode(tokenResponse)
}
