package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project_api/util"

	"github.com/bitly/go-simplejson"
	"github.com/julienschmidt/httprouter"
)

//Users 用户表json
type Users struct {
	ID       int    `form:"id" json:"id"`
	Account  string `form:"account" json:"account"`
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Sex      string `form:"sex" json:"sex"`
	Phone    string `form:"phone" json:"phone"`
	Level    string `form:"level" json:"level"`
}

// InsertNode 增加数据调用前json准备
func InsertNode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var user Users

	err := json.Unmarshal(body, &user)
	if err != nil {
		util.ResponseJSON(w, 1002, util.ErrorCode(1002), simplejson.New())
		return
	}

	userService := &UserService{}
	if !util.Regexphone(user.Phone) {
		util.ResponseJSON(w, -2, util.ErrorCode(-2), simplejson.New())
		return
	}
	id, err := userService.Insertdata(user)
	if err != nil || id <= 0 {
		util.ResponseJSON(w, -1, util.ErrorCode(-1), simplejson.New())
		return
	}
	idJSON := simplejson.New()
	idJSON.Set("id", id)
	util.ResponseJSON(w, 0, util.ErrorCode(0), idJSON)
}

// DeleteNode 删除数据调用前json准备
func DeleteNode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var users Users
	err := json.Unmarshal(body, &users)
	if err != nil {
		util.Logger.Error(err)
		return
	}

	user := UserService{}
	//fmt.Println("==", users.Account, users.Password)
	isExistData := user.IsExistData(users)
	fmt.Println(isExistData)
	if !isExistData {
		util.ResponseJSON(w, -3, util.ErrorCode(-3), simplejson.New())
		return
	}

	errs := user.Deletedata(users)
	if errs != nil {
		util.ResponseJSON(w, -1, util.ErrorCode(-1), simplejson.New())
		return
	}
	util.ResponseJSON(w, 0, util.ErrorCode(0), simplejson.New())
}

// UpdateNode 更新数据前的json准备
func UpdateNode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	users := UserService{}
	var user Users //数据
	err := json.Unmarshal(body, &user)
	if !users.IsExistData(user) {
		util.ResponseJSON(w, -3, util.ErrorCode(-3), simplejson.New())
		return
	}
	if err != nil {
		util.ResponseJSON(w, 1002, util.ErrorCode(1002), simplejson.New())
		return
	}
	if !util.Regexphone(user.Phone) {
		util.ResponseJSON(w, -4, util.ErrorCode(-4), simplejson.New())
		return
	}
	userservice := &UserService{}
	err = userservice.Updatedata(user)

	if err != nil {
		util.ResponseJSON(w, -1, util.ErrorCode(-1), simplejson.New())
		return
	}
	util.ResponseJSON(w, 0, util.ErrorCode(0), simplejson.New())
}

// SeleNode 查找数据
func SeleNode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	userservice := &UserService{}
	res, err := userservice.Seledata()
	if err != nil {
		util.ResponseJSON(w, -1, util.ErrorCode(-1), simplejson.New())
		return
	}
	util.ResponseJSON(w, 0, util.ErrorCode(0), res)
}

//IsExistData 判断数据账号是否存在
func (p *UserService) IsExistData(node Users) bool {
	var count int
	sql := "select count(*) from user where account=? and password=? " //insert into set 比 insert into values 清晰明了，容易查错 ，但是不能批量增加数据
	err := util.Mysqldb.QueryRow(sql, node.Account, node.Password).Scan(&count)
	//fmt.Print(err)
	if err != nil {
		util.Logger.Error(err)
		return false
	}
	//fmt.Println("count:", count)
	if count >= 1 {
		return true
	}
	return false
}

// UserService 增加存储的用户结构体
type UserService struct {
}

// Insertdata 增加数据
func (p *UserService) Insertdata(node Users) (int64, error) {
	sql := "insert into user set account=?,password=?,name=?,sex=?,phone=?,level=?"
	res, err := util.Mysqldb.Exec(sql, node.Account, node.Password, node.Name, node.Sex, node.Phone, node.Level)
	if err != nil {
		util.Logger.Error(err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		util.Logger.Error(err)
		return 0, err
	}
	return id, nil
}

// Seledata 查询数据
func (p *UserService) Seledata() ([]Users, error) {
	sql := "select * from user"
	rows, err := util.Mysqldb.Query(sql)
	defer rows.Close()
	if err != nil {
		util.Logger.Error(err)
	}
	//获取的集合进行遍历存到数组当中
	var users []Users
	for rows.Next() {
		var user Users
		err = rows.Scan(&user.ID, &user.Account, &user.Password, &user.Name, &user.Sex, &user.Phone, &user.Level)
		if err != nil {
			util.Logger.Error(err)
			continue
		}
		users = append(users, user)
	}
	if err != nil {
		util.Logger.Error(err)
		return nil, err
	}
	return users, nil
}

// Updatedata 更新数据
func (p *UserService) Updatedata(node Users) error {
	sql := "update user set name=?,sex=?,phone=?,level=? where account=? and password=?"
	_, err := util.Mysqldb.Exec(sql, node.Name, node.Sex, node.Phone, node.Level, node.Account, node.Password)
	if err != nil {
		util.Logger.Error(err)
		return err
	}
	return nil
}

// Deletedata 删除数据
func (p *UserService) Deletedata(node Users) error {
	sql := "delete from user where account=? and password=?"
	_, err := util.Mysqldb.Exec(sql, node.Account, node.Password)
	if err != nil {
		util.Logger.Error(err)
		return err
	}
	return nil
}
