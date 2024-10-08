package storage

import (
	"strconv"
)

type Variable struct {
	type_of_val string
	str         string
	integer     int
}

type Storage struct {
	inner map[string]Variable
}

func NewStorage() (Storage, error) {
	return Storage{
		inner: make(map[string]Variable),
	}, nil
}

func (r Storage) Set(key, value string) {
	if digit, err := strconv.Atoi(value); err == nil {
		r.inner[key] = Variable{integer: digit, type_of_val: "D"}
	} else {
		r.inner[key] = Variable{str: value, type_of_val: "S"}
	}

}

func (r Storage) Get(key string) interface{} {
	v_strct, cde := r.inner[key]

	if !cde {
		return nil
	}
	if v_strct.type_of_val == "D" {
		return v_strct.integer
	} else if v_strct.type_of_val == "S" {
		return v_strct.str
	}
	return nil
}

func (r Storage) GetKind(key string) string {
	v_strct, cde := r.inner[key]
	if !cde {
		return "sth went wrong"
	}
	return v_strct.type_of_val

}
