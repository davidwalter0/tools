// https://gist.githubusercontent.com/Rabbit52/a8a44c3c4cd514052952/raw/4ed391af30d39cde52c71bb7e1b5056189edb998/set_cookiejar.go
package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/davidwalter0/tools/trace/httptrace"
)

func main() {
	var MapKey = "Set-Cookie"
	var MapKey2 = "vpn0.me"
	jar, _ := cookiejar.New(nil)
	var cookies = []*http.Cookie{}
	// var cookieValue string

	var cookieValue = "MTUyNzQ3OTE1MnxEdi1CQkFFQ180SUFBUkFCRUFBQUpmLUNBQUVHYzNSeWFXNW5EQThBRFdGMWRHaGxiblJwWTJGMFpXUUVZbTl2YkFJQ0FBRT18jT6qXznIa7VOltg_b0j9XQb8IVx2ABqp7JNkVs_QN5g="
	cookie := &http.Cookie{
		Name:   MapKey2,
		Value:  cookieValue,
		Path:   "/",
		Domain: "vpn0.me",
	}
	cookies = append(cookies, cookie)
	u, _ := url.Parse("http://weibo.cn/search/?vt=4")
	jar.SetCookies(u, cookies)
	fmt.Println(jar.Cookies(u))
	client := &http.Client{
		Jar: jar,
	}
	postData := url.Values{}
	postData.Set("keyword", "尹相杰")
	postData.Set("smblog", "搜微博")
	r, _ := http.NewRequest("POST", "http://vpn0.me:8080/login", strings.NewReader(postData.Encode()))
	// r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Cookie", fmt.Sprintf("%s=%s", MapKey, cookieValue))
	r.Header.Add("Cookie", fmt.Sprintf("%s=%s", MapKey2, cookieValue))
	_, err := client.Do(r)
	if err != nil {
		panic(nil)
	}

	// w, err := client.Do(r)
	// if err != nil {
	// 	panic(nil)
	// }
	fmt.Println(httptrace.RequestMeta(r))
	fmt.Println(JSONify(r.Jar))
	// trace.EchoRequestMeta(w, r)
}
