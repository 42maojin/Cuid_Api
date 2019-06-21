package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//数据库连接
func connect() *sql.DB {

	// mysqlHost := "192.168.1.1"
	// mysqlDB := "go_test"

	// mysqlDataSoureceName := mysqlHost + mysqlDB

	db, err := sql.Open("mysql", "root:admin@/go_test?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Users 表
type Users struct {
	ID       int    `form:"id" json:"id"`
	Account  string `form:"account" json:"account"`
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Sex      int    `form:"sex" json:"sex"`
	Phone    int    `form:"phone" json:"phone"`
	Level    int    `form:"level" json:"level"`
}

// Response 返回值
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

//增加数据
func insertdata(w http.ResponseWriter, r *http.Request) {
	var response Response
	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	if r.Method == "POST" {

		account := r.FormValue("account")
		password := r.FormValue("password")
		name := r.FormValue("name")
		sex := r.FormValue("sex")
		phone := r.FormValue("phone")
		fmt.Printf("ppp:%s", phone)
		fmt.Printf("sss:%s", sex)
		var count int64

		row := db.QueryRow("select count(*) from user where account=?", account).Scan(&count)
		if err != nil {
			log.Print(err)

			// Info Debug Error Warning Fatal
			// Log.Debug()
			// Log.Error()

		}
		fmt.Println(row)
		fmt.Println("type :%T", row)

		if count == 0 {
			log.Print("没有结果")
			// right, a := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, phone)
			right := regexphone(phone)
			fmt.Println("right :", right)

			if right {
				level := r.FormValue("level")
				fmt.Println(account)
				fmt.Println(password)
				fmt.Println(name)
				_, err = db.Exec(
					"insert into user(account,password,name,sex,phone,level) values(?,?,?,?,?,?)",
					account, password, name, sex, phone, level)
				if err != nil {
					log.Panicln(err)
				}

				response.Status = 1
				response.Message = "Successfully increased data"

				log.Println("数据插入成功")

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)

			} else if !right {
				//fmt.Fprintf(w, "请输入正确的手机号码")
				response.Status = 0
				response.Message = "Insert data failed-请输入正确的手机号码"

				log.Println("数据插入失败")

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			}
		} else {
			response.Status = 0
			response.Message = "Insert data failed-该账户已存在"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}

	} else if r.Method == "GET" {
		fmt.Println("bad request~")
	}
}

//查询数据
func seledata(w http.ResponseWriter, r *http.Request) {
	var user Users
	var response Response
	var arruser []Users

	db := connect()
	defer db.Close()

	row, err := db.Query("select * from user")
	if err != nil {
		log.Print(err)
	}
	for row.Next() {
		if err := row.Scan(&user.ID, &user.Account, &user.Password, &user.Name, &user.Sex, &user.Phone, &user.Level); err != nil {
			log.Fatal(err.Error())
		} else {
			arruser = append(arruser, user)
		}
	}
	response.Status = 1
	response.Message = "Success"
	response.Data = arruser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//更新数据
func updatedata(w http.ResponseWriter, r *http.Request) {
	var response Response
	db := connect()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	Account := r.FormValue("account")
	Password := r.FormValue("password")
	Name := r.FormValue("name")
	Sex := r.FormValue("sex")
	Phone := r.FormValue("phone")
	Level := r.FormValue("level")
	var count int
	errs := db.QueryRow("select count(*) from user where account=? and password=?", Account, Password).Scan(&count)
	if errs != nil {
		log.Print(errs)
	}
	if count == 1 {
		right := regexphone(Phone)
		if right {
			_, err = db.Exec("update user set name=?,sex=?,phone=?,level=? where account=? and password=?", Name, Sex, Phone, Level, Account, Password)
			if err != nil {
				log.Print(err)
			}
			response.Status = 1
			response.Message = "Update data successfully"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			response.Status = 0
			response.Message = "Insert data failed-请输入正确的手机号码"

			log.Println("数据插入失败")

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}

	} else {
		response.Status = 0
		response.Message = "Update data fail=请检查要更新的账号和密码是否正确"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}

//删除数据
func deletedata(w http.ResponseWriter, r *http.Request) {
	var response Response
	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	Account := r.FormValue("account")
	Password := r.FormValue("password")
	var count int
	errs := db.QueryRow("select count(*) from user where account=? and password=?", Account, Password).Scan(&count)
	if errs != nil {
		log.Print(errs)
	}
	if count == 1 {
		_, err = db.Exec("delete from user where account=? and password=?", Account, Password)
		if err != nil {
			log.Print(err)
		}
		response.Status = 1
		response.Message = "Success delete"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response.Status = 0
		response.Message = "fail delete==无该账号"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

//判断手机号码是否正确
func regexphone(phone string) bool {
	right, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, phone)
	return right
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", insertdata).Methods("POST")   //增
	router.HandleFunc("/user", deletedata).Methods("DELETE") //删
	router.HandleFunc("/user", updatedata).Methods("PUT")    //改
	router.HandleFunc("/user", seledata).Methods("GET")      //查
	http.Handle("/", router)
	fmt.Println("Connected to port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
