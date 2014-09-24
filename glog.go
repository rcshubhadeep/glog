package glog

import (
	"appengine"
	"appengine/user"
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("static/tmpl/*"))

type AdminPage struct {
	Admin_email string
	LogoutUrl   string
}

func init() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/admin", Admin)
}

func Index(w http.ResponseWriter, r *http.Request) {
	// err := templates.ExecuteTemplate(w, "indexPage", nil)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(c, "/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func Admin(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	an := AdminPage{"", ""}
	var templ = ""

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	c.Debugf("User present %v", u)
	an.Admin_email = u.Email
	templ = "adminIndex"
	url, err := user.LogoutURL(c, "/admin")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	an.LogoutUrl = url
	err = templates.ExecuteTemplate(w, templ, an)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
