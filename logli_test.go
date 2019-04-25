package logli

import (
	"bytes"
	"log"
	"os"
	"testing"
)

const msg = "Welcome, logli"

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestColorize(t *testing.T) {
	got := Colorize(msg, Color.LightBlue)
	want := vary(Color.LightBlue, TypeFace.Normal, msg)
	if got != want {
		t.Errorf(`Colorize(%s, Color.LightBlue) = %s; want %s`, msg, got, want)
	}
	return
}

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	Debug(msg)

	t.Log(buf.String())
}

func BenchmarkDebug(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Debug(msg)
	}
}

func ExamplePrintF() {
	timeDisabled = true
	PrintF("CustomPrintNT - %s", msg)
	// Output: CustomPrintNT - Welcome, logli
}
