package products

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	repository := NewRepository()
	service := NewService(repository)
	u := NewHandler(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("", u.GetProducts)

	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("token", os.Getenv("TOKEN"))

	return req, httptest.NewRecorder()
}

func TestGetAllProducts(t *testing.T) {
	// Arrange
	initialDatabase := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}

	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products?seller_id=1", "")

	// Act
	var resp []Product
	r.ServeHTTP(rr, req)

	// Assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initialDatabase, resp)
	assert.Equal(t, len(initialDatabase), len(resp))
}

func TestGetAllProductsFail_StatusBadRequest(t *testing.T) {
	// Arrange
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products?seller_id=", "")

	// Act
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAllProductsFail_StatusInternalServer(t *testing.T) {
	// Arrange
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products?seller_id=#", "")

	// Act
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
