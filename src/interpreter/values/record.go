package values

import . "interpreter"

// RecordValue

type RecordValue interface {
	Size() int
	GetName(i int) string
	GetValue(i int) Value
	GetValueByName(name string) Value
}

// ConcreteRecordValue

type ConcreteRecordAttribute struct {
	N string
	V Value	
}

type ConcreteRecordValue struct {
	atts []ConcreteRecordAttribute
}

func (this *ConcreteRecordValue) Equal(that Value) bool {
	switch that := that.(type) {
	case RecordValue:
		if len(this.atts) != that.Size() {
			return false
		}

		for i, att := range this.atts {
			if att.N != that.GetName(i) {
				return false
			}
			if !att.V.Equal(that.GetValue(i)) {
				return false
			}
		}
		return true
	}
	return false
}

func (r *ConcreteRecordValue) Size() int {
	return len(r.atts)
}

func (r *ConcreteRecordValue) GetName(i int) string {
	return r.atts[i].N
}

func (r *ConcreteRecordValue) GetValue(i int) Value {
	return r.atts[i].V
}

func (r *ConcreteRecordValue) GetValueByName(name string) Value {
	for _, att := range r.atts {
		if att.N == name {
			return att.V
		}
	}
	panic("unexpected attribute name")
}

func NewConcreteRecordValue(atts []ConcreteRecordAttribute) *ConcreteRecordValue {
	return &ConcreteRecordValue{atts}
}