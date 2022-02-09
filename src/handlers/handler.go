package handlers

import (
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"time"
)

const SecretKey = "secret"

type Handler struct {
	DB    *sql.DB
	Cache *redis.Client
	App   *fiber.App
	ctx   context.Context
}

func NewHandler(ctx context.Context, db *sql.DB, cache *redis.Client, app *fiber.App) (*Handler, error, error) {
	//	creating a new baseHandler and assigning the baseHandler variable with the pointer
	//	also creating dbErr and cacheError which will store the values of database connections
	var (
		handler  = &Handler{}
		dbErr    error
		cacheErr error
	)

	//	assigning newly initialized handler the values from main.go
	handler.ctx = ctx

	//	assigning database error to the result of the Connect function found in the database package
	handler.DB = db
	dbErr = handler.DB.Ping()

	//	assigning cache error to the result of pinging a new redis connection
	handler.Cache = cache
	cacheErr = handler.Cache.Ping(context.Background()).Err()

	handler.App = app

	//	return the memory address of the newly created handler, as well as
	//	any possible mySQL or redis errors we received when trying to connect
	return handler, dbErr, cacheErr
}

func (handler *Handler) CreateApp(config fiber.Config) {
	handler.App = fiber.New(config)
}

func (handler *Handler) CreateRouter(cors fiber.Handler) {
	handler.App.Use(cors)
}

func (handler *Handler) Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "hi",
	})
}

//	default method for any error response returned through the api
func (handler *Handler) ErrResponse(c *fiber.Ctx, statusCode int, fiberMap fiber.Map) error {
	c.Status(statusCode)
	return c.JSON(fiberMap)
}

//	default method for any success response returned through the api
func (handler *Handler) SuccessResponse(c *fiber.Ctx, statusCode int, fiberMap interface{}) error {
	c.Status(statusCode)
	return c.JSON(fiberMap)
}

//	generate a signed jwt token string
func (handler *Handler) GenerateJWT(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
}

//	set the jwt cookie for the user using the signed jwt token string
func (handler *Handler) SetCookie(c *fiber.Ctx, jwt string, expires time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Expires:  expires,
		HTTPOnly: true,
	})
}

//	retrieve the signed jwt token string to parse for information
func (handler *Handler) GetCookie(c *fiber.Ctx) string {
	return c.Cookies("jwt")
}
