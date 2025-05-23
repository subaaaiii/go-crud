package categorycontroller

import (
	"go-crud/entities"
	"go-crud/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.Created_at = time.Now()
		category.Updated_At = time.Now()

		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		category := categorymodel.GetDetail(id)
		data := map[string]any{
			"category": category,
		}
		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var category entities.Category
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.Updated_At = time.Now()

		if ok := categorymodel.Update(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)

	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	if err := categorymodel.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
