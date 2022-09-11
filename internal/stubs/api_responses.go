package stubs

// PostTokenResponse is a dummy JSOn response for getting access token
func PostTokenResponse() []byte {
	return []byte(`
	{
		"token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsInVpZCI6Mn0.eyJpYXQiOjE2MDM4MjQyODMsIm5iZiI6MTYwMzgyNDI4MywiZXhwIjoxNjAzODI3ODgzfQ.ufW8sCrf_W2RFpVvH6zri0l7pJLnkPXCZi1zc10ZvOg",
		"expires_in": 3600
	}
`)
}

// PostCollectResponse is a dummy JSON response for requesting a payment
func PostCollectResponse() string {
	return `
	{
		  "reference": "bcedde9b-62a7-4421-96ac-2e6179552a1a",
		  "ussd_code": "*126# for MTN or #150*50# for ORANGE",
		  "operator": "mtn or orange"
	}`
}

// GetPendingTransactionResponse is a dummy JSON response for the Transaction Status
func GetPendingTransactionResponse() []byte {
	return []byte(`
	{
		"reference": "bcedde9b-62a7-4421-96ac-2e6179552a1a",
		"status": "PENDING",
		"amount": 1,
		"currency": "XAF",
		"operator": "MTN",
		"code": "CP201027T00005",
		"operator_reference": "1880106956"
	}`)
}

// GetSuccessfulTransactionResponse is a dummy JSON response for the Transaction Status
func GetSuccessfulTransactionResponse() []byte {
	return []byte(`
	{
		"reference": "bcedde9b-62a7-4421-96ac-2e6179552a1a",
		"status": "SUCCESSFUL",
		"amount": 1,
		"currency": "XAF",
		"operator": "MTN",
		"code": "CP201027T00005",
		"operator_reference": "1880106956"
	}`)
}

// GetPendingAirtimeTransactionResponse is a dummy JSON response for a pending airtime transaction
func GetPendingAirtimeTransactionResponse() []byte {
	return []byte(`
	{
		"reference": "971e32ae-bb5a-420a-a38a-c2931536609f",
		"external_reference": "5577006791947779410",
		"status": "PENDING",
		"amount": 100,
		"currency": "XAF",
		"operator": "ORANGE_CM",
		"code": "CP220804U0649K",
		"type": "AIRTIME",
		"reason": ""
	}
`)
}

// GetSuccessfullAirtimeTransactionResponse is a dummy JSON response for a successful airtime transaction
func GetSuccessfullAirtimeTransactionResponse() []byte {
	return []byte(`
	{
		"reference": "971e32ae-bb5a-420a-a38a-c2931536609f",
		"external_reference": "5577006791947779410",
		"status": "SUCCESSFUL",
		"amount": 100,
		"currency": "XAF",
		"operator": "ORANGE_CM",
		"code": "CP220804U0649K",
		"type": "AIRTIME",
		"reason": ""
	}
`)
}

// GetBalanceResponse is a dummy JSON response for the transaction balance
func GetBalanceResponse() string {
	return `
	{
		"total_balance": 0,
		"mtn_balance": 0,
		"orange_balance": 0,
		"currency": "XAF"
	}
`
}

// PostTransferResponse is a dummy JSON response for airtime transfer requests
func PostTransferResponse() []byte {
	return []byte(`
	{
		"reference":"26676007-1c31-46d7-9c71-acb031cf0de4",
		"status":"PENDING"
	}
`)
}

// PostWithdrawResponse is a dummy JSON response for withdraw requests
func PostWithdrawResponse() []byte {
	return []byte(`
	{
		"reference":"26676007-1c31-46d7-9c71-acb031cf0de4",
		"status":"PENDING"
	}
`)
}

// GetHistoryResponse is a dummy JSON response for the transaction history
func GetHistoryResponse() string {
	return `
	{
		"data": [
			{
				"datetime": "2021-01-29T09:52:34.876707Z",
				"code": "CP210129D0001P",
				"operator_tx_code": "MP210129.1052.A35072",
				"operator": "Orange",
				"phone_number": "237696546822",
				"description": "Test",
				"external_user": "",
				"amount": 5,
				"charge_amount": 0.05,
				"debit": 0,
				"credit": 4.95,
				"status": "SUCCESSFUL",
				"reference_uuid": "25c63c72-8485-4059-85ad-fdb4bfb26c21"
			},
			{
				"datetime": "2021-01-25T12:44:11.808507Z",
				"code": "CP210125D0001N",
				"operator_tx_code": "2171591856",
				"operator": "MTN",
				"phone_number": "237679587525",
				"description": "Test",
				"external_user": "",
				"amount": 5,
				"charge_amount": 0.05,
				"debit": 0,
				"credit": 4.95,
				"status": "FAILED",
				"reference_uuid": "769dc5c3-1a98-4788-bac4-2daaa49a58b6"
			}
		]
	}
`
}
