package adapter

import (
	"log"
	"net/http"
	"service/rest-api/internal/port/in"

	"github.com/labstack/echo/v4"

	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoInterface interface {
    SignIn(user *UserLogin) (string, error)
}

type cognitoClient struct {
    cognitoClient *cognito.CognitoIdentityProvider
    appClientID   string
}

type UserHandler struct {
	service in.UserService
}

func NewUserHandler(s in.UserService) UserHandler {
	return UserHandler{service: s}
}

func (handler *UserHandler) GetAllUsers(context echo.Context) error {
	users, err := handler.service.GetAll()
	if err != nil {
		log.Printf("Could not get all users: %v", err)

		return context.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Unable to retrieve users",
		})
	}
	return context.JSON(http.StatusOK, users)
}

func NewCognitoClient(appClientId string) CognitoInterface {
    config := &aws.Config{Region: aws.String("us-east-2")}
    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
    client := cognito.New(sess)

    return &cognitoClient{
        cognitoClient: client,
        appClientID:   appClientId,
    }
}

func (c *cognitoClient) SignIn(user *UserLogin) (string, error) {
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

