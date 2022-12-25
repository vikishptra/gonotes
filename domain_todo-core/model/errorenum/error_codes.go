package errorenum

import "vikishptra/shared/model/apperror"

const (
	SomethingError apperror.ErrorType = "ER0000 something error"
	MessageEmpty   apperror.ErrorType = "ER0001 message empty"
	DataNull       apperror.ErrorType = "ER0002 data not found"
	ObjSame        apperror.ErrorType = "ER0002 data duplicated"
)
