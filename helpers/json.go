package helpers

func AmountJson(amount float32) map[string]float32 {
	return map[string]float32{"amount": amount}
}

func CashbackJson(amount float32) map[string]float32 {
	return map[string]float32{"amount": 0, "cashback": amount}
}
