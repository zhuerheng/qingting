package query

import (
	"../database"
	"net/http"
	"strconv"
)

func Query(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		date, ok1 := r.PostForm["datetime"]
		name, ok2 := r.PostForm["name"]
		if ok1 && ok2 && name[0] != "" && date[0] != "" {
			num, found, err := database.DbQueryNum(date[0], name[0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(err)
			}
			if !found {
				templates["query"].Execute(w, "没有找到相关数据!")
				return
			}
			templates["query"].Execute(w, "日期："+date[0]+" 名称："+name[0]+" 消息数量："+strconv.Itoa(num))

		} else {
			templates["query"].Execute(w, "date 或 name 不能为空!")
			return
		}
	} else {
		templates["query"].Execute(w, nil)
	}
}
