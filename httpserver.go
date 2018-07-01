package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"google.golang.org/appengine" // Required external App Engine library
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

type Person struct {
	Fname       string
	Lname       string
	Phone       string
	Email       string
	NotLoggedIn bool
	Message     string
}

type SecurityRoles struct {
	Admin   string
	Orgname string
	Country string
	State   string
	City    string
	Sname   string
	L1      string
	L2      string
	L3      string
	L4      string
	L5      string
	L6      string
	L7      string
	L8      string
	L9      string
}

var UEmail string = "Login"
var tpl *template.Template

// template_path := filepath.Join(filepath.Dir(root_path), "templates")
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

//var SetupsecurityTmp = parseTemplate("/static/setupSecurity.gtpl")

//listTmpl   = parseTemplate("list.html")
// 	return listTmpl.Execute(w, r, books)
func main() {

	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	//r := mux.NewRouter()
	// r.Handle("/", SayHelloWorld)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/getformdata", GetFormDataHandler)
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/setupSecurity", SetupSecurity)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/postSignup", PostSignupFormDataHandler)

	http.HandleFunc("/postsecuritydata", PostSecurityFormDataHandler)
	http.HandleFunc("/testtemplate", TestTemplateHandler)

	// This will serve files under http://localhost:8000/static/<filename>

	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	// http.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	appengine.Main()
	// r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(Dir))))

	// [START new_key]

}
func SetUserContext(r *http.Request) Person {
	ctx := appengine.NewContext(r) //This code setups the username in the header
	u := user.Current(ctx)
	UEmail = fmt.Sprintf("%v", u) // it is here because Context changes before templates are excuted
	p1 := new(Person)
	p1.Email = UEmail

	if u == nil {

		p1.Email = "Login"
		p1.Message = "Sign Up"
		p1.NotLoggedIn = true

		//	UEmail = "Login"
	} else {
		p1.Message = "Welcome"
	}
	return *p1
}

func SetupSecurity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		// tpl.ExecuteTemplate(w, "setupSecurity.gohtml", SetUserContext(r))
		t, _ := template.ParseFiles("ui/setupSecurity.gohtml")
		t.Execute(w, nil)
		// not included in olde script 	w.Write([]byte(fmt.Sprintf(html)))
	} else {
		r.ParseForm()
	}

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r) //This code setups the username in the header
	u := user.Current(ctx)
	UEmail = fmt.Sprintf("%v", u) // it is here because Context changes before templates are excuted
	p1 := new(Person)
	p1.Email = UEmail

	if u == nil {

		p1.Email = "Login"
		p1.Message = "Sign Up"
		p1.NotLoggedIn = true

		//	UEmail = "Login"
	} else {
		p1.Message = "Welcome"
	}

	//get request method
	if r.Method == "GET" {

		t, _ := template.ParseFiles("static/index.html")
		t.Execute(w, p1) //fmt.Printf(r.URL.Path[1:])
	} else {
	}

	//	fmt.Fprintf(w, userlist(r.URL.Path[1:]))

}

func GetFormDataHandler(w http.ResponseWriter, r *http.Request) {

	html := `<h1>Contact  : </h1>

               // replace example.com to your machine domain name or localhost
               <form action="http://localhost:8080/process_form_data" method="post">
                <div>
                 <label>Name : </label>
                 <input type="text" name="name" id="name" >
                </div>
                <div>
                 <label>Phone : </label>
                 <input type="text" name="phone" id="phone" >
                </div>
                <div>
                  <input type="submit" value="Send">
                </div>
              </form>`

	w.Write([]byte(fmt.Sprintf(html)))
}
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	// UEmail = fmt.Sprintf("%v", u)

	log.Printf("current user: %v ", UEmail)
	log.Printf("%#v", u)
	log.Printf("%#v", ctx)

	if u == nil {

		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)

	}

	url, _ := user.LogoutURL(ctx, "/")
	fmt.Fprintf(w, `<a Welcome, %s\n ! (<a href="%s">sign out</a>`, u, url)

}

func Signup(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "signup.gohtml", SetUserContext(r))
		//	tpl, _ := template.ParseFiles("/templates/signup.old.gohtml")
		//	tpl.Execute(w, SetUserContext(r))
	} else {
		r.ParseForm()
		// logic part of log in

	}

}

func TestTemplateHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "index.gohtml", SetUserContext(r))
}
func PostSignupFormDataHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	p1 := Person{
		Fname: r.Form.Get("fname"),
		Lname: r.Form.Get("lname"),
		Email: r.Form.Get("email"),
	}
	log.Printf("person is ", p1.Fname)
	//uname := r.Form.Get["username"]
	//fmt.Fprintf(w, " %s\n  username:", r.Form["username"])
	//fmt.Fprintf(w, "%s\n password:", r.Form["password"])
	//fmt.Fprintf(w, "%s\n ")
	//func Current(c context.Context) *User
}
func PostSecurityFormDataHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	// [START new_context]
	ctx := appengine.NewContext(r)
	// [END new_context]
	s1 := SecurityRoles{

		Admin:   r.Form.Get("admin"),
		Orgname: r.Form.Get("orgname"),
		Country: r.Form.Get("country"),
		State:   r.Form.Get("state"),
		City:    r.Form.Get("city"),
		Sname:   r.Form.Get("sname"),
		L1:      r.Form.Get("l1"),
		L2:      r.Form.Get("l2"),
		L3:      r.Form.Get("l3"),
		L4:      r.Form.Get("l4"),
		L5:      r.Form.Get("l5"),
		L6:      r.Form.Get("l6"),
		L7:      r.Form.Get("l7"),
		L8:      r.Form.Get("l8"),
		L9:      r.Form.Get("l9"),
	}
	key := datastore.NewIncompleteKey(ctx, "securityroles", nil)
	// [END new_key]
	// [START add_post]
	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "securityroles", nil), &s1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var s2 SecurityRoles
	if err = datastore.Get(ctx, key, &s2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stored and retrieved the School named %q", s2.Sname)
}
