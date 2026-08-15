package errs

import "errors"

var (
	ErrPgNoRows              = errors.New("postgres no rows")
	ErrPgUniqueViolation     = errors.New("postgres unique violation")
	ErrPgForeignKeyViolation = errors.New("postgres foreign key violation")
	ErrPgNotNullViolation    = errors.New("postgres not null violation")
	ErrPgCheckViolation      = errors.New("postgres check violation")
)
