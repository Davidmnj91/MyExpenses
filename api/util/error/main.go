package error


type Interface interface {
	New() *Error
}

type Error struct {
	HTTP HttpErrorInterface
}

func New() *Error  {
	http := &HTTP{}
	return &Error{HTTP: http}
}
