package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func Contains(ss []string, s string) bool {
	for _, ts := range ss {
		if ts == s {
			return true
		}
	}

	return false
}

func Difference(s1 []string, s2 []string) []string {
	diff := make([]string, 0)
	for _, s := range s1 {
		if !Contains(s2, s) {
			diff = append(diff, s)
		}
	}
	return diff
}

func Redirect(c *gin.Context, u string, v url.Values)  {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	parsedUrl.RawQuery = v.Encode()
	c.Redirect(http.StatusFound, parsedUrl.String())
}
