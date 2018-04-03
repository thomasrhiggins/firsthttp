package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
  "strings"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

)

func handler(w http.ResponseWriter, r *http.Request) {
    //http.FileServer(http.Dir("./static"))
    //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {

        t, _ := template.ParseFiles("index.html")
        t.Execute(w, nil)//fmt.Printf(r.URL.Path[1:])
      }else {}

    fmt.Fprintf(w, userlist(r.URL.Path[1:], ))

}

func userlist( urlpath string) string {
  names := make(map[string]string)
names["tom"] = "   toms url is really good anotehreone1" // assign value by key
names["alf"] = "   alf is another url notheron10e"
names["max"] = "    max nother3"
names["login"] = "login(w, r)"

return names[urlpath]

}

func makeEntity() {
        type SecurityRoles struct {
          admin string
          orgname string
          country string
          state  string
          city  string
          sname string
          l1 string
          l2 string
          l3 string
          l4 string
          l5 string
          l6 string
          l7 string
          l8 string

          l9 string

        }
      var S SecurityRoles
      S.admin = "/admin"
      S.orgname = "/NSW Dept Education"
      S.country ="/AU" /* country code lookup dodo */
      S.state = "/NSW"
      S.city = "/Mossvale"
      S.sname ="/Mossvale High School"
      S.l1 = "/Principal"
      S.l2 = "/Head Teachers"
      S.l3 = "/Teachers"
      S.l4 = "/Students"


      fmt.Printf(S.orgname)
}


func main() {

  makeEntity()


  http.Handle("/", http.FileServer(http.Dir("./static")))
  //http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    http.HandleFunc("/setupSecurity", setupSecurity)
    fs := http.FileServer(http.Dir("static"))

    //http.Handle("/", http.StripPrefix("/static/", fs))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    log.Fatal(http.ListenAndServe(":8080", nil))


    //log.Fatal(http.ListenAndServe(":8080", nil))
    //fmt.Printf(log)
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    //fmt.Fprintf(w, "Hello astaxie!") // write data to response
    //fmt.Fprintf(w, userlist(r.URL.Path[1:], ))
    mongo()
    t, _ := template.ParseFiles("static/index.html")
    t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("static/login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
        sayhelloName(w, r)
    }
}

func setupSecurity(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("static/setupSecurity.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
        sayhelloName(w, r)
    }
}


    type Person struct {
        Name  string
        Phone string
    }

    func mongo() {
        session, err := mgo.Dial("localhost")
        if err != nil {
            panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")
        err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
            &Person{"Cla", "+55 53 8402 8510"})
        if err != nil {
            log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)
    }
