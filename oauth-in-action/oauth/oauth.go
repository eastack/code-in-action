package oauth

import "net/url"

type (
	Client struct {
		ClientId     string
		ClientName   string
		ClientUri    string
		ClientSecret string
		RedirectUris []string
		Scope        string
		LogoUri      string
	}

	Server struct {
		AuthorizationEndpoint string
		TokenEndpoint         string
	}

	Resource struct {
		Description string
		Name        string
	}

	TokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	ApproveInfo struct {
		AuthorizationEndpointRequest url.Values
		Scope                        []string
		User                         string
	}

	AssignedAuthorizationInfo struct {
		RefreshToken string
		AccessToken  string
		ClientId     string
		Scope        string
	}
)
