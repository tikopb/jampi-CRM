// container.go
package usecase

import (
	"time"

	userRepo "JampiCrm/internal/repository/user"
	authUsecase "JampiCrm/internal/usecase/auth"

	midUtil "JampiCrm/internal/delivery/auth"

	"crypto/rand"
	"crypto/rsa"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Container struct {
	AuthUsecase authUsecase.Usecase

	Middleware midUtil.MidlewareInterface
}

func NewContainer(db *gorm.DB) *Container {

	secret := GetEnv("key_secret")
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	userRepo, err := userRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Second, 48*time.Hour)
	if err != nil {
		panic("errorr repo")
	}
	authUsecase := authUsecase.GetUsecase(userRepo)

	middleware := midUtil.GetAuthMiddleware(authUsecase)

	return &Container{
		AuthUsecase: authUsecase,
		Middleware:  middleware,
	}
}

func GetEnv(param string) string {
	var value string

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic("config environment not found!")
	}

	value = viper.GetString("key_secret")
	return value
}
