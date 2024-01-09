package models

type AuthToken struct {
	Model
	AccessToken  string
	RefreshToken string
	AppId        uint32
	AccountId    uint32
}

func NewAuthToken(accessToken, refreshToken string, appId, accountId uint32) *AuthToken {
	return &AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AppId:        appId,
		AccountId:    accountId,
	}
}
