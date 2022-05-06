package methods

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetTransactionsAmount(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/block/:block_number/total")
	c.SetParamNames("block_number")
	c.SetParamValues("11509797")

	if assert.NoError(t, GetTransactionsAmount(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
