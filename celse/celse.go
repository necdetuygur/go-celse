/* EXAMPLE
package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		SessionSet(w, r, "foo", "bar")
		foo := SessionGet(w, r, "foo")
		fmt.Println(foo)
		fmt.Fprint(w, foo)
	})
	_ = http.ListenAndServe("0.0.0.0:8000", nil)
}
*/

package celse

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var SessionKey string = "NecdetKey1"
var SessionID string = "NecdetID1"
var Store = sessions.NewCookieStore([]byte(SessionKey))

func Set(w http.ResponseWriter, r *http.Request, key string, value string) {
	session, _ := Store.Get(r, SessionID)
	session.Values[key] = value
	err := sessions.Save(r, w)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Get(w http.ResponseWriter, r *http.Request, key string) string {
	session, err := Store.Get(r, SessionID)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return "ERR3"
	}
	return fmt.Sprintf("%v", session.Values[key])
}
