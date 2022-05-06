package methods

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"testtask/data"
)

func calculateAmount(transactions data.Transactions) float64 {
	hexValues := transactions.ValueList()
	var floatValues = make([]uint64, len(hexValues), len(hexValues))
	for i := 0; i < len(hexValues); i++ {
		floatValues[i], _ = strconv.ParseUint(strings.TrimPrefix(hexValues[i], "0x"), 16, 64)
	}
	var result uint64 = 0
	for i := 0; i < len(floatValues); i++ {
		result += floatValues[i]
	}
	res := float64(result) * math.Pow(10, -18)
	return res
}

func readResponse(blockNumber int) data.EthData {
	var jsonData data.EthData
	url := "https://api.etherscan.io/api?module=proxy&" +
		"action=eth_getBlockByNumber&tag=" +
		fmt.Sprintf("%x", blockNumber) +
		"&boolean=true&apikey=" + viper.GetString("APIKEY")
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(response.Body)
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&jsonData)
	if err != nil {
		log.Fatalln(err)
	}
	return jsonData
}

func getEthData(blockNumber int) (int, float64) {
	jsonData := readResponse(blockNumber)
	transactions := len(jsonData.Result.Transactions)
	amount := calculateAmount(jsonData.Result.Transactions)
	return transactions, amount
}
