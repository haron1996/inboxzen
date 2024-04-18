package utils

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func ConstructGoogleConfig() (*oauth2.Config, error) {
	scopes := []string{
		gmail.GmailReadonlyScope,
		gmail.GmailComposeScope,
		gmail.GmailModifyScope,
		"https://www.googleapis.com/auth/gmail.settings.basic",
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	}

	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		return nil, err
	}

	return config, nil
}
