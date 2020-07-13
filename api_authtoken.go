package avaclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DanglToken represents a DanglIT auth token
type DanglToken struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken string `json:"access_token"`

	// TokenType is the type of token.
	// The Type method returns either this or "Bearer", the default.
	TokenType string `json:"token_type"`

	// ExpiresIn is the optional expiration time in seconds of the access token.
	ExpiresIn int32 `json:"expires_in"`

	// Issued time
	Issued time.Time

	// Scope
	Scope string `json:"scope"`
}

// IsTokenValid checks if the token is expired
func (st *DanglToken) IsTokenValid() bool {
	if int32(time.Now().Sub(st.Issued).Seconds()) > (st.ExpiresIn - 120) { // 120 = tolerance time in seconds
		return false
	}
	return true
}

func (st *DanglToken) String() string {
	return st.TokenType + " " + st.AccessToken
}

func newTokenRequest(tokenURL, clientID, clientSecret, scope string, v url.Values) (*http.Request, error) {
	v.Set("grant_type", "client_credentials")
	if clientID != "" {
		v.Set("client_id", clientID)
	}
	if clientSecret != "" {
		v.Set("client_secret", clientSecret)
	}
	if scope != "" {
		v.Set("scope", scope)
	}

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

// NewCustomToken  retrieves a token from the auth endpoint
func NewCustomToken(clientID, clientSecret, scope, tokenURL string, v url.Values) (*DanglToken, error) {
	req, err := newTokenRequest(tokenURL, clientID, clientSecret, scope, v)
	if err != nil {
		return nil, err
	}
	token, err := parseJSONToken(req)

	return token, err
}

// NewToken  retrieves a token from the default auth endpoint
func NewToken(clientID, clientSecret string) (*DanglToken, error) {
	token, err := NewCustomToken(
		clientID,
		clientSecret,
		"avacloud",
		"https://identity.dangl-it.com/connect/token",
		url.Values{},
	)

	return token, err
}

func parseJSONToken(req *http.Request) (*DanglToken, error) {
	ctx, _ := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	r, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if code := r.StatusCode; code < 200 || code > 299 {
		return nil, fmt.Errorf("api_authtoken: invalid response code: %v", code)
	}

	var token DanglToken
	err = json.NewDecoder(r.Body).Decode(&token)
	token.Issued = time.Now()

	if err != nil {
		return nil, err
	}

	if token.AccessToken == "" {
		return nil, errors.New("api_authtoken: server response missing access_token")
	}

	return &token, nil
}
