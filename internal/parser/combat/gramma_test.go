package combat

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

//go:embed testdata/combat.log
var data []byte

func Test(t *testing.T) {
	// data = []byte("19:37:08.518  CMBT   | Damage             Feen|0000001437 ->             Feen|0000001437 6227.20 (h:6227.20 s:0.00) (suicide) TRUE_DAMAGE|IGNORE_DAMAGE_SCALE|IGNORE_SHIELD|SUICIDE|KICK_EXCE <FriendlyFire> \r\n")

	// data = []byte("22:04:33.627  CMBT   | Damage     PlasmaHurtWS|0000000271 ->           Greedo|0000005981 370.37 (h:0.00 s:370.37)  THERMAL  \r\n")

	var count int

	rd := bufio.NewReader(bytes.NewReader(data))
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			t.Logf("MAIN READ: %q\n", err)
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
		// 	t.Logf("MAIN LEX: %q\n", err.Error())
		// 	break
		// }

		// res, err := l.Parse(line)
		// println(err)

		// enc := json.NewEncoder(os.Stdout)
		// enc.SetIndent("", "  ")
		// enc.Encode(toks)
	}

	fmt.Println(count)
}
