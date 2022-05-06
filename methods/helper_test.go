package methods

import (
	"testing"
	"testtask/data"
)

func Test_CalculateAmountZeroVal(t *testing.T) {

	t1 := data.Transactions{
		data.Transaction{
			Value: "0",
		},
		data.Transaction{
			Value: "0",
		},
		data.Transaction{
			Value: "0",
		},
	}

	got := calculateAmount(t1)
	check := 0.0

	if got != check {
		t.Errorf("got %v wanted %v", got, check)
	}

}

func Test_ReadResponseIdValue(t *testing.T) {
	got := readResponse(1)
	check := data.EthData{
		Jsonrpc: "",
		ID:      0,
		Result:  data.Result{},
	}
	if got.ID == check.ID {
		t.Errorf("got %v wanted %v", got, check)
	}
}

func Test_GetEthData(t *testing.T) {
	got1, got2 := getEthData(11509797)
	check1, check2 := 155, 2.285405

	if got1 != check1 && got2 != check2 {
		t.Errorf("values does not match")
	}
}
