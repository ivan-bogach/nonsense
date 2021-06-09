package nonsense

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"strings"
)

// SliceContainsString reports whether slaice sl contains string s.
func SliceContainsString(sl []string, st string) bool {
	for _, s := range sl {
		if s == st {
			return true
		}
	}
	return false
}

// SendStringToTelegram sends string s to a telegram bot with a token from telega.txt located in your root directory
func SendStringToTelegram(s string) int {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile(usr.HomeDir + "/telega.txt")
	if err != nil {
		log.Fatal("you must create a file telega.txt with https://api.telegram.org/bottoken (e. g. https://api.telegram.org/bot110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw) in the root directory")
		return 0
	}
	str := strings.TrimSuffix(string(b), "\n")
	sendStr := str + s

	resp, err := http.Get(sendStr)
	if err != nil {
		SendStringToTelegram("Again: " + s)
	}

	defer resp.Body.Close()
	return resp.StatusCode
}
