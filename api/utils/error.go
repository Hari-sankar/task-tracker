package utils

import (
	"database/sql"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
)

type ErrorStruct struct {
	Code    int
	Message string
}

func NewErrorStruct(code int, message string) ErrorStruct {
	return ErrorStruct{
		Code:    code,
		Message: message,
	}
}

func (err ErrorStruct) Error() string {
	return err.Message
}

// HandleDBError handles database errors and returns a custom error struct
func HandleDBError(err error) ErrorStruct {
	if err == nil {
		return ErrorStruct{} // No error, return an empty struct
	}

	// Handle "no rows" error
	if errors.Is(err, sql.ErrNoRows) {
		return NewErrorStruct(404, "resource not found")
	}

	// Check if the error is a Postgres-specific error
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		log.Printf("Postgres error: Code=%s, Message=%s", pgErr.Code, pgErr.Message)
		switch pgErr.Code {
		case "23505": // unique violation
			return NewErrorStruct(409, "resource already exists")
		case "23503": // foreign key violation
			return NewErrorStruct(400, "invalid reference")

		case "23502": // Not null violation
			return NewErrorStruct(400, "missing required field")

		case "42P01": // undefined table
			return NewErrorStruct(500, "internal server error - table not found")
		case "28P01": // invalid password
			return NewErrorStruct(401, "invalid credentials")
		default:
			return NewErrorStruct(503, "database connection error")
		}
	}

	// Generic fallback for other errors
	return NewErrorStruct(500, "internal server error")
}
