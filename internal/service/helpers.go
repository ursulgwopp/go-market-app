package service

import (
	"errors"
	"regexp"

	"github.com/ursulgwopp/go-market-app/internal/models"
)

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// // Check for at least one uppercase letter
	// if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
	// 	return errors.New("password must contain at least one uppercase letter")
	// }

	// // Check for at least one lowercase letter
	// if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
	// 	return errors.New("password must contain at least one lowercase letter")
	// }

	// Check for at least one digit
	if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
		return errors.New("password must contain at least one digit")
	}

	// // Check for at least one special character
	// if matched, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, password); !matched {
	// 	return errors.New("password must contain at least one special character")
	// }

	return nil
}

func validateUsername(username string) error {
	// Check length
	if len(username) < 3 || len(username) > 30 {
		return errors.New("username must be between 3 and 20 characters long")
	}

	// Check for allowed characters (alphanumeric and underscores)
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username); !matched {
		return errors.New("username can only contain alphanumeric characters and underscores")
	}

	return nil
}

func validateEmail(email string) error {
	// Regular expression for validating an Email
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	// Check if the email matches the regex
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func validateProduct(product models.ProductRequest) error {
	if len(product.Name) < 1 || len(product.Name) > 50 {
		return errors.New("product name must be between 1 and 50 characters long")
	}

	if len(product.Description) < 1 || len(product.Description) > 150 {
		return errors.New("product description must be between 1 and 150 characters long")
	}

	if product.Quantity < 1 || product.Quantity > 9999 {
		return errors.New("invalid product quantity")
	}

	if product.Price < 1 {
		return errors.New("invalid price")
	}

	return nil
}
