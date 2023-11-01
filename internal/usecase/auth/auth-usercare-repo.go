package auth

import "JampiCrm/internal/model"

type Usecase interface {
	RegisterUser(request model.RegisterRequest) (model.User, error)
	Login(request model.LoginRequest) (model.UserSessionRespond, error)
	CheckSession(sessionData model.UserSession) (userID string, err error)
	RefreshToken(refreshToken string) (model.UserSession, error)
	LogOutUser(request model.UserSession)
}
