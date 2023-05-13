package web

import (
	"net/http"
	"os"

	infa "github.com/Knowckx/infa"
	"github.com/Knowckx/infa/util"
	"github.com/rs/zerolog/log"
	"github.com/zellyn/kooky"
	"github.com/zellyn/kooky/browser/chrome"
)

func GetHostCookies(host string) []*http.Cookie {
	koo := GetChromeCookies()
	outs := []*http.Cookie{}
	for _, ko := range koo {
		if ko.Domain == host {
			ht := ko.Cookie
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
	infa.Printf("Domain Name value Path -- len %d", len(cks))
	for _, ck := range cks {
		PrintCookie(ck)
	}
}

func PrintCookie(ck *http.Cookie) {
	val := util.ShortStr(ck.Value)
	infa.Printf("%s %s %s %s", ck.Domain, ck.Name, val, ck.Path)
}
