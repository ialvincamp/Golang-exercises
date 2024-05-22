package tests

import (
	"exercise4/internal/database"
	"exercise4/internal/models"
	"exercise4/internal/server"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func TestAdd(t *testing.T) {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Create a Fiber app for testing

	app := fiber.New()
	database.ConnectToDB()
	database.DB.AutoMigrate(&models.User{}, &models.Claims{})
	// Inject the Fiber app into the server
	s := &server.FiberServer{App: app}
	// Define a route in the Fiber app
	app.Post("/login", s.LoginUser)

	// Create a test HTTP request
	payload := `{"Username":"alvin01","Password":"testpassword"}`
	req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
	if err != nil {
		t.Fatalf("error creating request. Err: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}

	// Check the HTTP status code
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status Created; got %v", resp.Status)
	}

	// Check the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if payload != string(body) {
		t.Errorf("expected response body to be %v; got %v", payload, string(body))
	}
}

// func TestSubtract(t *testing.T) {
// 	result := Subtract(5, 3)
// 	if result != 2 {
// 		t.Errorf("Subtract(5, 3) = %d; want 2", result)
// 	}
// }
// func TestHandlers(t *testing.T) {
// 	// Create a Fiber app for testing
// 	app := fiber.New()
// 	// Inject the Fiber app into the server
// 	s := &server.FiberServer{App: app}
// 	// Define a route in the Fiber app
// 	app.Get("/", s.HelloWorldHandler)
// 	// Create a test HTTP request
// 	req, err := http.NewRequest("GET", "/", nil)
// 	if err != nil {
// 		t.Fatalf("error creating request. Err: %v", err)
// 	}
// 	// Perform the request
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Fatalf("error making request to server. Err: %v", err)
// 	}
// 	// Your test assertions...
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK; got %v", resp.Status)
// 	}
// 	expected := "{\"message\":\"Hello World\"}"
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("error reading response body. Err: %v", err)
// 	}
// 	if expected != string(body) {
// 		t.Errorf("expected response body to be %v; got %v", expected, string(body))
// 	}
// }

// func Login(t *testing.T) {
// 	// Create a Fiber app for testing
// 	app := fiber.New()
// 	// Inject the Fiber app into the server
// 	s := &server.FiberServer{App: app}
// 	// Define a route in the Fiber app
// 	app.Get("/", s.HelloWorldHandler)
// 	// Create a test HTTP request
// 	req, err := http.NewRequest("GET", "/", nil)
// 	if err != nil {
// 		t.Fatalf("error creating request. Err: %v", err)
// 	}
// 	// Perform the request
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Fatalf("error making request to server. Err: %v", err)
// 	}
// 	// Your test assertions...
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK; got %v", resp.Status)
// 	}
// 	expected := "{\"message\":\"Hello World\"}"
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("error reading response body. Err: %v", err)
// 	}
// 	if expected != string(body) {
// 		t.Errorf("expected response body to be %v; got %v", expected, string(body))
// 	}
// }

// func register(t *testing.T) {

// 	data := []byte(`{
// 		"Email": "alvintest6@gmail.com",
// 		"Username": "user6",
// 		"Password": "mypassword"

// 	}`)

// 	// Create a Fiber app for testing
// 	app := fiber.New()
// 	// Inject the Fiber app into the server
// 	s := &server.FiberServer{App: app}
// 	// Define a route in the Fiber app
// 	app.Post("/register", s.CreateUser)
// 	// Create a test HTTP request
// 	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(data))
// 	if err != nil {
// 		t.Fatalf("error creating request. Err: %v", err)
// 	}

// 	// Perform the request
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Fatalf("error making request to server. Err: %v", err)
// 	}
// 	// Your test assertions...
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK; got %v", resp.Status)
// 	}
// 	expected := "{\"Success\":200}"
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("error reading response body. Err: %v", data)
// 	}
// 	if expected != string(body) {
// 		t.Errorf("expected response body to be %v; got %v", expected, string(body))
// 	}
//}
