package errors

type CMDError struct{
	Name string
}

func (e CMDError) Error() string {
	return e.Name
}


type CustomizeError struct{
	Name string
}

func (ce CustomizeError) Error() string{
	return ce.Name
}