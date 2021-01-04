package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"github.com/rwrr/oauth2-in-action/oauth"
	"github.com/rwrr/oauth2-in-action/util"
	"github.com/rwrr/randomstring"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var client = oauth.Client{
	ClientId:     "oauth-client-1",
	ClientName:   "oauth-client-one",
	ClientUri:    "http://localhost:9000",
	ClientSecret: "oauth-client-secret-1",
	LogoUri:      "https://dss0.bdstatic.com/6Ox1bjeh1BF3odCf/it/u=279126261,584360358&fm=218&app=92&f=JPEG?w=121&h=75&s=5402177445305C321052A0D80200C0FB",
	RedirectUris: []string{"http://localhost:9000/callback"},
	Scope:        "foo",
}

var server = oauth.Server{
	AuthorizationEndpoint: "http://localhost:9001/authorize",
	TokenEndpoint:         "http://localhost:9001/token",
}

var accessToken string
var refreshToken string
var scope string
var state string
var protectedResourceApi = "http://localhost:9002/resource"

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"access_token": accessToken,
		"Scope":        scope,
	})
}

func Authorize(c *gin.Context) {
	accessToken = ""
	state = randomstring.Generate(5)

	query := make(url.Values)
	query.Set("response_type", "code")
	query.Set("client_id", client.ClientId)
	query.Set("redirect_uri", client.RedirectUris[0])
	query.Set("state", state)
	query.Set("scope", client.Scope)

	log.Println("redirect", server.AuthorizationEndpoint)
	util.Redirect(c, server.AuthorizationEndpoint, query)
}

func Callback(c *gin.Context) {
	if error := c.Query("error"); error != "" {
		// it's an error response, act accordingly
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": error,
		})
		return
	}

	if c.Request.URL.Query().Get("state") != state {
		log.Printf("State DOES NOT MATCH: expected %s got %s\n", state, c.Request.URL.Query().Get("state"))
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "State value did not match",
		})
		return
	}

	formData := make(url.Values)
	formData.Set("grant_type", "authorization_code")
	//formData.Set("grant_type", "refresh_token")
	formData.Set("code", c.Request.URL.Query().Get("code"))
	formData.Set("redirect_uri", client.RedirectUris[0])

	_, body, errors := gorequest.New().
		Post(server.TokenEndpoint).
		SetBasicAuth(client.ClientId, client.ClientSecret).
		Type(gorequest.TypeForm).
		Send(formData).
		End()

	log.Println(body)
	if errors != nil {
		print(errors)
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": errors,
		})
		return
	}

	responseToken := oauth.TokenResponse{}
	json.NewDecoder(strings.NewReader(body)).Decode(&responseToken)
	accessToken = responseToken.AccessToken

	c.HTML(http.StatusOK, "index.html", gin.H{
		"access_token": accessToken,
		"Scope":        scope,
	})
}

func FetchResource(c *gin.Context) {
	if accessToken == "" {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Missing Access Token",
		})
	}
	log.Printf("Making request with access token %s\n", accessToken)

	response, body, errors := gorequest.New().
		Post(protectedResourceApi).
		AppendHeader("Authorization", "Bearer "+accessToken).
		End()

	if response.StatusCode >= 200 && response.StatusCode <= 300 {
		resource := oauth.Resource{}
		json.NewDecoder(strings.NewReader(body)).Decode(&resource)

		c.HTML(http.StatusOK, "data.html", gin.H{
			"description": resource.Description,
			"name":        resource.Name,
		})
	} else {
		if refreshToken != "" {
			refreshAccessToken()
			return
		} else {
			c.HTML(http.StatusOK, "error.html", gin.H{
				"error": response.StatusCode,
			})
			return
		}
	}

	if errors != nil {
		print(errors)
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": errors,
		})
		return
	}
}

func refreshAccessToken() {
	formData := make(url.Values)
	formData.Set("grant_type", "refresh_token")
	formData.Set("refresh_token", refreshToken)

	response, body, _ := gorequest.New().
		Post(server.TokenEndpoint).
		SetBasicAuth(client.ClientId, client.ClientSecret).
		Type(gorequest.TypeForm).
		Send(formData).
		End()

	if response.StatusCode >= 200 && response.StatusCode <= 300 {
		print(response.Body)
		print(body)

		//c.HTML(http.StatusOK, "data.html", gin.H{
		//	"description": resource.Description,
		//	"name":        resource.Name,
		//})
	} else {
		//if refreshToken != "" {
		//	refreshAccessToken()
		//	return
		//} else {
		//	c.HTML(http.StatusOK, "error.html", gin.H{
		//		"error": response.StatusCode,
		//	})
		//	return
		//}
	}
}
