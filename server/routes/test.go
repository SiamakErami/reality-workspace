func AccountExists(req *http.Request, res http.ResponseWriter, ctx context.Context) *http.Response {

	res.Header().Set("Content-Type", "application/json")

	// Set the logger
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller: true, // Report the file name and line number
		ReportTimestamp: true, // Report the timestamp
		TimeFormat: "2006-01-02 15:04:05", // Set the time format
		Prefix: "ACCOUNTS (AE)", // Set the prefix
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