/*
AVACloud Oauth Authentication Helper.
*/

package avaclient

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type OauthAuthenticator struct {
	clientcredentials.Config
}

func newOauthAuthenticator(clientId, clientSecret, tokenUrl string, scopes []string) OauthAuthenticator {
	return OauthAuthenticator{
		clientcredentials.Config{
			ClientID:       clientId,
			ClientSecret:   clientSecret,
			TokenURL:       tokenUrl,
			Scopes:         scopes,
			EndpointParams: nil,
			AuthStyle:      oauth2.AuthStyleAutoDetect,
		},
	}
}

// NewCustomToken  retrieves an OAuth token-source from the auth endpoint
func NewCustomToken(clientId, clientSecret, tokenUrl string, scopes []string) (oauth2.TokenSource, error) {
	oauthConfig := newOauthAuthenticator(clientId, clientSecret, tokenUrl, scopes)

	source := oauthConfig.TokenSource(context.Background())

	// request a token
	_, err := source.Token()
	if err != nil {
		return nil, err
	}

	return source, nil
}

// NewToken  retrieves an OAuth token-source from the default auth endpoint
func NewToken(clientID, clientSecret string) (oauth2.TokenSource, error) {
	token, err := NewCustomToken(
		clientID,
		clientSecret,
		"https://identity.dangl-it.com/connect/token",
		[]string{"avacloud"},
	)

	return token, err
}
