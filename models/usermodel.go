package models

import (
	"errors"
	"strings"
	"regexp"
	"golang.org/x/crypto/bcrypt"
)

// Address represents the address information
type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}

// User represents the user model
type User struct {
	ID            int     `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	PhoneNo       string  `json:"phone_no"`
	Password      string  `json:"-"`
	IsEmailVerify bool    `json:"is_email_verify"`
	IsPhoneVerify bool    `json:"is_phone_verify"`
	Address       Address `json:"address"`
}

// NewUser creates a new User instance with validation
func NewUser(firstName, lastName, email, phoneNo, password string, address Address) (*User, error) {
	// Validate FirstName
	if firstName == "" {
		return nil, errors.New("Please enter the FirstName")
	}

	// Validate FirstName length
	if len(firstName) < 3 || len(firstName) > 15 {
		return nil, errors.New("FirstName length should be between 3 and 15 characters")
	}

	// Validate LastName length
	if len(lastName) > 0 && (len(lastName) < 3 || len(lastName) > 15) {
		return nil, errors.New("LastName length should be between 3 and 15 characters")
	}

	// Combine FirstName and LastName to form Name
	name := strings.TrimSpace(firstName + " " + lastName)

	// Validate Email
	if email == "" {
		return nil, errors.New("Please enter the Email")
	}

	// Add additional email validation if needed

	// Validate PhoneNo
	if phoneNo != "" && len(phoneNo) != 10 {
		return nil, errors.New("PhoneNo should be a 10-digit number")
	}

	// Validate and hash Password
	if password == "" {
		return nil, errors.New("Please enter the Password")
	}
	passwordRegex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`
	if matched, _ := regexp.MatchString(passwordRegex, password); !matched {
		return nil,errors.New("Password should have at least one uppercase, one lowercase, one digit, and be at least 8 characters long")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:     firstName,
		LastName:      lastName,
		Name:          name,
		Email:         email,
		PhoneNo:       phoneNo,
		Password:      string(hashedPassword),
		IsEmailVerify: false,
		IsPhoneVerify: false,
		Address:       address,
	}, nil
}
 
