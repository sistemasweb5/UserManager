package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"strconv"
	"strings"
	"poc/go-aws-cognito/clientCognito"	
	"os"
)


type UserResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	CustomID      string `json:"custom_id"`
	EmailVerified bool   `json:"email_verified"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	clientCognito := clientCognito.NewCognitoClient(os.Getenv("COGNITO_APP_CLIENT_ID"))
	r := gin.Default()
	r.POST("user", func(context *gin.Context) {
		err := CreateUser(context, clientCognito)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusCreated, gin.H{"message": "user created"})
	})
	r.POST("user/confirmation", func(context *gin.Context) {
		err := ConfirmAccount(context, clientCognito)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusCreated, gin.H{"message": "user confirmed"})
	})
	r.POST("user/login", func(context *gin.Context) {
		token, err := SignIn(context, clientCognito)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusCreated, gin.H{"token": token})
	})
	r.GET("user", func(context *gin.Context) {
		user, err := GetUserByToken(context, clientCognito)
		if err != nil {
			if err.Error() == "token not found" {
				context.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
				return
			}
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"user": user})
	})
	r.PATCH("user/password", func(context *gin.Context) {
		err := UpdatePassword(context, clientCognito)
		if err != nil {
			if err.Error() == "token not found" {
				context.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
				return
			}
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "password updated"})
	})
	fmt.Println("Server is running on port 8181")
	err = r.Run(":8181")
	if err != nil {
		panic(err)
	}
}

func CreateUser(c *gin.Context, cognito clientCognito.CognitoInterface) error {
	var user clientCognito.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return errors.New("invalid json")
	}
	err := cognito.SignUp(&user)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func ConfirmAccount(c *gin.Context, cognito clientCognito.CognitoInterface) error {
	var user clientCognito.UserConfirmation
	if err := c.ShouldBindJSON(&user); err != nil {
		return errors.New("invalid json")
	}
	err := cognito.ConfirmAccount(&user)
	if err != nil {
		return fmt.Errorf("could not confirm user: %w", err)
	}
	return nil
}

func SignIn(c *gin.Context, cognito clientCognito.CognitoInterface) (string, error) {
	var user clientCognito.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		return "", errors.New("invalid json")
	}
	token, err := cognito.SignIn(&user)
	if err != nil {
		return "", fmt.Errorf("could not sign in: %w", err)
	}
	return token, nil
}

func GetUserByToken(c *gin.Context, cognito clientCognito.CognitoInterface) (*UserResponse, error) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		return nil, errors.New("token not found")
	}
	cognitoUser, err := cognito.GetUserByToken(token)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", err)
	}
	user := &UserResponse{}
	for _, attribute := range cognitoUser.UserAttributes {
		switch *attribute.Name {
		case "sub":
			user.ID = *attribute.Value
		case "name":
			user.Name = *attribute.Value
		case "email":
			user.Email = *attribute.Value
		case "custom:custom_id":
			user.CustomID = *attribute.Value
		case "email_verified":
			emailVerified, err := strconv.ParseBool(*attribute.Value)
			if err == nil {
				user.EmailVerified = emailVerified
			}
		}
	}
	return user, nil
}

func UpdatePassword(c *gin.Context, cognito clientCognito.CognitoInterface) error {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		return errors.New("token not found")
	}
	var user clientCognito.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		return errors.New("invalid json")
	}
	err := cognito.UpdatePassword(&user)
	if err != nil {
		return errors.New("could not update password")
	}
	return nil
}
