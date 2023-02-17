package service

import "errors"

var ErrInvalidToken = errors.New("invalid token")

type AuthService struct {
	secret string
}

func NewAuthService(secret string) *AuthService {
	return &AuthService{secret: secret}
}

func (s *AuthService) ValidateToken(token string) error {
	if token == "" {
		return ErrInvalidToken
	}
	// TODO: implement JWT validation against s.secret
	return nil
}
