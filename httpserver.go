package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//	mgo "gopkg.in/mgo.v2"
)

type Person struct {
	FName string
	LName string
	Phone string
	Email string
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

func SetupSecurity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/setupSecurity.gtpl")
		t.Execute(w, nil)
		//	w.Write([]byte(fmt.Sprintf(html)))
	} else {
		r.ParseForm()
	}

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {

		t, _ := template.ParseFiles("static/index.html")
		t.Execute(w, nil) //fmt.Printf(r.URL.Path[1:])
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
	// fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

	}
}

func PostFormDataHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(" 149 error is ", err)
	}

	//person := new(Person)
	org := new(SecurityRoles)
	decoder := schema.NewDecoder()

	err = decoder.Decode(org, r.PostForm)

	if err != nil {
		fmt.Println("159 err0r ", err)
		//	mongo()
	}
	fmt.Println(" 121 before mongo call")
	mongo(*org)
	// fmt.Println(person)
	fmt.Println(org)
	fmt.Println(org.State)

	w.Write([]byte(fmt.Sprintf("org  is:  %v \n", org)))
	//	w.Write([]byte(fmt.Sprintf("\n L9 is : %v \n", org.l2))) */
}

func main() {

	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()
	// r.Handle("/", SayHelloWorld)
	r.HandleFunc("/getformdata", GetFormDataHandler)
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/setupSecurity", SetupSecurity)
	r.HandleFunc("/login", Login)
	r.HandleFunc("/postsecuritydata", PostFormDataHandler)
	// This will serve files under http://localhost:8000/static/<filename>

	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(Dir))))

	//	r.PathPrefix("static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
func mongo(o SecurityRoles) {
	fmt.Println("158 mongo called")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("security").C("organisation")
	fmt.Println("168 before insert statement")
	err = c.Insert(&SecurityRoles{"Admin", o.Orgname, o.Country, o.State, o.City, o.Sname, o.L1, o.L2, o.L3, o.L4, o.L5, o.L6, o.L7, o.L8, o.L9})

	if err != nil {
		fmt.Println("173 Mongo Error ")
		log.Fatal(err)
	}

	result := SecurityRoles{}
	err = c.Find(bson.M{"city": "Bathurst"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("School Name: ", result.Sname)
}
