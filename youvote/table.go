package youvote

import (
	"github.com/eoscanada/eos-go"
	"encoding/json"
)

func GetList(table string) (data []map[string]interface{}, err error) {
	abi := eos.New(API_URL)
	r, err := abi.GetTableRows(eos.GetTableRowsRequest{
		Code: CONTRACT_NAME,
		Scope: CONTRACT_NAME,
		Table: table,
		JSON: true,
	})
	if err != nil {
		return
	}

	_data := r.Rows
	json.Unmarshal(_data, &data)
	return
}
