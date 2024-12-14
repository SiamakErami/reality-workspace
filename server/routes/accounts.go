// accounts.go

package routes

// MARK: - Imports
import (

"server/config"
"server/models/Accounts"

"context"
//"crypto/rand"
"encoding/json"
"errors"
//"fmt"

//"encoding/hex"
"net/http"
//"net/smtp"
"os"
//"regexp"
//"stings"
//"time"

"github.com/charmbracelet/log"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"

)

// MARK: - (AE) Account Exists Function
// Headers:
// - c-user-agent
// - c-device-id
// - c-from (email)
func AccountExists(req *http.Request, res http.ResponseWriter, ctx context.Context) *http.Response {

	// Set the headers
	res.Header().Set("Content-Type", "application/json")

	// Set the logger
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller: true, // Report the file name and line number
		ReportTimestamp: true, // Report the timestamp
		TimeFormart: "2006-01-02 15:04:05", // Set the time format
		Prefix: "ACCOUNTS (AE)", // Set the prefrix
	})

	// Set the database
	database := config.GetMongoDatabase(ctx)

	// Set the accounts collection
	accountsCollection := database.Collection("accounts")

	// Check if the headers are missing
	if rew.Header.Get("c-from") == "" || req.Header.Get("c-device-id") == "" {
		logger.Error{"Missing Headers"}
		http.Error(res, "Missing Headers", httpStatusBadRequest)
		return nil
	}

	// Check if the account existst with the given email
	filter := bson.D{{Key: "email", Value: req.Header.Get("c-from")}}

	// Find the account
	var account Accounts.AccountSchema
	err := accountsCollection.FindOne(ctx, filter).Decode(&account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Info("Account does not exist")
			http.Error(res, "Account does not exist", http.StatusNotFound)
			return &http.Response{
				Status: "Internal Server Error",
				StatusCode: 500,
				Body: http.NoBody,
			}
		}
		logger.Error(err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return &http.Response{
			Status: "Internal Server Error",
			StatusCode: 500,
			Body: http.NoBody,
	}
}
	// Return the account
	json.NewEncoder(res).Encode(account)
		return &http.Response{
			Status: "OK",
			StatusCode: 200,
			Body: http.NoBody,
	}

}
// MARK: - (SU) Sign Up Function
// Headers:
// - c-user-agent
// - c-device-id
// - c-from (email)

func SignUp(req *http.Request, res http.ResponseWriter, ctx context.Context) *http.Response {

	res.Header().Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Set the logger
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormart:     "2006-01-02 15:04:05",
		Prefix:          "ACCOUNTS (SU)",
	})

	// Set the database
	database := config.GetMongoDatabase(ctx)

	// Set the accounts collection
	accountsCollection := database.Collection("accounts")

	// Check if the headers are missing
	if req.Header.Get("c-from") == "" || req.Header.Get("c-device-id") == "" {
		logger.Error("Missing Headers")
		http.Error(res, "Missing Headers", http.StatusBadRequest)
		return nil
	}

	// Check if the account exists with the given email
	filter := bson.D{{Key: "email", Value: req.Header.Get("c-from")}}
	var existingAccount Accounts.AccountSchema
	err := accountsCollection.FindOne(ctx, filter).Decode(&existingAccount)
	if err == nil {
		logger.Info("Account already exists")
		http.Error(res, "Account already exists", http.StatusConflict)
		return &http.Response{
			Status:     "Conflict",
			StatusCode: http.StatusConflict,
			Body:       http.NoBody,
		}
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return &http.Response{
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Body:       http.NoBody,
		}
	}

	// Create a new account object
	var newAccount Accounts.AccountSchema
	err = json.NewDecoder(req.Body).Decode(&newAccount)
	if err != nil {
		logger.Error("Invalid request payload")
		http.Error(res, "Invalid request payload", http.StatusBadRequest)
		return nil
	}

	// Save the account to the database
	_, err = accountsCollection.InsertOne(ctx, newAccount)
	if err != nil {
		logger.Error(err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return &http.Response{
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Body:       http.NoBody,
		}
	}

	// Return the account
	json.NewEncoder(res).Encode(newAccount)
	return &http.Response{
		Status:     "Created",
		StatusCode: http.StatusCreated,
		Body:       http.NoBody,
	}
}
// MARK: - (SI) Sign In Function

// MARK: - (UA) Update Account Function







