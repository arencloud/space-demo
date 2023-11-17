package tests


import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/arencloud/space-demo/controllers"
	"github.com/arencloud/space-demo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)
func TestGetRoute(t *testing.T) {
	testsGet := []struct {
	  description  string 
	  route        string 
	  expectedCode int    
	  testCaseNamae string
	  body models.User
	}{
	  {
		description:  "GET HTTP status 200",
		route:        "/api/users/list",
		expectedCode: 200,
		testCaseNamae: "Testing Scenario #1: get all users",
	  },
	  {
		description:  "GET HTTP status 404, when route is not exists",
		route:        "/api/users/not-found",
		expectedCode: 404,
		testCaseNamae: "Testing Scenario #2: wrong handler endpoint",
	  },
	}

	testsGetById := []struct {
		description  string 
		route        string 
		expectedCode int    
		testCaseNamae string
		body models.User
	  }{
		{
		  description:  "GET HTTP status 404, when route is not exists",
		  route:        "/api/users/not-found",
		  expectedCode: 404,
		  testCaseNamae: "Testing Scenario #2: wrong handler endpoint",
		},
		{
		  description:  "GET HTTP status 200 by ID",
		  route:        "/api/users/list/1",
		  expectedCode: 200,
		  testCaseNamae: "Testing Scenario #3: get user by ID",
		},
		{
		  description:  "GET HTTP status 404 no user found",
		  route:        "/api/users/list/1000",
		  expectedCode: 404,
		  testCaseNamae: "Testing Scenario #4: no user found with ID",
		},
	  }
  
	// Define Fiber app.
	appGet := fiber.New()
	appGetByID := fiber.New()

	u := controllers.New()
	appGet.Get("/api/users/list", u.FindUsersHandler)
	for _, test := range testsGet {
		fmt.Println(test.testCaseNamae)
	  req := httptest.NewRequest("GET", test.route, nil)
	  resp, _ := appGet.Test(req, 1)
  
	  // Verify, if the status code is as expected
	  assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

	appGetByID.Get("/api/users/list/:id", u.FindUserByIdHandler)
	for _, test := range testsGetById {
		fmt.Println(test.testCaseNamae)
	  req := httptest.NewRequest("GET", test.route, nil)
	  resp, _ := appGetByID.Test(req, 1)
  
	  // Verify, if the status code is as expected
	  assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
  }