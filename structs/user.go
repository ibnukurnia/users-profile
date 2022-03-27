package structs

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RiskProfile struct {
	Id           int     `json:"-"`
	UserId       int     `json:"-"`
	MmPercent    float32 `json:"mmPercent"`
	BondPercent  float32 `json:"bondPercent"`
	StockPercent float32 `json:"stockPercent"`
}

type UserDetail struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Age         int         `json:"age"`
	RiskProfile RiskProfile `json:"riskProfile"`
}
