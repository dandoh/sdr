package auth

import (
	"io/ioutil"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"encoding/json"
	"github.com/dandoh/sdr/util"
	"github.com/dandoh/sdr/model"
	"time"
	"os"
	"github.com/enodata/faker"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string    `json:"email"`
}

func confirmLogin(requestbody LoginRequest) (uint, bool) {
	email := requestbody.Email
	password := requestbody.Password
	return model.GetUserID(email, password)
}

func confirmSignUp(requestbody SignupRequest) bool {
	username := requestbody.Username
	password := requestbody.Password
	email := requestbody.Email
	if !model.IsUserExisted(username, email) {
		avatar := faker.Avatar().String()
		var user model.User = model.User{Name: username, PasswordMD5: util.GetMD5Hash(password), Email: email, Avarta:avatar}
		model.CreateUser(&user)
		return true
	}
	return false
}

func SignupFunc(w http.ResponseWriter, req *http.Request) {
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

func LoginFunc(w http.ResponseWriter, req *http.Request) {
	// get username & password
	jwtSecret := []byte(os.Getenv("SECRET_KEY"));
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
		requestBody.Email,
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
