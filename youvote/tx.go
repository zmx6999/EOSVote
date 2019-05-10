package youvote

import "github.com/eoscanada/eos-go"

func NewAction(action string, actor string, actionData interface{}) *eos.Action {
	return &eos.Action{
		Account: eos.AN(CONTRACT_NAME),
		Name: eos.ActN(action),
		Authorization: []eos.PermissionLevel{
			{Actor: eos.AN(actor), Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(actionData),
	}
}

func PushAction(action *eos.Action, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	api := eos.New(API_URL)

	keyBag := &eos.KeyBag{}
	if err = keyBag.ImportPrivateKey(privateKey); err != nil {
		return
	}
	api.SetSigner(keyBag)

	txOpts := &eos.TxOptions{}
	if err = txOpts.FillFromChain(api); err != nil {
		return
	}

	tx := eos.NewTransaction([]*eos.Action{action}, txOpts)
	signedTx, packedTx, err := api.SignTransaction(tx, txOpts.ChainID, eos.CompressionNone)
	if err != nil {
		return
	}

	r, err = api.PushTransaction(packedTx)
	if err != nil {
		return
	}

	return
}
