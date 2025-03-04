package sample

import "github.com/google/wire"

func InitalizeService() *SimpleService {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil
}
