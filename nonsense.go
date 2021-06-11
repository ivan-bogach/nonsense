package nonsense

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/user"
	"sort"
	"strconv"
	"strings"
	"time"
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
	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		// ABSURDLY large keys, for ABSURDLY dumb devices like raspberry.
		TLSHandshakeTimeout: 60 * time.Second,
	}
	c := &http.Client{
		Transport: t,
	}

	resp, err := c.Get(sendStr)
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

// CompareStringSlices reports whether slices a and b are equal.
func CompareStringSlices(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// StringSlicesIntersection returns two slices intersection
func StringSlicesIntersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

// OnlyUnique returns only unique strings from slice (drop duplicates)
func OnlyUnique(slice []string) []string {
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	uniqSlice := make([]string, 0, len(uniqMap))

	keys := make([]string, 0, len(uniqMap))
	for k := range uniqMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	uniqSlice = append(uniqSlice, keys...)

	return uniqSlice
}

// StringSlicesUnion returns union of two slices
func StringSlicesUnion(one, two []string) []string {
	var union []string
	union = append(union, one...)
	union = append(union, two...)
	return OnlyUnique(union)
}

// StringSliceDifference returns difference of two slices
func StringSliceDifference(one, two []string) []string {
	var difference []string
	for _, str := range one {
		if SliceContainsString(two, str) {
			continue
		} else {
			difference = append(difference, str)
		}
	}
	return difference
}
