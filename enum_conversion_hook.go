package main

import (
	"context"
	"fmt"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
	"xorm.io/xorm/contexts"
	"xorm.io/xorm/convert"
	"xorm.io/xorm/log"
)

type IsProtobufEnum interface {
	IsProtobufEnum() bool
}

type EnumConversionHook struct {
	Logger log.Logger
}

func (e EnumConversionHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	for i, arg := range c.Args {
		_arg, ok := e.convert(arg)
		if !ok {
			continue
		}
		if _, ok := (_arg).(convert.Conversion); ok {
			if _, ok := (arg).(protoreflect.Enum); ok {
				enum, _ := (arg).(fmt.Stringer)
				c.Args[i] = enum.String()
			}
		}
	}

	return c.Ctx, nil
}

func (e *EnumConversionHook) convert(arg any) (any, bool) {
	defer func() {
		if err := recover(); err != nil {

		}
	}()

	rv := reflect.ValueOf(arg)
	if rv.Kind() != reflect.Ptr {
		return reflect.New(rv.Type()).Interface(), true
	}
	return arg, false
}

func (e *EnumConversionHook) AfterProcess(c *contexts.ContextHook) error {
	return nil
}

//engine.AddHook(&repository.EnumConversionHook{Logger: logger})
