package factory

type IFactory interface {
	method() (any, error)
	attribute() (any, error)
	asset() (any, error)
}

func GetFactory(str string) IFactory {
	if str == "factory1" {
		return &Factory1{}
	}
	return &Factory2{}
}

type Factory1 struct {
			
}

func (f *Factory1) method() (any, error) {
	return nil, nil
}

func (f *Factory1) attribute() (any, error) {
	return nil, nil
}

func (f *Factory1)asset() (any, error) {
	return nil, nil
}

type Factory2 struct {
			
}

func (f *Factory2) method() (any, error) {
	return nil, nil
}

func (f *Factory2) attribute() (any, error) {
	return nil, nil
}

func (f *Factory2)asset() (any, error) {
	return nil, nil
}