package error

type ApplicationErrorInterface interface {
	Message() string
	Error() error
	InvalidId(err error) *AppError
	NotFound(err error) *AppError
	ValidationError(err error) *AppError
	AlreadyExists(err error) *AppError
	RepositoryError(err error) *AppError
	NotAuthenticated(err error) *AppError
	NotAuthorized(err error) *AppError
	UnknownError(err error) *AppError
}

// AppError defines an application (domain) error
type AppError struct {
	message string
	error   error
}

// Message return http error message
func (appError *AppError) Message() string {
	return appError.message
}

// Error return http error message
func (appError *AppError) Error() error {
	return appError.error
}

// InvalidId return invalidId error message
func (appError *AppError) InvalidId(err error) *AppError {
	return &AppError{
		message: "Invalid ID value, must be ID complaint",
		error:   err,
	}
}

// NotFound return not found error message
func (appError *AppError) NotFound(err error) *AppError {
	return &AppError{
		message: "Record not found",
		error:   err,
	}
}

// ValidationError return validation error message
func (appError *AppError) ValidationError(err error) *AppError {
	return &AppError{
		message: "Validation error",
		error:   err,
	}
}

// AlreadyExists return already exists error message
func (appError *AppError) AlreadyExists(err error) *AppError {
	return &AppError{
		message: "Resource already exists",
		error:   err,
	}
}

// RepositoryError return repository error message
func (appError *AppError) RepositoryError(err error) *AppError {
	return &AppError{
		message: "Error in repository operation",
		error:   err,
	}
}

// NotAuthenticated return not authenticated error message
func (appError *AppError) NotAuthenticated(err error) *AppError {
	return &AppError{
		message: "Not Authenticated",
		error:   err,
	}
}

// NotAuthorized return not authorized error message
func (appError *AppError) NotAuthorized(err error) *AppError {
	return &AppError{
		message: "Not Authorized",
		error:   err,
	}
}

// UnknownError return unknown error message
func (appError *AppError) UnknownError(err error) *AppError {
	return &AppError{
		message: "Something went wrong",
		error:   err,
	}
}