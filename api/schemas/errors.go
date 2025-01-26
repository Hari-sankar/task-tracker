package schemas

import "task-tracker/api/utils"

//Defining custom errors
var (
	ErrUserNotFound = utils.NewErrorStruct(404, "user not found")

	ErrUsernameExists = utils.NewErrorStruct(409, "username already exists")

	ErrInvalidCredentials = utils.NewErrorStruct(401, "invalid credentials")

	ErrUnauthorized = utils.NewErrorStruct(401, "unauthorized access")

	ErrInvalidInput = utils.NewErrorStruct(400, "invalid input data")

	ErrInternalServer = utils.NewErrorStruct(500, "internal server error")

	ErrDatabaseConnection = utils.NewErrorStruct(503, "database connection error")

	ErrInvalidToken = utils.NewErrorStruct(401, "invalid or expired token")
)
