package operators

import . "interpreter"

func Scan(pg Plugin, cout chan []Value) {
    c := pg.GetCollection().NewCursor()
    for !c.IsDone() {
        out := make([]Value, 1)
        out[0] = c.Next()
        cout <- out
    }
    c.Close()
    close(cout)
}