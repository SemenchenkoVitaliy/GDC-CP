package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "strconv"
	// "strings"

	dbDriver "github.com/SemenchenkoVitaliy/GDC-CP/dbDriver/mongo"
	"github.com/SemenchenkoVitaliy/GDC-CP/netutils"
	"github.com/SemenchenkoVitaliy/GDC-CP/utils"
	"github.com/gorilla/mux"
)

var (
	db = dbDriver.NewDatabase()

	fsPublicUrl string
	publicUrl   string
)

func rootGET(w http.ResponseWriter, r *http.Request) {
	lectionUrls, err := db.GetLectionUrls(0, 10, "upddate")
	if err != nil {
		netutils.InternalError(w, err, "Get top lections urls")
		return
	}

	lections, err := db.GetLectionMultiple(lectionUrls)
	if err != nil {
		netutils.InternalError(w, err, "Get top lections")
		return
	}

	fmt.Println(lections)

	result, err := json.Marshal(lections)
	if err != nil {
		netutils.InternalError(w, err, "JSON convert in roorGET")
		return
	}

	w.Header().Set("Content-Type", "text/json")
	w.Write([]byte(result))
}

func lectionGET(w http.ResponseWriter, r *http.Request) {
	data, err := db.GetLectionSingle(mux.Vars(r)["lection"])
	if err != nil {
		netutils.NotFoundError(w, nil, "No such lection: "+mux.Vars(r)["lection"])
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		netutils.InternalError(w, err, "JSON convert in lectionAllGET")
		return
	}
	w.Header().Set("Content-Type", "text/json")
	w.Write([]byte(result))
}

func rootPOST(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	switch action {
	case "add":
		name := r.FormValue("name")

		lection := dbDriver.Product{
			Name: name,
		}

		err := db.AddLection(lection)
		if err != nil {
			http.Error(w, "no such action", 400)
		}
	default:
		http.Error(w, "no such action: "+action, 400)
		return
	}
	w.WriteHeader(200)
}

func lectionPOST(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	switch action {
	case "remove":
		db.RemoveLection(mux.Vars(r)["lection"])
	case "changeName":
		db.ChangeNameLection(mux.Vars(r)["lection"], r.FormValue("name"))
	case "changeLecturer":
		db.ChangeLecturerLection(mux.Vars(r)["lection"], r.FormValue("name"))
	case "changeCourse":
		db.ChangeCourseLection(mux.Vars(r)["lection"], r.FormValue("name"))
	case "addTag":
		db.AddTagLection(mux.Vars(r)["lection"], r.FormValue("name"))
	case "remTag":
		db.RemoveTagLection(mux.Vars(r)["lection"], r.FormValue("name"))
	case "addGroup":
		db.AddGroupLection(mux.Vars(r)["lection"], r.FormValue("name"))
	case "remGroup":
		db.RemoveGroupLection(mux.Vars(r)["lection"], r.FormValue("name"))
	case "addContent":
		r.ParseMultipartForm(32 << 20)
		fhs := r.MultipartForm.File["file"]
		for _, header := range fhs {
			if header.Filename == "" {
				return
			}
			file, err := header.Open()
			defer file.Close()

			if err != nil {
				utils.Log(err, "form file in addCover in"+mux.Vars(r)["lection"])
				http.Error(w, err.Error(), 500)
			}

			var buf bytes.Buffer
			io.Copy(&buf, file)

			filePath := fmt.Sprintf(
				"/data/%v/%v",
				mux.Vars(r)["lection"],
				header.Filename,
			)
			mainServer.WriteFile(filePath, buf.Bytes())

			db.AddContentLection(mux.Vars(r)["lection"], header.Filename)
		}
	case "remContent":
		db.RemoveContentLection(mux.Vars(r)["lection"], r.FormValue("name"))
	default:
		http.Error(w, "no such action: "+action, 400)
		return
	}

	w.WriteHeader(200)
}
