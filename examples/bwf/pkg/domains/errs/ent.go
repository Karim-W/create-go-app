package errs

type Entity struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error"`
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
	trace string,
) *Entity {
	return &Entity{
		Code:         code,
		ErrorMessage: err,
		Details:      msg,
		Trace:        trace,
	}
}

func NewInternal(
	Message string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         500,
		ErrorMessage: "Internal server error",
		Details:      Message,
		Trace:        Trace,
	}
}

func NewBadRequest(
	Message string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         400,
		ErrorMessage: "Bad request",
		Details:      Message,
		Trace:        Trace,
	}
}

func NewUnauthorized(
	Message string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         401,
		ErrorMessage: "Unauthorized",
		Details:      Message,
		Trace:        Trace,
	}
}

func NewForbidden(
	Message string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         403,
		ErrorMessage: "Forbidden",
		Details:      Message,
		Trace:        Trace,
	}
}

func NewNotFound(
	Message string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         404,
		ErrorMessage: "Not found",
		Details:      Message,
		Trace:        Trace,
	}
}

func NewConflict(
	Message string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         409,
		ErrorMessage: "Conflict",
		Details:      Message,
		Trace:        Trace,
	}
}

func NewNotAcceptable(
	Message string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         406,
		ErrorMessage: "Not acceptable",
		Details:      Message,
		Trace:        Trace,
	}
}

func NewTooEarly(
	Message string,
	Trace string,
) *Entity {
	return &Entity{
		Code:         425,
		ErrorMessage: "Too early",
		Details:      Message,
		Trace:        Trace,
	}
}
