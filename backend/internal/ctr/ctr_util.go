package ctr

import (
	"fmt"
	"reflect"
	"strings"
)

func Reflect(controllers ...CTR) ([]*Ctr, error) {
	ctrs := make([]*Ctr, len(controllers))
	var err error
	for i, c := range controllers {
		if ctrs[i], err = reflectController(c); err != nil {
			return nil, err
		}
	}
	return ctrs, nil
}

func reflectController(controller any) (*Ctr, error) {
	var (
		ctr *Ctr
	)

	if initer, ok := controller.(CTR); ok {
		ctr = initer.Init()
	} else {
		return nil, fmt.Errorf("controller:%v is missing init method", reflect.TypeOf(controller).Elem())
	}

	controllerType := reflect.TypeOf(controller)
	controllerValue := reflect.ValueOf(controller)

	for i := controllerType.NumMethod() - 1; i >= 0; i-- {
		method := controllerType.Method(i)

		if method.Type.NumOut() == 1 && method.Type.Out(0).AssignableTo(reflect.TypeOf((*Route)(nil))) {
			routeFn, ok := controllerValue.MethodByName(method.Name).Interface().(func() *Route)
			if !ok {
				continue
			}
			route := routeFn()
			route.ctr = ctr
			if route.name == "" {
				route.name = strings.Replace(fmt.Sprintf("%v.%s", controllerType, method.Name), "*", "", 1)
			}
			ctr.Routes = append(ctr.Routes, route)
		}
	}

	return ctr, nil
}
