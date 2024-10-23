package adapter

import (
	"net/http"

	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	"github.com/labstack/echo/v4"
)

type CognitoInterface interface {
	SignIn(user domain.UserLogin) (string, error)
}

type AuthHandler struct {
	service in.AuthService
}

func NewAuthHandler(s in.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	var user domain.UserLogin

	// Bind the JSON request body to the UserLogin struct
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Call the service to authenticate the user
	token, err := h.service.SignIn(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid login credentials"})
	}

	// Return the token as JSON response
	return c.JSON(http.StatusOK, map[string]string{"access_token": token})
}

type CognitoClient struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientID   string
}

func NewCognitoClient(appClientID string) *CognitoClient {
	config := &aws.Config{Region: aws.String("us-east-2")}
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	client := cognito.New(sess)

	return &CognitoClient{
		cognitoClient: client,
		appClientID:   appClientID,
	}
}

func (c *CognitoClient) SignIn(user domain.UserLogin) (string, error) {
	authInput := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": user.Email,
			"PASSWORD": user.Password,
		}),
		ClientId: aws.String(c.appClientID),
	}

	result, err := c.cognitoClient.InitiateAuth(authInput)
	if err != nil {
		return "", err
	}
	return *result.AuthenticationResult.AccessToken, nil
}
