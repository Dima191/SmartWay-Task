package models

import "errors"

var (
	ErrInternal                   = errors.New("internal error")
	ErrNoData                     = errors.New("no data")
	ErrAlreadyExists              = errors.New("already exists")
	ErrEmployeeAlreadyExists      = errors.New("employee already exists")
	ErrPassportAlreadyExists      = errors.New("passport already exists")
	ErrUnknownReference           = errors.New("unknown reference")
	ErrUnknownCompanyReference    = errors.New("unknown company reference")
	ErrUnknownDepartmentReference = errors.New("unknown department reference")
	ErrCompanyDepProvide          = errors.New("provide both values: company and department")
)
