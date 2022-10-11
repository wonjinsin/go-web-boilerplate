package model

import (
	"pikachu/util"
	"time"
)

// Token ...
type Token struct {
	IssuedAt        int64    `json:"iat"`
	Expire          int64    `json:"exp"`
	JwtID           *string  `json:"jti,omitempty"`
	Issuer          string   `json:"iss"`
	Audience        string   `json:"aud"`
	Subject         string   `json:"subject"` // UID
	Type            string   `json:"type"`    // Type of token
	AuthorizedParty string   `json:"azp"`     // Authorized party
	Email           string   `json:"email"`
	Roles           []string `json:"scope"`
	Scope           []string `json:"roles"`
}

// MakeUserPayload ...
func (t *Token) MakeUserPayload(projectName string, domain string, user *User) *Token {
	return &Token{
		IssuedAt:        time.Now().Unix(),
		Expire:          time.Now().AddDate(0, 0, 1).Unix(),
		Issuer:          domain,
		Audience:        util.TokenAudienceAccount,
		Subject:         user.UID,
		Type:            util.TokenTypeBaerer,
		AuthorizedParty: projectName,
		Email:           user.Email,
	}
}
