package controllers

import (
	mo "go-credital/db/mongo"
	"go-credital/models/mongo"
	"go-credital/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// M - default data type
type M map[string]interface{}

// Login controler
func Login(c echo.Context) error {
	// model := mongo.SysUser{
	// 	Name:     "youngky",
	// 	Password: "inipassword",
	// 	Username: "banghaji",
	// 	Group:    "lawak",
	// }

	// sess := mo.GetConnection()
	// panelRepo := mo.InitMongoRepo(sess, "sys_user")
	// t := M{
	// 	"name": "youngky",
	// }
	// o := mongo.SysUser{}
	// user, err := panelRepo.FindOne(t, o)
	// bsonBytes, _ := bson.Marshal(user)
	// bson.Unmarshal(bsonBytes, &o)

	// log.Println(o.ID.Hex())

	// if err != nil {
	// 	log.Println(err)
	// }

	//fields := utils.GetFieldsName(model)
	//log.Println(fields)
	// sini
	data := M{}
	return c.Render(http.StatusOK, "login", data)
}

// DoLogin controller
func DoLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	model := mongo.SysUser{}
	pannelRepo := mo.InitMongoRepo("sys_user")

	params := bson.M{
		"username": username,
	}

	sysUserBson, _ := pannelRepo.FindOne(params, &model)

	data := M{
		"error":    "Username dan Password tidak sama",
		"username": username,
	}

	if sysUserBson == nil {
		data["error"] = "Anda belum terdaftar"
	} else {
		bsonBytes, _ := bson.Marshal(sysUserBson)
		bson.Unmarshal(bsonBytes, &model)
		if password == model.Password {
			utils.SetSession(c, "user", model)

			return c.Redirect(http.StatusFound, "/dashboard")
		}
	}

	return c.Render(http.StatusOK, "login", data)
}

// Dashboard - dashboard controller
func Dashboard(c echo.Context) error {
	sess := utils.GetSession(c)
	user := sess.Values["user"].(mongo.SysUser)
	log.Println(user.Name)
	data := M{"message": "Hello World!"}
	return c.Render(http.StatusOK, "dashboard", data)
}

// Logout - logout controller
func Logout(c echo.Context) error {
	utils.DeleteSession(c)
	return c.Redirect(http.StatusFound, "/login")
}
