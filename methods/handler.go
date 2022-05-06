package methods

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

//GetTransactionsAmount e.GET("/api/block/:block_number/total")
func GetTransactionsAmount(c echo.Context) error {
	callback := c.QueryParam("callback")
	blockNumber, err := strconv.Atoi(c.Param("block_number"))
	if err != nil {
		log.Fatalln(err)
	}
	var payload struct {
		Transactions int     `json:"transactions"`
		Amount       float64 `json:"amount"`
	}
	payload.Transactions, payload.Amount = getEthData(blockNumber)
	return c.JSONP(http.StatusOK, callback, payload)
}
