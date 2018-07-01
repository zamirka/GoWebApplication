package simpleweb

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func SayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func SayhelloNameToZamir(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Zamir!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("simpleweb/login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()

		token := r.Form.Get("token")
		if token != "" {
			//check token validity
		} else {
			//give error if no token
		}

		// logic part of log in
		username := r.Form.Get("username")
		if len(username) != 0 {
			fmt.Println("username:", r.Form["username"])
		} else {
			fmt.Println("Username field was empty")
		}
		fmt.Println("password:", r.Form["password"])
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			fmt.Println("No proper age provided")
		} else {
			getint, err := strconv.Atoi(r.Form.Get("age"))
			if err != nil {
				fmt.Println("Anyway couldn't convert age to string")
			}

			switch getint {
			case 100:
				fmt.Println("You can not live so long")
			case 0:
				fmt.Println("It's too early for you")
			case 10:
				fmt.Println("You are too young")
			case 50:
				fmt.Println("You are too old")
			default:
				fmt.Println("Your age is OK")
			}
		}
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprint("%x", h.Sum(nil))

		t, _ := template.ParseFiles("simpleweb/upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
