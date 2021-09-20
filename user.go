package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type ConsoleUser struct {
	username []string
	password []string
	cookieJar *cookiejar.Jar
}

func (cu *ConsoleUser)Login()  {
	print_info("Authentication to the console with credentials")

	options := cookiejar.Options{
        
    }

    jar, err := cookiejar.New(&options)
    if err != nil {
        panic_with_msg("Something went wront", err)
    }

	cu.cookieJar = jar

	client := http.Client{Jar: cu.cookieJar}
    resp, err := client.PostForm(fmt.Sprintf("http://%s:8080/login", TARGET), url.Values{
         "password": cu.password,
         "username" : cu.username,        
	})

	if err != nil {
		panic_with_msg("Unable to login somehow. Dunno why", err)
	}
	
	body, err := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(body), "forgot_password") {
		panic_with_msg("Unable to login with credentials ! Something is wrong", err)
	}

	print_good("Successfully authenticated to the administrator interface ! ")
}