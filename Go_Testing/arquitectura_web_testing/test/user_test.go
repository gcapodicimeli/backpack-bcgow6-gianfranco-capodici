package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/cmd/server/handler"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/internal/domain"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/internal/users"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockStore users.MockStorage) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	repository := users.NewRepository(&mockStore)
	service := users.NewService(repository)
	u := handler.NewUser(service)

	r := gin.Default()

	pr := r.Group("/users")
	pr.GET("/", u.GetAll)

	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", os.Getenv("TOKEN"))

	return req, httptest.NewRecorder()
}

func TestGetAllUsers(t *testing.T) {
	// Arrange
	mockStorage := users.MockStorage{
		DataMock: []domain.User{
			{Id: 1,
				Name:     "Juan",
				LastName: "Perez",
				Email:    "juan@gmail.com",
				Age:      39,
				Height:   1.78,
				Active:   true,
				Date:     "02-02-2022",
			},
			{Id: 2,
				Name:     "Juana",
				LastName: "Lopez",
				Email:    "juana@gmail.com",
				Age:      79,
				Height:   1.58,
				Active:   true,
				Date:     "02-06-2022",
			},
		},
	}

	r := createServer(mockStorage)
	req, rr := createRequestTest(http.MethodGet, "/users/", "")

	// Act
	resp := &web.Response{}
	r.ServeHTTP(rr, req)

	// Assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, len(mockStorage.DataMock), reflect.ValueOf(resp.Data).Len())
}
