package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const data = `19:44:04.074  CMBT   | Damage Megabomb_RW_BlackHole|0000000155 ->            tuman|0000000824   0.00 (h:0.00 s:0.00)  KINETIC  
`

func main() {
	// yyDebug = 2
	p := yyNewParser()
	raw, err := os.ReadFile("/Users/pavelmilko/Desktop/my/sclogparser/test/parser/combat/testdata/heal.txt")
	_ = raw
	_ = err
	// raw = []byte(data)
	l := newLexer(string(raw), time.Now())

	_ = p.Parse(l)

	fmt.Println(len(l.res))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	// enc.Encode(l.res)
}
