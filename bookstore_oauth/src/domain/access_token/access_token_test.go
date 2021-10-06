package access_token

import "testing"

func TestAccessTokenIsExpired(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("brand new access token should not be expired")
	}
}
