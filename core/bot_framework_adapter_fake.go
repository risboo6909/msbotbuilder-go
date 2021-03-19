package core

import (
	"context"

	"github.com/risboo6909/msbotbuilder-go/connector/auth"
	"github.com/risboo6909/msbotbuilder-go/schema"
)

// All the mocks and stubs for BotFrameworkAdapter goes here.

// MockTokenValidator stub for bypassing the authentication
type MockTokenValidator struct {
}

// AuthenticateRequest mock implementation for authentication
func (jv *MockTokenValidator) AuthenticateRequest(ctx context.Context, activity schema.Activity, authHeader string, credentials auth.CredentialProvider, channelService string) (auth.ClaimsIdentity, error) {
	claims := map[string]interface{}{
		"1": "1",
	}
	return auth.NewClaimIdentity(claims, true), nil
}
