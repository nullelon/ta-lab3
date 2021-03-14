package ta_lab3

type List interface {
	Add(el interface{})
	Insert(index int, el interface{})
	Get(index int) interface{}
	Find(el interface{}) (index int, ok bool)
	Remove(index int)
	Length() int
}
