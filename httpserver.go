package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os" //for handler testing
	"path/filepath"
	"strings"
	"time"

	"google.golang.org/appengine" // Required external App Engine library
	"google.golang.org/appengine/user"
	// packages written by me
	"templmanager"
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

type UserData struct {
	Name        string
	City        string
	Nationality string
}

type SkillSet struct {
	Language string
	Level    string
}

type SkillSets []*SkillSet

type Configuration struct {
	LayoutPath  string
	IncludePath string
}

var UEmail string = "Login"
var tpl *template.Template
var tlates *template.Template
var newtlates *template.Template
var FileToDisplay string = "ui/show.gohtml"

// template_path := filepath.Join(filepath.Dir(root_path), "templates")

func init() {

	tpl = template.Must(template.ParseGlob("template/*"))
	log.Println("tpl name is :", tpl.DefinedTemplates(), "tpl.name = ", tpl.Name())
}

//var SetupsecurityTmp = parseTemplate("/static/setupSecurity.gtpl")

//listTmpl   = parseTemplate("list.html")
// 	return listTmpl.Execute(w, r, books)
func main() {
	var format string = time.RFC1123
	th := timeHandler(format)
	loadConfiguration("config.json")
	templmanager.LoadTemplates()
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
	http.HandleFunc("/postSignup", PostSignupFormMessage) //postSignup

	http.HandleFunc("/postsecuritydata", PostSecurityFormDataHandler)
	http.Handle("/test", th)
	//	http.Handle("/tParse", templatehandler)
	// This will serve files under http://localhost:8000/static/<filename>
	http.HandleFunc("/", index)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/skillset", skillSet)

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
		tpl.ExecuteTemplate(w, "setupSecurity.gohtml", SetUserContext(r))
		//tpl, err := template.ParseFiles("templates/setupSecurity.gohtml")

		/*	if err != nil {
				log.Println("toms parsefiles temlpate error ", err)
				panic(err)
			}
			tpl.Execute(w, nil)
			if err != nil {
				log.Println("toms execute error ", err)

				panic(err)
			}
		*/
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

func PostSignupFormMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method) //get request method
	if r.Method == "POST" {
		log.Println("Post Signup Form execution - file to display ", FileToDisplay)
		FileToDisplay = "ui/postSignupMessage.gohtml"
		log.Println("Post Signup Form execution - file to display ", FileToDisplay)

		var tmpl = `<html>
<head>
    <title>Hello World HEader!</title>
`
		var ftds = template.Must(template.ParseFiles(FileToDisplay))
		log.Println("tplates names is :", ftds.DefinedTemplates())
		//tlates.AddParseTree(FileToDisplay, tpl.Tree)
		log.Println("tpl name after pass tree :", ftds.DefinedTemplates())
		err := ftds.ExecuteTemplate(w, "postSignup", SetUserContext(r))
		if err != nil {
			panic(err)
		} // parsing of template string
		//t.ExecuteTemplate(w, "index.gohtml", SetUserContext(r))
		log.Println("this is tpl ", tmpl)
		// t, _ := template.New(FileToDisplay) //setp 1
		//t.Execute(w, "Hello World!")        //step 2t.execute()
		//FileTemplateParseHandler(FileToDisplay) //
		//tpl.ExecuteTemplate(w, "signup.gohtml", SetUserContext(r))
		//	tpl, _ := template.ParseFiles("/templates/signup.old.gohtml")
		//	tpl.Execute(w, SetUserContext(r))
	} else {
		r.ParseForm()
		// logic part of log in

	}
}

func TestTemplateHandler(w http.ResponseWriter, r *http.Request) {
	tx := template.New("crap")                                                       // Create a template.
	tx, _ = template.ParseFiles("ui/show.gohtml", "templates/header-include.gohtml") // Parse template file.
	log.Println("Tx name is :", tx.DefinedTemplates())
	pattern := filepath.Join("templates/", "*.gohtml")
	log.Println("template pattern ", pattern)
	tx.ExecuteTemplate(w, "show.gohtml", SetUserContext(r)) //tx.Execute(w, SetUserContext(r)) // merge.

	//tpl.ExecuteTemplate(w, "index.gohtml", SetUserContext(r))
}

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
		tpl.Execute(w, "index.gohtml")
	}
	return http.HandlerFunc(fn)
}

func FileTemplateParseHandler(FileToDisplay string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//fileToDisplay =  ui/show.gohtml//	var fileToAddToTemplate string
		log.Println("FileTemplateParseHandler(FileToDisplay", FileToDisplay)
		var allTemplateFiles []string
		allTemplateFiles = append(allTemplateFiles, FileToDisplay)
		log.Println(" ")
		log.Println("Template string files inital load ", allTemplateFiles)
		log.Println(" ")
		files, err := ioutil.ReadDir("./templates")
		if err != nil {
			fmt.Println(err)

		}
		for _, file := range files {
			filename := file.Name()
			if strings.HasSuffix(filename, ".gohtml") {
				allTemplateFiles = append(allTemplateFiles, "templates/"+filename)

			}

			log.Println("Template files build ", allTemplateFiles)
		}

		tlates := template.New("crap1")

		tlates, err = template.ParseFiles(allTemplateFiles...)
		if err != nil {
			log.Println("firstemplate load failed -----------", err)

		}
		log.Println("final defined templates ", tlates.DefinedTemplates())

		log.Println("final defined FILETODISPLAY  ", FileToDisplay)
		log.Println("Base = ", Base(FileToDisplay))
		FileToDisplay = Base(FileToDisplay)
		tlates.ExecuteTemplate(w, FileToDisplay, SetUserContext(r))

	}
	return http.HandlerFunc(fn)
}
func loadConfiguration(fileName string) {
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("error:", err)
	}
	log.Println("layout path: ", configuration.LayoutPath)
	log.Println("include path: ", configuration.IncludePath)
	templmanager.SetTemplateConfig(configuration.LayoutPath, configuration.IncludePath)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := templmanager.RenderTemplate(w, "index.tmpl", nil)
	if err != nil {
		log.Println(err)
	}
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	userData := &UserData{Name: "Asit Dhal", City: "Bhubaneswar", Nationality: "Indian"}
	err := templmanager.RenderTemplate(w, "aboutme.tmpl", userData)
	if err != nil {
		log.Println(err)
	}
}

func skillSet(w http.ResponseWriter, r *http.Request) {
	skillSets := SkillSets{&SkillSet{Language: "Golang", Level: "Beginner"},
		&SkillSet{Language: "C++", Level: "Advanced"},
		&SkillSet{Language: "Python", Level: "Advanced"}}
	err := templmanager.RenderTemplate(w, "skillset.tmpl", skillSets)
	if err != nil {
		log.Println(err)
	}
}
