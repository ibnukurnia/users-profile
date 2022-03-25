package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

type RiskProfile struct {
	
	Id int64 `json:"-"`
	UserId int64 `json:"-"`
	MmPercent float32 `json:"mmPercent"`
	BondPercent float32 `json:"bondPercent"`
	StockPercent float32 `json:"stockPercent"`

}


	
var dbClient *sql.DB

func connect() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/user_profile")
	if err != nil {
		panic(err.Error())
	}

	dbClient = db
}


func main() {

	log.Println("Running at Port 8000")
	
	r := mux.NewRouter()

	connect()
	
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Base Api"))
	})

	r.HandleFunc("/users", getUsers).Methods("OPTIONS","GET")
	r.HandleFunc("/users/{id}", getUser).Methods("OPTIONS", "GET")
	r.HandleFunc("/users", createUser).Methods("OPTIONS", "POST")
	
	http.ListenAndServe(":8000", r)

}


func createUser(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/json")
	var user User
	var riskProfile RiskProfile
	err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	ageIndex := 55 - user.Age
	
	if ageIndex >= 30 {
		riskProfile.StockPercent = 72.5
		riskProfile.BondPercent = 21.5
		riskProfile.MmPercent = 100 - (riskProfile.StockPercent+riskProfile.BondPercent)
	}else if ageIndex >= 20{
		riskProfile.StockPercent = 54.5
		riskProfile.BondPercent = 25.5
		riskProfile.MmPercent = 100 - (riskProfile.StockPercent+riskProfile.BondPercent)
	}else if ageIndex < 20{
		riskProfile.StockPercent = 34.5
		riskProfile.BondPercent = 45.5
		riskProfile.MmPercent = 100 - (riskProfile.StockPercent+riskProfile.BondPercent)
	}

	resUser,err := dbClient.Exec("insert into users(name,age) values (?,?)", user.Name, user.Age)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	id,_ := resUser.LastInsertId()

	user.Id = id

	_,err = dbClient.Exec("insert into risk_profile(user_id,bond_percent,stock_percent,mm_percent) values (?,?,?,?)", 
	id, riskProfile.BondPercent, riskProfile.StockPercent, riskProfile.MmPercent)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	

	json.NewEncoder(w).Encode(user)
}


func getUsers(w http.ResponseWriter, r *http.Request)  {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	var users []User

	rows, err := dbClient.Query("SELECT * FROM users LIMIT  ? OFFSET  ?",limit, offset)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next(){
		var user User

		err = rows.Scan(&user.Id, &user.Name, &user.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		users = append(users, user)
	}

	
	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	
}

func getUser(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id:= params["id"]

	log.Println(id)
	
	rows, err := dbClient.Query("SELECT * from risk_profile INNER JOIN users ON users.id = risk_profile.user_id WHERE users.id = ?",id)

	if err != nil {
		w.Write([]byte("Not found"))
		return
	}

	for rows.Next(){
		var user User
		var riskProfile RiskProfile
		
		
		err = rows.Scan(
			&riskProfile.Id, &riskProfile.UserId, &riskProfile.BondPercent, &riskProfile.StockPercent, 
			&riskProfile.MmPercent,&user.Id, &user.Name, &user.Age )

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		type Response struct {
			User interface{} `json:"user"`
			RiskProfile interface{} `json:"risk"`
		}

		res := Response{User: user,RiskProfile: riskProfile}
		
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(res)
	
	}


}

