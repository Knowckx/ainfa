package web

import (
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/zellyn/kooky"
	"github.com/zellyn/kooky/chrome"
	infa "github.tools.sap/aeolia/in-fa"
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

func PrintCookies(cks []*http.Cookie) {
	infa.Printf("Domain Name val Path")
	for _, ck := range cks {
		PrintCookie(ck)
	}
}

func PrintCookie(ck *http.Cookie) {
	val := ShortStr(ck.Value)
	infa.Printf("%s %s %s %s", ck.Domain, ck.Name, val, ck.Path)
}

func ShortStr(in string) string {
	leng := 10
	out := in
	if len(out) > leng {
		out = out[0:leng] + ".."
	}
	return out
}
