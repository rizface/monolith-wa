package provider

import (
	"database/sql"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/core/port"
	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/helper"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userrepository  port.UserRepositoryInterface
	redisrepository port.RedisRepositoryInterface
	db              *sql.DB
	redis           *redis.Client
}

func InitAuthService(
	userrepository port.UserRepositoryInterface,
	redisrepository port.RedisRepositoryInterface,
	db *sql.DB,
	redis *redis.Client,
) *AuthService {
	return &AuthService{
		userrepository:  userrepository,
		redisrepository: redisrepository,
		db:              db,
		redis:           redis,
	}
}

func (u *AuthService) Register(userdomain *domain.UserRequestDomain) (*entity.User, *constant.ErrorBuilder) {
	existingUser, err := u.userrepository.FindByUsername(u.db, userdomain.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, constant.InternalServerError(err.Error())
	}
	if existingUser != nil {
		return nil, constant.USERNAME_IS_EXISTS
	}

	hashedPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(userdomain.Password), bcrypt.DefaultCost)
	userdomain.Password = string(hashedPasswordBytes)
	err = u.userrepository.Create(u.db, userdomain)
	if err != nil {
		return nil, constant.InternalServerError(err.Error())
	}
	user, _ := u.userrepository.FindByUsername(u.db, userdomain.Username)
	return user, nil
}

func (u *AuthService) Login(userdomain *domain.UserRequestDomain) (fiber.Map, *constant.ErrorBuilder) {
	existing, err := u.userrepository.FindByUsername(u.db, userdomain.Username)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, constant.USER_NOT_FOUND
	}

	passwordIsMatch := bcrypt.CompareHashAndPassword(
		[]byte(existing.Password),
		[]byte(userdomain.Password),
	) == nil

	if !passwordIsMatch {
		return nil, constant.PASSWORD_WRONG
	}

	token, err := helper.GenerateJWT(existing)
	if err != nil {
		return nil, constant.InternalServerError(err.Error())
	}
	result := fiber.Map{
		"user": existing.ConvertToResponseDomain(),
		"token": fiber.Map{
			"bearer": token,
		},
	}

	err = u.redisrepository.StoreJWT(u.redis, token.(string), result)
	if err != nil {
		return nil, constant.InternalServerError(err.Error())
	}

	return result, nil
}
