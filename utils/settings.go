package utils

import (
	"go-credital/db"

	"gopkg.in/mgo.v2/bson"
)

var menu []bson.M

//InitSettings -> init settings
func InitSettings() {
	initMenu()
}

func initMenu() {
	pannelRepo := db.InitMongoRepo("sys_menu")
	menuRaw, _ := pannelRepo.FindMenu(nil)

	parent := []bson.M{}
	index := map[string]int{}
	out := map[string][]interface{}{}

	i := 0
	for _, v := range menuRaw {
		if v["parent"].(string) == "" {
			parent = append(parent, v)
			index[v["name"].(string)] = i
			i++
		} else {
			out[v["parent"].(string)] = append(out[v["parent"].(string)], v)
		}
	}

	for k, v := range out {
		temp := parent[index[k]]
		temp["child"] = v
		parent[index[k]] = temp
		//log.Println(temp, v)
	}

	menu = parent

}

// GetMenu -> get menu settings
func GetMenu() []bson.M {
	return menu
}
