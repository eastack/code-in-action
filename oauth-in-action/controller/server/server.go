package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rwrr/oauth2-in-action/jsondb"
	"github.com/rwrr/oauth2-in-action/oauth"
	"github.com/rwrr/oauth2-in-action/util"
	"github.com/rwrr/randomstring"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func init() {
	log.SetOutput(os.Stdout)
}

var approveInfos = make(map[string]oauth.ApproveInfo)

var clients = []oauth.Client{
	{
		ClientId:     "oauth-client-1",
		ClientName:   "oauth-client-one",
		ClientUri:    "http://localhost:9000",
		ClientSecret: "oauth-client-secret-1",
		LogoUri:      "https://dss0.bdstatic.com/6Ox1bjeh1BF3odCf/it/u=279126261,584360358&fm=218&app=92&f=JPEG?w=121&h=75&s=5402177445305C321052A0D80200C0FB",
		RedirectUris: []string{"http://localhost:9000/callback"},
		Scope:        "foo bar",
	},
}

var server = oauth.Server{
	AuthorizationEndpoint: "http://localhost:9001/authorize",
	TokenEndpoint:         "http://localhost:9001/token",
}

var authorizeQuery = make(map[string]url.Values)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"oauthClients": clients,
		"oauthServer":  server,
	})
}

func Authorize(c *gin.Context) {
	clientId := c.Query("client_id")
	client := getClient(clientId)

	if client == nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Unknown client " + clientId,
		})
	} else if !util.Contains(client.RedirectUris, c.Query("redirect_uri")) {
		log.Printf("Mismatched redirect URI, expected %s got %s",
			client.RedirectUris, c.Query("redirect_uri"))
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Invalid redirect URI",
		})
	} else {
		queryScope := strings.Split(c.Query("scope"), " ")
		dbScope := strings.Split(client.Scope, " ")

		if len(util.Difference(queryScope, dbScope)) > 0 {
			redirectUri := c.Query("redirect_uri")
			query := make(url.Values)
			query.Set("error", "invalid_scope")

			util.Redirect(c, redirectUri, query)
			return
		}

		requestId := randomstring.Generate(8)
		query := c.Request.URL.Query()
		authorizeQuery[requestId] = query

		c.HTML(http.StatusOK, "approve.html", gin.H{
			"client":    client,
			"requestId": requestId,
			"scopes":    queryScope,
		})
	}
}

func Approve(c *gin.Context) {
	requestId := c.PostForm("requestId")
	request := authorizeQuery[requestId]
	delete(authorizeQuery, requestId)

	if request == nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "No matching authorization request",
		})
	}

	decision := c.PostForm("approve")

	if decision == "approve" {
		if request.Get("response_type") == "code" {
			code := randomstring.Generate(8)
			postForm := c.Request.PostForm
			requestScope := make([]string, 0)
			for k, _ := range postForm {
				if strings.HasPrefix(k, "scope_") {
					requestScope = append(requestScope, strings.TrimPrefix(k, "scope_"))
				}
			}

			client := getClient(request.Get("client_id"))
			cachedScope := strings.Split(client.Scope, " ")

			if len(util.Difference(requestScope, cachedScope)) > 0 {
				redirectUri := request.Get("redirect_uri")
				query := make(url.Values)
				query.Set("error", "invalid_scope")

				util.Redirect(c, redirectUri, query)
				return
			}

			// 存储用户信息，请求信息，和 scope 信息
			approveInfos[code] = oauth.ApproveInfo{
				User:                         "",
				Scope:                        requestScope,
				AuthorizationEndpointRequest: request,
			}

			// 通过浏览器重定向将 code 传给客户端
			redirectUri := request.Get("redirect_uri")
			query := make(url.Values)
			query.Set("code", code)
			query.Set("state", request.Get("state"))

			util.Redirect(c, redirectUri, query)
			return
		} else {
			redirectUrl := request.Get("redirect_uri")
			query := make(url.Values)
			query.Set("error", "unsupported_response_type")

			util.Redirect(c, redirectUrl, query)
			return
		}
	} else {
		query := make(url.Values)
		query.Set("error", "access_denied")
		redirectUrl := request.Get("redirect_uri")

		util.Redirect(c, redirectUrl, query)
		return
	}
}

// Token 用于根据客户端的请求向其颁发access_token
func Token(c *gin.Context) {
	// 尝试从Header中获取客户端认证标记
	clientId, clientSecret, ok := c.Request.BasicAuth()
	if ok {
		// 如果已经在header中获得到了认证信息，而提交的表单中再次包含认证信息是错误的
		if c.PostForm("client_id") != "" {
			log.Println("Client attempted to authenticate with multiple methods")
			c.HTML(http.StatusUnauthorized, "error.html", gin.H{
				"error": "invalid_client",
			})
			return
		}
	} else {
		clientId = c.PostForm("client_id")
		clientSecret = c.PostForm("client_secret")
	}

	client := getClient(clientId)

	// 如果客户端不存在
	if client == nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Unknown client " + clientId,
		})
		return
	}

	// 校验客户端密钥
	if client.ClientSecret != clientSecret {
		log.Printf("Mismatched client secret, expected %s got %s\n", client.ClientSecret, clientSecret)
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"error": "invalid_client",
		})
		return
	}

	// 客户端身份认证通过，进入token颁发流程
	if c.PostForm("grant_type") == "authorization_code" {
		approveInfo, exists := approveInfos[c.PostForm("code")]
		if exists {
			delete(approveInfos, c.PostForm("code"))

			if approveInfo.AuthorizationEndpointRequest.Get("client_id") == clientId {
				auth := oauth.AssignedAuthorizationInfo{
					AccessToken: randomstring.Generate(32),
					RefreshToken: randomstring.Generate(32),
					Scope:       strings.Join(approveInfo.Scope, " "),
					ClientId:    clientId,
				}

				jsondb.Write("by-access-token", auth.AccessToken, auth)
				jsondb.Write("by-refresh-token", auth.RefreshToken, auth)

				log.Printf("Issuing access token %s\n", auth.AccessToken)
				log.Printf("with scope %s\n", auth)

				c.JSON(http.StatusOK, gin.H{
					"access_token": auth.AccessToken,
					"refresh_token": auth.RefreshToken,
					"token_type": "Bearer",
					"scope": auth.Scope,
				})
			}
		}

	} else if c.PostForm("grant_type") == "refresh_token" {
		log.Println("todo refresh token")
	}
}

func getClient(clientId string) *oauth.Client {
	for i, client := range clients {
		if client.ClientId == clientId {
			return &clients[i]
		}
	}
	return nil
}
