package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SetSlackCookies(req *http.Request) {
	cookie1 := http.Cookie{
		Name:  "_ga",
		Value: "GA1.2.512403146.1409452713",
	}
	cookie2 := http.Cookie{
		Name:  "a",
		Value: "2621832485%2C2725022032%2C2817841753%2C2880540297",
	}
	cookie3 := http.Cookie{
		Name:  "a-2621832485",
		Value: "%2BoVYKJaTNZQlNR2XX6ApZRbeBvcjEzgyN7eWRyl8q6%2FLj0yXXOwCUf13BT49gjq6JcWpKLuq0%2F6cpaeiqOvvKQ%3D%3D",
	}
	cookie4 := http.Cookie{
		Name:  "a-2725022032",
		Value: "%2FVvZzuTNU2lbqPgzYsl528rmTd41ZLxX9%2B5xcsD4G6cNYpdvly5Mw1tt%2BdetiDSsUhlbmlgQdoHiST8avtRtnA%3D%3D",
	}
	cookie5 := http.Cookie{
		Name:  "a-2817841753",
		Value: "QDGu9OYAnKkjWmzTcMb7v9BJQHc2%2FZziyS4ej%2FPHgF2StHXuHQjXiC34Ks1I%2FR2wPnRT5jNYiVT5j%2FMrAObOQg%3D%3D",
	}
	cookie6 := http.Cookie{
		Name:  "a-2880540297",
		Value: "j7prH%2BmbALWAoy9Tblh%2BHa88cREjx5CxWpbTDxZcohU1tMuNrfsIr%2BA0TNDlM8tb4BlDwDjTmoqb3oYAALPc7A%3D%3D",
	}
	cookie7 := http.Cookie{
		Name:  "fbm_569627156411038",
		Value: "base_domain=.slack.com",
	}

	req.AddCookie(&cookie1)
	req.AddCookie(&cookie2)
	req.AddCookie(&cookie3)
	req.AddCookie(&cookie4)
	req.AddCookie(&cookie5)
	req.AddCookie(&cookie6)
	req.AddCookie(&cookie7)
}

func SetFormValues(req *http.Request, fname string, lname string, email string) {
	q := req.URL.Query()
	q.Set("first_name", fname)
	q.Set("last_name", lname)
	q.Set("email", email)
	q.Set("token", "xoxs-2331842482-2725022032-2725022108-1e80107e17")
	q.Set("set_active", "true")
	q.Set("_attempts", "1")
	req.URL.RawQuery = q.Encode()
}

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://gophers.slack.com/api/users.admin.invite?t=1414871617&", nil)

	SetSlackCookies(req)
	SetFormValues(req, "Casey", "Rosengren", "caseyrosengren+5@gmail.com")

	// fmt.Println(req.URL)
	// fmt.Println(req.Cookies())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error is: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading of response error is: ", err)
	}

	fmt.Println(string(body))
}
