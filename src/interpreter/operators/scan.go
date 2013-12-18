package operators

import . "interpreter"

func Scan(pg Plugin, cout chan []Value) {
    c := pg.GetCollection().NewCursor()
    for !c.IsDone() {
        cout <- WrapValue(c.Next())
    }
    c.Close()
    close(cout)
}