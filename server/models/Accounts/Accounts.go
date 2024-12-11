package Accounts;

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MARK: - Account Struct
type AccountSchema struct {
	ID          primitive.ObjectID          `json:"_id" bson:"_id"`                   // unique object id
	AccountID   *string                      `json:"account_id" bson:"account_id"`    // unique account id
	FirstName   *string                      `json:"first_name" bson:"first_name"`    // user first name
	LastName    *string                      `json:"last_name" bson:"last_name"`      // user last name
	Birthday    *string                     `json:"birthday" bson:"birthday"`         // user birthday
	Email       *string                      `json:"email" bson:"email"`              // user email address
	Phone       *string                      `json:"phone" bson:"phone"`              // user phone number
	CreatedAt   *time.Time                   `json:"created_at" bson:"created_at"`    // account creation date

}