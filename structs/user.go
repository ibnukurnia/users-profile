package structs

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"-"`
}

type UserRegist struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
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
	Password    string      `json:"-"`
	RiskProfile RiskProfile `json:"riskProfile"`
}

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
