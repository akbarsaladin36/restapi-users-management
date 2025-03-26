package services

import (
	"restapi-users-management/helpers"
	"restapi-users-management/inputs"
	"restapi-users-management/models"
	"restapi-users-management/repositories"
	"strings"
	"time"
)

type AuthService interface {
	FindOneService(user_username string) (models.User, error)
	RegisterService(registerInput inputs.RegisterInput) (models.User, error)
	LoginService(loginInput inputs.LoginInput) (models.User, error)
	FindSessionService(user_uuid string) (models.Session, error)
	CreateSessionService(loginUser models.User, tokenString string) (models.Session, error)
	UpdateSessionService(user_uuid string, tokenString string, session models.Session) (models.Session, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(authRepository repositories.AuthRepository) *authService {
	return &authService{authRepository}
}

func (as *authService) FindOneService(user_username string) (models.User, error) {
	user, err := as.authRepository.FindOne(user_username)

	return user, err
}

func (as *authService) RegisterService(registerInput inputs.RegisterInput) (models.User, error) {
	hashedPassword, _ := helpers.HashPassword(registerInput.Password)

	userUUID := strings.ReplaceAll(helpers.GenerateUUID(registerInput.Username), "-", "")

	user := models.User{
		UserUuid:                userUUID,
		UserUsername:            registerInput.Username,
		UserEmail:               registerInput.Email,
		UserPassword:            hashedPassword,
		UserStatusCd:            "active",
		UserRole:                "user",
		UserCreatedDate:         time.Now(),
		UserCreatedUserUuid:     userUUID,
		UserCreatedUserUsername: registerInput.Username,
	}

	newRegisterUser, err := as.authRepository.Create(user)

	return newRegisterUser, err
}

func (as *authService) LoginService(loginInput inputs.LoginInput) (models.User, error) {
	checkUser, err := as.authRepository.FindOne(loginInput.Username)

	helpers.CheckPassword(checkUser.UserPassword, loginInput.Password)

	return checkUser, err
}

func (as *authService) FindSessionService(user_uuid string) (models.Session, error) {
	checkSession, err := as.authRepository.FindSession(user_uuid)

	return checkSession, err
}

func (as *authService) CreateSessionService(loginUser models.User, tokenString string) (models.Session, error) {
	sessionId := helpers.GenerateSessionToken(32)

	session := models.Session{
		SessionId:                  sessionId,
		SessionToken:               tokenString,
		SessionUserUuid:            loginUser.UserUuid,
		SessionUserUsername:        loginUser.UserUsername,
		SessionUserRole:            loginUser.UserRole,
		SessionStartAt:             time.Now(),
		SessionExpiredAt:           time.Now().Add(24 * time.Hour),
		SessionStatusCd:            "active",
		SessionCreatedDate:         time.Now(),
		SessionCreatedUserUuid:     loginUser.UserUuid,
		SessionCreatedUserUsername: loginUser.UserUsername,
	}

	newSession, err := as.authRepository.CreateSession(session)

	return newSession, err
}

func (as *authService) UpdateSessionService(user_uuid string, tokenString string, session models.Session) (models.Session, error) {
	checkSession, _ := as.authRepository.FindSession(user_uuid)

	checkSession.SessionToken = tokenString
	checkSession.SessionStartAt = time.Now()
	checkSession.SessionExpiredAt = time.Now().Add(24 * time.Hour)
	checkSession.SessionStatusCd = "active"
	checkSession.SessionUpdatedDate = time.Now()
	checkSession.SessionUpdatedUserUuid = checkSession.SessionUserUuid
	checkSession.SessionUpdatedUserUsername = checkSession.SessionUserUsername

	updateSession, err := as.authRepository.UpdateSession(user_uuid, checkSession)

	return updateSession, err
}
