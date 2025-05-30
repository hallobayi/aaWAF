package errors

import (
	"fmt"

	"php-parser/pkg/position"
)

// Error parsing error
type Error struct {
	Msg string
	Pos *position.Position
}

// NewError creates and returns new Error
func NewError(msg string, p *position.Position) *Error {
	return &Error{
		Msg: msg,
		Pos: p,
	}
}

func (e *Error) String() string {
	atLine := ""
	if e.Pos != nil {
		atLine = fmt.Sprintf(" at line %d", e.Pos.StartLine)
	}

	return fmt.Sprintf("%s%s", e.Msg, atLine)
}
