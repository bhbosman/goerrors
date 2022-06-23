package goerrors

import (
	"errors"
	"fmt"
)

var NotImplemented error = errors.New("not implemented")

var InvalidParam error = errors.New("invalid param")
var TimeOut error = errors.New("invalid param")
var InvalidType error = errors.New("not implemented")
var InvalidSignature error = errors.New("InvalidSignature")
var InvalidState error = errors.New("InvalidState")
var InvalidSequenceNumber error = errors.New("InvalidSequenceNumber")
var DuplicateKey error = errors.New("DuplicateKey")

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
