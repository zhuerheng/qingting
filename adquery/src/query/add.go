package query

import (
	"../database"
	"fmt"
	"html/template"
	"net/http"
)

var templates map[string]*template.Template

func init() {
	fmt.Println("init ok!")
	templates = make(map[string]*template.Template)
	for _, tmpl := range []string{"add", "query"} {
		t, err := template.ParseFiles("../tmpl/" + tmpl + ".html")
		if err != nil {
			panic(err)
		}
		templates[tmpl] = t
	}
}

func add(name string, url string) (int, error) { //int 1 for name 2 for url
	//检查要添加的字段是否已经存在
	_, found, err := database.DbQueryFromAd_name(name, "name", "url")
	if err != nil {
		return 0, err
	}
	if found {
		return 1, nil
	}
	_, found, err = database.DbQueryFromAd_name(url, "url", "name")
	if err != nil {
		return 0, err
	}
	if found {
		return 2, nil
	}

	//添加字段
	err = database.DbAdd(name, url)
	return 0, err
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		name, ok1 := r.PostForm["name"]
		url, ok2 := r.PostForm["url"]
		if ok1 && ok2 && name[0] != "" && url[0] != "" {
			if temp, err := add(name[0], url[0]); err != nil {
				http.Error(w, err.Error(), http.StatusOK)
				panic(err)
			} else if temp == 1 {
				templates["add"].Execute(w, "name: "+name[0]+" 已经存在！")
				return
			} else if temp == 2 {
				templates["add"].Execute(w, "url: "+url[0]+" 已经存在!")
				return
			}
		} else {
			templates["add"].Execute(w, "name 或 url 不能为空!")
			return
		}
		templates["add"].Execute(w, "成功添加 name:"+name[0])
	} else {
		templates["add"].Execute(w, nil)
	}
}
