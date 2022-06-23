package goerrors

import "fmt"

type InvalidParamError struct {
	ParameterName string
	Message       string
}

func NewInvalidParamError(parameterName string, message string) *InvalidParamError {
	return &InvalidParamError{ParameterName: parameterName, Message: message}
}

func NewInvalidNilParamError(parameterName string) *InvalidParamError {
	return &InvalidParamError{
		ParameterName: parameterName,
		Message:       fmt.Sprintf("%v can not be nil", parameterName),
	}
}

func (self *InvalidParamError) Error() string {
	return fmt.Sprintf("Param: %v, Message: %v", self.ParameterName, self.Message)
}
