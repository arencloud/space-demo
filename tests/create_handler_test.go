package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/arencloud/space-demo/controllers"
	"github.com/arencloud/space-demo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoute(t *testing.T) {
	tests := []struct {
	  description  string 
	  route        string 
	  expectedCode int    
	  testCaseNamae string
	  body models.User
	}{
	  {
		description:  "POST HTTP status 201",
		route:        "/api/users/create",
		expectedCode: 201,
		testCaseNamae: "Testing Scenario #1: create user hadler",
		body: models.User{
			Name: "fake",
			Email: "fake@fake.com",
		},
	  },
	  {
		description:  "POST HTTP status 404, when route is not exists",
		route:        "/api/users/not-found",
		expectedCode: 404,
		testCaseNamae: "Testing Scenario #2: wrong handler endpoint",
		body: models.User{
			Name: "fake",
			Email: "fake@fake.com",
		},
	  },
	  {
		description:  "POST HTTP status 409, when data duplicated",
		route:        "/api/users/create",
		expectedCode: 409,
		testCaseNamae: "Testing Scenario #3: duplicate user",
		body: models.User{
			ID: 1,
			Name: "fake",
			Email: "fake@fake.com",
		},
	  },
	}
  
	// Define Fiber app.
	app := fiber.New()
	

	u := controllers.New()
	app.Post("/api/users/create", u.CreateUserHandler)
	for _, test := range tests {
		fmt.Println(test.testCaseNamae)

		pjuser, _ := json.Marshal(&test.body)
	  req := httptest.NewRequest("POST", test.route, bytes.NewReader(pjuser))
	  resp, _ := app.Test(req, 1)
  
	  // Verify, if the status code is as expected
	  assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
  }