package repository

import (
	"os"
	"service/rest-api/internal/core/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoRepository struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientID   string
	userPoolID    string
}

func NewCognitoRepository() *CognitoRepository {
	config := &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))}
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	return &CognitoRepository{
		cognitoClient: cognito.New(sess),
		appClientID:   os.Getenv("COGNITO_APP_CLIENT_ID"),
		userPoolID:    os.Getenv("COGNITO_USER_POOL_ID"),
	}
}

func (c *CognitoRepository) SignUp(user *domain.User) error {
	input := &cognito.SignUpInput{
		ClientId: aws.String(c.appClientID),
		Username: aws.String(user.Email),
		Password: aws.String(user.Password),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("name"),
				Value: aws.String(user.Name),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	}

	_, err := c.cognitoClient.SignUp(input)
	return err
}

func (c *CognitoRepository) ConfirmAccount(confirmation *domain.UserConfirmation) error {
	input := &cognito.ConfirmSignUpInput{
		ClientId:         aws.String(c.appClientID),
		Username:         aws.String(confirmation.Email),
		ConfirmationCode: aws.String(confirmation.Code),
	}

	_, err := c.cognitoClient.ConfirmSignUp(input)
	return err
}
