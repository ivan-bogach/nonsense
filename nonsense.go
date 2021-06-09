package nonsense

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"strconv"
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
func SendStringToTelegram(s string) (int, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile(usr.HomeDir + "/telega.txt")
	if err != nil {
		return 0, fmt.Errorf("you must create a file telega.txt with https://api.telegram.org/bottoken (e. g. https://api.telegram.org/bot110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw) in the root directory")
	}
	str := strings.TrimSuffix(string(b), "\n")
	sendStr := str + s

	resp, err := http.Get(sendStr)
	if err != nil {
		return 0, fmt.Errorf("this is an %s error: %s", "SendStringToTelegram", err)
	}

	defer resp.Body.Close()
	return resp.StatusCode, nil
}

// StringIsNumeric reports whether s contains numeric value.
func StringIsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
