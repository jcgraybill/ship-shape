package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	InfoLogger   *log.Logger
	UpdateLogger *log.Logger
	DrawLogger   *log.Logger
	mem          runtime.MemStats
)

func init() {
	tm, err := os.OpenFile("tm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err == nil {
		UpdateLogger = log.New(tm, "", log.Ldate|log.Ltime|log.Lshortfile)
		DrawLogger = log.New(tm, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		log.Fatal(err)
	}

	summary, err := os.OpenFile("summary.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		InfoLogger = log.New(summary, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		log.Fatal(err)
	}
}

func (g *Game) instrument() {
	if g.count%300 == 0 {
		runtime.ReadMemStats(&mem)
		InfoLogger.Printf("tick %d summary tps %.0f fps %.0f alloc %d sys %d gc %d", g.count, ebiten.CurrentTPS(), ebiten.CurrentFPS(), mem.Alloc/1048576, mem.Sys/1048576, mem.NumGC)
	}
}

func (g *Game) measure(f func()) (uint64, uint64) {
	runtime.ReadMemStats(&mem)
	t0 := time.Now()
	m0 := mem.HeapAlloc
	f()
	t1 := time.Now()
	runtime.ReadMemStats(&mem)
	m1 := mem.HeapAlloc
	return uint64(t1.Sub(t0).Microseconds()), m1 - m0
}

func (g *Game) SplitToString(a []uint64, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(v, 10)
	}
	return strings.Join(b, sep)
}
