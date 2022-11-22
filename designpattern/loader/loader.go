package loader

import (
	"context"
	"errors"
	"fmt"
)

const  (
	loader1 = "loader1"
	loader2 = "loader2"
	loader3 = "loader3"	
)

var (
	errNotExisted = errors.New("loader not existed")
)


type loaderFunc func(ctx context.Context, params string) (interface{}, error)

var ruleLoader =  map[string]interface{} {
	loader1: loaderFunc1(),
	loader2: loaderFunc2(),
	loader3: loaderFunc3(),
}

func loaderFunc1() loaderFunc {
	return func(ctx context.Context, params string) (interface{}, error) {
		fmt.Println(ctx, params)
		return nil, nil
	}
}

func loaderFunc2() loaderFunc {
	return func(ctx context.Context, params string) (interface{}, error) {
		fmt.Println(ctx, params)
		return nil, nil
	}
}

func loaderFunc3() loaderFunc {
	return func(ctx context.Context, params string) (interface{}, error) {
		fmt.Println(ctx, params)
		return nil, nil
	}
}

// usage

func usage(ctx context.Context, key, params string) (interface{}, error){
	ruleLoader, ok := ruleLoader[key] 
	if !ok {
		return nil, errNotExisted
	}
	switch loader := ruleLoader.(type) {
	case loaderFunc:
		return loader(ctx, params)
	default:
		return nil, errNotExisted
	}	
}