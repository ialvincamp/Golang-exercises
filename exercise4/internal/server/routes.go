package server

import (
	db "exercise4/internal/database"
	"exercise4/internal/models"
	"exercise4/internal/util"
	"math/rand"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type CacheItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var t = cache.New(5*time.Minute, 10*time.Minute)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
	s.App.Post("/register", s.CreateUser)
	s.App.Post("/login", s.LoginUser)
	s.App.Post("/cache/set", s.setCache)
	s.App.Get("/cache/get/:key", s.getCache)
	s.App.Delete("/cache/delete/:key", s.deleteCache)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) setCache(c *fiber.Ctx) error {

	item := new(CacheItem)

	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	t.Set(item.Key, item.Value, 1*time.Minute) // Set value with expiration of 1 minute
	return c.SendString("Value set in cache")
}

func (s *FiberServer) getCache(c *fiber.Ctx) error {
	key := c.Params("key")
	value, found := t.Get(key)
	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Value not found in cache",
		})
	}
	return c.JSON(fiber.Map{
		"value": value,
	})
}

func (s *FiberServer) deleteCache(c *fiber.Ctx) error {
	key := c.Params("key")
	t.Delete(key)
	return c.SendString("Value deleted from cache")
}

func (s *FiberServer) CreateUser(c *fiber.Ctx) error {
	u := new(models.User)

	if err := c.BodyParser(u); err != nil {
		return c.JSON(fiber.Map{
			"error": true,
			"input": "Please review your input",
		})
	}

	// Hashing the password with a random salt
	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(
		password,
		rand.Intn(bcrypt.MaxCost-bcrypt.MinCost)+bcrypt.MinCost,
	)

	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)

	if err := db.DB.Create(&u).Error; err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"general": "Something went wrong, please try again later.",
		})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := util.GenerateTokens(u.UUID.String())
	accessCookie, refreshCookie := util.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"status":        200,
	})
}

func (s *FiberServer) LoginUser(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}

	u := new(models.User)
	if res := db.DB.Where(
		&models.User{Email: input.Identity}).Or(
		&models.User{Username: input.Identity},
	).First(&u); res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "general": "Invalid Credentials."})
	}

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return c.JSON(fiber.Map{"error": true, "general": "Invalid Credentials."})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := util.GenerateTokens(u.UUID.String())
	accessCookie, refreshCookie := util.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"status":        200,
	})
}
