package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Variable struct {
	type_of_val string
	str         string
	integer     int
}

type Storage struct {
	inner  map[string]Variable
	logger *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, cde := zap.NewProduction()
	if cde != nil {
		return Storage{}, cde
	}
	defer logger.Sync()
	logger.Info("Created new Storage object")
	return Storage{
		inner:  make(map[string]Variable),
		logger: logger,
	}, nil
}

func (r Storage) Set(key, value string) {
	defer r.logger.Sync()
	if digit, err := strconv.Atoi(value); err == nil {
		r.inner[key] = Variable{integer: digit, type_of_val: "D"}
	} else {
		r.inner[key] = Variable{str: value, type_of_val: "S"}
	}
	r.logger.Info("Added value on key:", zap.String("key", key), zap.Any("data", value))
	r.logger.Sync()
}

func (r Storage) Get_Var(key string) (Variable, bool) {
	res, err := r.inner[key]
	if !err {
		return Variable{}, false
	}
	return res, true
}

func (r Storage) Get(key string) *string {
	res, err := r.Get_Var(key)
	if !err {
		return nil
	}
	r.logger.Info("returned value by key", zap.String("key", key))
	r.logger.Sync()
	return &res.str
}

func (r Storage) GetKind(key string) string {
	v_strct, cde := r.inner[key]
	if !cde {
		return "sth went wrong"
	}
	r.logger.Info("returned type of value")
	r.logger.Sync()
	return v_strct.type_of_val

}
