package test


import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/dandoh/sdr/app"
	"github.com/dandoh/sdr/model"
	"context"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var initialized = false
var appHandlerFunc http.HandlerFunc
var authContext model.AuthorContext

func TestMain(m *testing.M)  {
	godotenv.Load(".test_env")
	if !initialized {
		appHandlerFunc = app.GraphqlHandlerFunc();
		authContext = model.AuthorContext{
			AuthorID: 1,
		}
		initialized = true
	}
	m.Run();
}

func TestAddGroup(t *testing.T) {
	req, _ := http.NewRequest("POST", "", nil) // TODO - put body
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appHandlerFunc)

	ctx := context.WithValue(req.Context(), "authorContext", authContext);
	handler.ServeHTTP(rr, req.WithContext(ctx))
}

