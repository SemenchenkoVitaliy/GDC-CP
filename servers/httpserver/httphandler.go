package httpserver

import (
	"html/template"
	"net/http"
	"strconv"

	dbDriver "github.com/SemenchenkoVitaliy/GDC-CP/dbDriver/mongo"
	"github.com/SemenchenkoVitaliy/GDC-CP/netutils"
	"github.com/gorilla/mux"
)

var (
	db = dbDriver.NewDatabase()

	fsPublicUrl string
	publicUrl   string
	templates   *template.Template
)

func loadTemplates() {
	templates = template.Must(template.ParseGlob("./HTML/*.gohtml"))
}

func rootGET(w http.ResponseWriter, r *http.Request) {
	lectionUrls, err := db.GetLectionUrls(0, 10, "name")
	if err != nil {
		netutils.InternalError(w, err, "Get top lections urls")
		return
	}

	lections, err := db.GetLectionMultiple(lectionUrls)
	if err != nil {
		netutils.InternalError(w, err, "Get top lections")
		return
	}

	if err = templates.ExecuteTemplate(w, "index", lections); err != nil {
		netutils.InternalError(w, err, "Execute template index")
		return
	}
}

func searchGET(w http.ResponseWriter, r *http.Request) {
	data := struct {
		PublicUrl string
	}{
		PublicUrl: publicUrl,
	}

	if err := templates.ExecuteTemplate(w, "search", data); err != nil {
		netutils.InternalError(w, err, "Execute template search")
		return
	}
}

func lectionGET(w http.ResponseWriter, r *http.Request) {
	lection, err := db.GetLectionSingle(mux.Vars(r)["lection"])
	if err != nil {
		netutils.NotFoundError(w, nil, "No such lection: "+mux.Vars(r)["lection"])
		return
	}

	data := struct {
		Lection   dbDriver.Product
		PublicUrl string
	}{
		Lection:   lection,
		PublicUrl: publicUrl,
	}

	if err = templates.ExecuteTemplate(w, "lection", data); err != nil {
		netutils.InternalError(w, err, "Execute template lection")
		return
	}
}

func adminGET(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	lectionUrls, err := db.GetLectionUrls(page*10, 10, "name")
	if err != nil {
		netutils.InternalError(w, err, "Get top 10 lection urls")
		return
	}

	lections, err := db.GetLectionMultiple(lectionUrls)
	if err != nil {
		netutils.InternalError(w, err, "Get top lections")
		return
	}

	data := struct {
		Lections  []dbDriver.Product
		PublicUrl string
	}{
		Lections:  lections,
		PublicUrl: publicUrl,
	}

	if err = templates.ExecuteTemplate(w, "admin", data); err != nil {
		netutils.InternalError(w, err, "Execute template admin")
		return
	}
}

func adminLectionGET(w http.ResponseWriter, r *http.Request) {
	lection, err := db.GetLectionSingle(mux.Vars(r)["lection"])
	if err != nil {
		netutils.NotFoundError(w, err, "Get lection "+mux.Vars(r)["lection"])
		return
	}

	data := struct {
		Lection   dbDriver.Product
		PublicUrl string
	}{
		Lection:   lection,
		PublicUrl: publicUrl,
	}

	if err = templates.ExecuteTemplate(w, "adminInfo", data); err != nil {
		netutils.InternalError(w, err, "Execute template adminInfo")
		return
	}
}
