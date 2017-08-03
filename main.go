package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/dandoh/sdr/util"

	"github.com/dandoh/sdr/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model.Init()
	model.InitType()

	setupServer()
}

var jwtSecret = []byte(os.Getenv("SECRET_KEY"));

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()

	// graphql Handler
	graphqlHandler := http.HandlerFunc(graphqlHandlerFunc)

	// login Handler
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/login", loginFunc)
	mux.HandleFunc("/signup", signupFunc)
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
		Schema:         model.QLSchema, // defined in another file
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
		fmt.Println("Chuan bi");
		fmt.Println(jwtSecret);
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

func confirmSignUp(requestbody SignupRequest) bool {
	username := requestbody.Username
	password := requestbody.Password
	email := requestbody.Email
	if (model.IsUserExisted(username, email) == false) {
		var user model.User = model.User{Name: username, PasswordMD5: util.GetMD5Hash(password), Email: email}
		model.CreateUser(&user)
		return true
	}
	return false
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string    `json:"email"`
}

func signupFunc(w http.ResponseWriter, req *http.Request) {
	// get username & password
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	requestBody := SignupRequest{}
	util.PrintBody(req)
	err := json.Unmarshal(bodyBytes, &requestBody)
	fmt.Printf("%+v", requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	success := confirmSignUp(requestBody)

	var result string = "fail"
	if (success) {
		result = "success"
	}

	statusResponse := struct {
		Status string `json:"status"`
	}{Status: result}

	json.NewEncoder(w).Encode(statusResponse)
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
	expireToken := time.Now().Add(time.Hour * 24 * 30).Unix()

	claims := Claims{
		userID,
		requestBody.Username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:8080",
			Id:        string(userID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("Chuan bi");
	fmt.Println(jwtSecret);
	signedToken, _ := token.SignedString(jwtSecret)

	//output token
	tokenResponse := struct {
		Token  string `json:"token"`
		UserID uint `json:"userId"`
	}{signedToken, userID}
	json.NewEncoder(w).Encode(tokenResponse)
}
