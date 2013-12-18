package plugins

import (
	"encoding/json"
	"os"
	. "interpreter"
    . "interpreter/values"
)

// JsonCollectionValue

type JsonCollectionValue struct {
	pg *JsonPlugin
	path []interface{}
}

func (v *JsonCollectionValue) Equal(that Value) bool {
    // FIXME: Implement (semantic) equality
    return false
}

func (v *JsonCollectionValue) NewCursor() Cursor {
	node := v.pg.data
	for _, p := range v.path {
		switch p := p.(type) {
		case string: node = node.(map[string]interface{})[p]
		case int: node = node.([]interface{})[p]
		default: panic("invalid json path")
		}
	}
	
	vs, ok := node.([]interface{})
	if !ok {
		panic("invalid collection")
	}

	return &JsonCursor{v, vs, -1}
}

// JsonCursor

type JsonCursor struct {
	v *JsonCollectionValue
    vs []interface{}
    pos int
} 

func (c *JsonCursor) IsDone() bool {
	c.pos += 1
	return c.pos == len(c.vs)
}

func (c *JsonCursor) newValue(any interface{}, path []interface{}) Value {
	switch v := any.(type) {
	case bool: return NewConcreteBoolValue(v)
	case float64: return NewConcreteFloatValue(v)
	case int: return NewConcreteIntValue(v)
	case string: return NewConcreteStringValue(v)
	case []interface{}:
		return &JsonCollectionValue{c.v.pg, path}
	case map[string]interface{}: 
	    atts := make([]ConcreteRecordAttribute, len(v))
	    i := 0
		for k, nv := range v {
	       	atts[i] = ConcreteRecordAttribute{k, c.newValue(nv, append(path, k))}
    		i += 1
    	}
    	return NewConcreteRecordValue(atts)
	default: panic("unsupported type")
	}
}

func (c *JsonCursor) Next() Value {
	return c.newValue(c.vs[c.pos], append(c.v.path, c.pos))
}

func (c *JsonCursor) Close() {}

// JsonPlugin

type JsonPlugin struct {
    fname string
    data interface{}
}

func (pg *JsonPlugin) Init() {
    f, err := os.Open(pg.fname)
    defer f.Close()
    if err != nil {
        panic(err)
    }

    // The JSON package reads the entire file at once
	dec := json.NewDecoder(f)
    if err := dec.Decode(&pg.data); err != nil {
    	panic(err)
    }
}

func (pg *JsonPlugin) Fini() {}

func (pg *JsonPlugin) GetCollection() CollectionValue {
    return &JsonCollectionValue{pg: pg}
}

func NewJsonPlugin(fname string) *JsonPlugin {
    return &JsonPlugin{fname: fname}
}