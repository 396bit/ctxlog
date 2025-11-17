package ctxlog

import (
	"context"
	"log"
	"testing"
)

func TestMultiline(t *testing.T) {

	log.SetFlags(log.Ldate | log.Ltime)
	//log.SetOutput(os.Stderr)

	c := context.Background()
	c1 := Add(c, "Prefix-#1")

	Print(c, "see no prefix")
	Print(c1, "see prefix #1")

	c2 := Add(c, "Prefix-#2")
	c12 := PersistentAdd(c1, "Prefix-#2")
	c23 := Add(c2, "Prefix-#3")

	Print(c12, "see prefix #1 + #2")
	Print(c1, "see prefix #1 + #2 too!")
	Print(c2, "see prefix #2")
	Print(c23, "see prefix #2 + #3")

}
