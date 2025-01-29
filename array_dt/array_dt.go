package array_dt

type Array_DT interface {
	Get() interface{}
	Add(a interface{})
}

type Value struct {
	value interface{}
}

func (v *Value) Add(a interface{}) {
	v.value = a
}

func (v Value) Get() interface{} {
	return v.value
}
