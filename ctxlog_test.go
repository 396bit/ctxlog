package ctxlog

import (
	"context"
	"log"
	"testing"
)

func TestMultiline(t *testing.T) {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//log.SetOutput(os.Stderr)

	c := context.Background()

	c1 := Add(c, "Prefix-#1")
	c2 := Add(c1, "Prefix-#2")

	Print(c, "string", 3, true)
	Print(c1, "string", 3, true)
	Print(c2, "string", 3, true)

	Print(c2, "this is a\nmulti\nline\nstring")
	Printf(c2, "string = '%s' and number = '%d'", "bla", 42)

	// if err != nil {
	// 	t.Error(err)
	// }
}
