package game

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"testing"
	"time"
)

//go:embed testdata/game.log
var data []byte

func Test(t *testing.T) {
	// data = []byte("01:15:07.375         | client: player 58 leave game\r\n")

	// data = []byte("01:14:58.327         | client: ADD_PLAYER 0 (XxKing68xX, 700176) status 4 team 1\r\n")

	// data = []byte("00:48:19.042         | client: ADD_PLAYER 0 (Bondagebear [], 3929961) status 4 team 1 group 5389724\r\n")

	var count int

	rd := bufio.NewReader(bytes.NewReader(data))
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("MAIN READ: %q\n", err)
			break
		}

		p := NewParser()

		res, err := p.Parse(time.Now(), line)
		if err != nil {
			var gErr *GrammaError
			if errors.As(err, &gErr) {
				t.Log("GRAMMA ERROR: ", gErr)
				t.FailNow()
			}
			if !errors.Is(err, ErrLineIsNotFinished) {
				t.Error(err)
				t.FailNow()
			}
		}

		if res != nil {
			count++
			// t.Logf("MAIN LINE: %q\n", line)
			// t.Logf("LOG LINE: %+#v\n", res.Line)
		}

		// _ = p.Parse(l)
		// if err := l.Err(); err != nil {
		// 	fmt.Printf("MAIN LEX: %q\n", err.Error())
		// 	break
		// }

		// res, err := l.Parse(line)
		// println(err)

		// enc := json.NewEncoder(os.Stdout)
		// enc.SetIndent("", "  ")
		// enc.Encode(toks)
	}
}
