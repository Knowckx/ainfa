package web

import (
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/zellyn/kooky"
	"github.com/zellyn/kooky/chrome"
)

func GetHostCookies(host string) []*http.Cookie {
	koo := GetChromeCookies()
	outs := []*http.Cookie{}
	for _, ko := range koo {
		if ko.Domain == host {
			ht := ko.HTTPCookie()
			outs = append(outs, &ht)
		}
	}
	return outs
}

func GetChromeCookies() []*kooky.Cookie {
	dir, _ := os.UserConfigDir() // "/<USER>/Library/Application Support/"
	cookiesFile := dir + "/Google/Chrome/Default/Cookies"
	cookies, err := chrome.ReadCookies(cookiesFile)
	if err != nil {
		log.Error().Stack().Err(err).Send()
	}
	return cookies
}
