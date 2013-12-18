package plugins

import (
	"encoding/csv"
    "io"
	"os"
    "strconv"
    . "interpreter"
    . "interpreter/values"
)

// CsvCollectionValue

type CsvCollectionValue struct {
    pg *CsvPlugin
}

func (c *CsvCollectionValue) Equal(that Value) bool {
    // FIXME: Implement (semantic) equality
    return false
}

func (c *CsvCollectionValue) NewCursor() Cursor {
    f, err := os.Open(c.pg.fname)
    if err != nil {
        panic(err)
    }

    return &CsvCursor{f: f, reader: csv.NewReader(f)}
}

// CsvCursor

type CsvCursor struct {
    f *os.File
    reader *csv.Reader
    record []string
} 

func (c *CsvCursor) IsDone() bool {
    var err error
    c.record, err = c.reader.Read()
    if err == io.EOF {
        return true
    } else if (err != nil) {
        panic(err)
    }
    return false
}

func (c *CsvCursor) Next() Value {
    atts := make([]ConcreteRecordAttribute, 0)
    for i, v := range c.record {
        atts = append(atts, ConcreteRecordAttribute{"_" + strconv.Itoa(i), NewConcreteStringValue(v)})
    }
    return NewConcreteRecordValue(atts)
}

func (c *CsvCursor) Close() {
    c.f.Close()
}

// CsvPlugin

type CsvPlugin struct {
    fname string
}

func (pg *CsvPlugin) Init() {}

func (pg *CsvPlugin) Fini() {}

func (pg *CsvPlugin) GetCollection() CollectionValue {
    return &CsvCollectionValue{pg}
}

func NewCsvPlugin(fname string) *CsvPlugin {
    return &CsvPlugin{fname: fname}
}