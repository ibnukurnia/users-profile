package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user-profile/dbConnection"
	"user-profile/structs"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user structs.UserRegist
	var riskProfile structs.RiskProfile

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	isUserExist,_ := dbConnection.DbClient.Query("select name from users where name = ?", user.Name)

	type UserExist struct {
		Name string
	}
	var userExist UserExist
	for isUserExist.Next(){
		err = isUserExist.Scan(&userExist.Name)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	if userExist.Name != ""{
		fmt.Fprintln(w,"user already exist")
		return
	}

	ageIndex := 55 - user.Age

	if ageIndex >= 30 {
		riskProfile.StockPercent = 72.5
		riskProfile.BondPercent = 21.5
		riskProfile.MmPercent = 100 - (riskProfile.StockPercent + riskProfile.BondPercent)
	} else if ageIndex >= 20 {
		riskProfile.StockPercent = 54.5
		riskProfile.BondPercent = 25.5
		riskProfile.MmPercent = 100 - (riskProfile.StockPercent + riskProfile.BondPercent)
	} else if ageIndex < 20 {
		riskProfile.StockPercent = 34.5
		riskProfile.BondPercent = 45.5
		riskProfile.MmPercent = 100 - (riskProfile.StockPercent + riskProfile.BondPercent)
	}

	hashPassword,_:=bcrypt.GenerateFromPassword([]byte(user.Password),10)

	resUser, err := dbConnection.DbClient.Exec("insert into users(name,age,password) values (?,?,?)", user.Name, user.Age,string(hashPassword))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id, _ := resUser.LastInsertId()

	user.Id = int(id)

	_, err = dbConnection.DbClient.Exec("insert into risk_profile(user_id,bond_percent,stock_percent,mm_percent) values (?,?,?,?)",
		id, riskProfile.BondPercent, riskProfile.StockPercent, riskProfile.MmPercent)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Fprintln(w,"success crating user")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	
	limitQuery :=  r.URL.Query().Get("limit")
	offsetQuery := r.URL.Query().Get("offset")

	var limit int
	var offset int

	if limitQuery != "" {
		limit,_ = strconv.Atoi(limitQuery)
	}else{
		limit = 10
	}

	if offsetQuery != "" {
		offset,_ = strconv.Atoi(offsetQuery)
	}else{
		offset = 0
	}

	var users []structs.User

	rows, err := dbConnection.DbClient.Query("SELECT * FROM users LIMIT  ? OFFSET  ?", limit, offset)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var user structs.User

		err = rows.Scan(&user.Id, &user.Name,&user.Password, &user.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if users != nil{
		json.NewEncoder(w).Encode(users)
	}else{
		fmt.Fprintf(w,"No Data")
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	rows, err := dbConnection.DbClient.Query("SELECT * from risk_profile INNER JOIN users ON users.id = risk_profile.user_id WHERE users.id = ?", id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var user structs.UserDetail

	for rows.Next() {
		
		err = rows.Scan(
			&user.RiskProfile.Id, &user.RiskProfile.UserId, &user.RiskProfile.BondPercent, &user.RiskProfile.StockPercent,
			&user.RiskProfile.MmPercent, &user.Id, &user.Name,&user.Password ,&user.Age)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	if user.Id != 0 {
		json.NewEncoder(w).Encode(user)
	}else{
		fmt.Fprintf(w,"User with id %s not found",id)
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user structs.UserLogin

	var userDetail structs.UserDetail

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	isUserExist,_ := 
	dbConnection.DbClient.Query("SELECT * from risk_profile INNER JOIN users ON users.id = risk_profile.user_id WHERE name = ?", 
	user.Name)
	
	for isUserExist.Next(){
		err = isUserExist.Scan(&userDetail.RiskProfile.Id, &userDetail.RiskProfile.UserId, &userDetail.RiskProfile.BondPercent, &userDetail.RiskProfile.StockPercent,
			&userDetail.RiskProfile.MmPercent, &userDetail.Id, &userDetail.Name,&userDetail.Password ,&userDetail.Age)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	if userDetail.Name == ""{
		fmt.Fprintln(w,"invalid user or password")
		return
	}

	matched := bcrypt.CompareHashAndPassword([]byte(userDetail.Password),[]byte(user.Password))

	if matched != nil {
		fmt.Fprintln(w,"invalid user or password")
		return
	}
	json.NewEncoder(w).Encode(userDetail)
}