package models

import "github.com/authink/inkstone/model"

type AuthToken struct {
	model.Created
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
	AppId        uint32 `db:"app_id"`
	AccountId    uint32 `db:"account_id"`
}

func NewAuthToken(accessToken, refreshToken string, appId, accountId uint32) *AuthToken {
	return &AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AppId:        appId,
		AccountId:    accountId,
	}
}

type AuthTokenWithApp struct {
	AuthToken
	AppName string `db:"app_name"`
}
