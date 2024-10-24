package repository

import (
	"service/rest-api/internal/core/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

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
		return c.appClientID, err
	}
	return *result.AuthenticationResult.AccessToken, nil
}

func (c *CognitoClient) Logout(accessToken string) error {
	logoutInput := &cognito.GlobalSignOutInput{
		AccessToken: aws.String(accessToken),
	}

	_, err := c.cognitoClient.GlobalSignOut(logoutInput)
	return err
}
