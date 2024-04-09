package errs

import (
	"database/sql"
	"strings"
)

type Entity struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error"`
	InternalCode string `json:"internal_code"`
	Details      string `json:"message"`
	Trace        string `json:"trace"`
}

func (e *Entity) Error() string {
	return e.ErrorMessage
}

func New(
	code int,
	err string,
	msg string,
	InternalCode string,
	trace string,
) *Entity {
	return &Entity{
		Code:         code,
		ErrorMessage: err,
		Details:      msg,
		Trace:        trace,
		InternalCode: InternalCode,
	}
}

func NewInternal(
	Message string,
	IntCode string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         500,
		ErrorMessage: "Internal server error",
		Details:      Message,
		Trace:        Trace,
		InternalCode: IntCode,
	}
}

func NewBadRequest(
	Message string,
	IntCode string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         400,
		ErrorMessage: "Bad request",
		Details:      Message,
		Trace:        Trace,
		InternalCode: IntCode,
	}
}

func NewUnauthorized(
	Message string,
	IntCode string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         401,
		ErrorMessage: "Unauthorized",
		Details:      Message,
		Trace:        Trace,
		InternalCode: IntCode,
	}
}

func NewForbidden(
	Message string,
	IntCode string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         403,
		ErrorMessage: "Forbidden",
		Details:      Message,
		Trace:        Trace,
		InternalCode: IntCode,
	}
}

func NewNotFound(
	Message string,
	IntCode string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         404,
		ErrorMessage: "Not found",
		Details:      Message,
		InternalCode: IntCode,
		Trace:        Trace,
	}
}

func NewConflict(
	Message string,
	IntCode string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         409,
		ErrorMessage: "Conflict",
		Details:      Message,
		InternalCode: IntCode,
		Trace:        Trace,
	}
}

func NewNotAcceptable(
	Message string,
	IntCode string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         406,
		ErrorMessage: "Not acceptable",
		Details:      Message,
		InternalCode: IntCode,
		Trace:        Trace,
	}
}

func NewTooEarly(
	Message string,
	InternalCode string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         425,
		ErrorMessage: "Too early",
		Details:      Message,
		Trace:        Trace,
		InternalCode: InternalCode,
	}
}

func SqlError(err error, code string, trace string) *Entity {
	if err == sql.ErrNoRows {
		return NewNotFound(
			"no result found",
			code,
			trace,
		)
	}

	errString := strings.ToLower(err.Error())

	if strings.Contains(errString, "conflict") {
		return NewConflict(
			"duplicate entry",
			code,
			trace,
		)
	}

	return NewInternal(
		err.Error(),
		code,
		trace,
	)
}
