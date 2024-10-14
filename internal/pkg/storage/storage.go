package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Variable struct {
	type_of_val Kind
	str         string
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
	switch kind := GetType(value); kind {
	case KindInt:
		r.inner[key] = Variable{str: value, type_of_val: kind}
	case KindStr:
		r.inner[key] = Variable{str: value, type_of_val: kind}
	case KindUD:
		r.logger.Error("unknown value type", zap.String("key", key), zap.Any("value", value))

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
		r.logger.Error("unknown value type", zap.String("key", key))
		return "sth went wrong"
	}
	r.logger.Info("returned type of value", zap.String("type", string(v_strct.type_of_val)))
	r.logger.Sync()
	return string(v_strct.type_of_val)

}

type Kind string

const (
	KindInt Kind = "D"
	KindStr Kind = "S"
	KindUD  Kind = "UN"
)

func GetType(data string) Kind {
	var val any
	val, err := strconv.Atoi(data)

	if err != nil {
		val = data
	}
	switch val.(type) {
	case int:
		return KindInt
	case string:
		return KindStr
	default:
		return KindUD
	}
}
