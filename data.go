// Data Stored convallis
package main

import (
	"fmt"

	//"log"
	"net/http"

	"google.golang.org/appengine/datastore"

	"google.golang.org/appengine" // Required external App Engine library
	//	"google.golang.org/appengine/user"
)

func PostSecurityFormDataHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	// [START new_context]
	ctx := appengine.NewContext(r)
	var tablename string = "securityroles"

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
	key := datastore.NewIncompleteKey(ctx, tablename, nil)
	// [END new_key]
	// [START add_post]
	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, tablename, nil), &s1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var s2 SecurityRoles
	if err = datastore.Get(ctx, key, &s2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stored and retrieved the School named %q", s2.Sname, " and State is ", s2.State)
}

func PostSignupFormDataHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	p1 := Person{
		Fname: r.Form.Get("fname"),
		Lname: r.Form.Get("lname"),
		Email: r.Form.Get("email"),
	}
	fmt.Fprintf(w, "person is ", p1.Fname)
	//uname := r.Form.Get["username"]
	fmt.Fprintf(w, " %s\n  username:", r.Form["username"])
	//fmt.Fprintf(w, "%s\n password:", r.Form["password"])
	//fmt.Fprintf(w, "%s\n ")
	//func Current(c context.Context) *User
}
