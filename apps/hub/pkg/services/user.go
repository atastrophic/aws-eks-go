package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/atastrophic/go-ms-with-eks/pkg/models"
	"github.com/atastrophic/go-ms-with-eks/pkg/password"
)

// UserService  to manage users
type UserService struct {
	users map[string]models.User
	passg *password.PasswordGenerator
}

// NewUserService ...
func NewUserService(passg *password.PasswordGenerator) *UserService {
	return &UserService{
		users: make(map[string]models.User),
		passg: passg,
	}
}

// Signup ...
func (svc *UserService) Signup(user models.User) error {

	if _, ok := svc.users[user.Username]; !ok {
		user.ID = uuid.New()
		hash, err := svc.passg.Generate(user.Password)
		if err != nil {
			return fmt.Errorf("error generating password %v", err)
		}
		user.Password = hash
		svc.users[user.Username] = user
		return nil
	}

	return fmt.Errorf("user already exists: %s", user.Username)
}

// Login ...
func (svc *UserService) Login(user models.User) (*models.Session, error) {

	if _, ok := svc.users[user.Username]; ok {
		hash, err := svc.passg.Generate(user.Password)
		if err != nil {
			return nil, fmt.Errorf("unable to generate password %v", err)
		}
		compare, err := svc.passg.Compare(user.Password, hash)
		if err != nil {
			return nil, fmt.Errorf("unable to compare password %v", err)
		}
		if compare {
			println(fmt.Sprintf("login successful... issuing jwt %-v", user))
			return svc.issueJwt(user)
		}
	}

	return nil, fmt.Errorf("username/password incorrect: %s", user.Username)
}

func (svc *UserService) issueJwt(user models.User) (*models.Session, error) {
	_ = time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Username:       user.Username,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &models.Session{AccessToken: tokenString}, nil
}
