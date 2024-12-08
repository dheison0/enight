package tokens

import (
	"server/extra"
	"time"
)

const TOKEN_LIFETIME = 2 * time.Hour

type TokenUser struct {
	User      string `json:"user"`
	ExpiresAt int64  `json:"expires_at"`
}

var tokenRelation = map[string]TokenUser{}

func init() {
	go (func() { // used to clean unused tokens
		for {
			for user, token := range tokenRelation {
				if token.ExpiresAt < time.Now().Unix() {
					delete(tokenRelation, user)
				}
			}
			time.Sleep(5 * time.Second)
		}
	})()
}

func Create(user string) string {
	token := extra.RandomString(10)
	tokenRelation[token] = TokenUser{
		User:      user,
		ExpiresAt: time.Now().Add(TOKEN_LIFETIME).Unix(),
	}
	return token
}

func GetUser(token string) string {
	if user, ok := tokenRelation[token]; ok {
		return user.User
	} else {
		return ""
	}
}

func Delete(token string) {
	delete(tokenRelation, token)
}
