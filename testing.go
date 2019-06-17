package scopeagent

import (
	"github.com/opentracing/opentracing-go"
	"runtime"
	"strings"
	"testing"
)

func InstrumentTest(t *testing.T, f func(t *testing.T)) {
	pc, _, _, _ := runtime.Caller(1)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	span := opentracing.StartSpan(t.Name(), opentracing.Tags{
		"span.kind":  "test",
		"test.name":  funcName,
		"test.suite": packageName,
	})
	defer span.Finish()
	defer func() {
		if t.Failed() {
			span.SetTag("test.status", "FAIL")
		} else if t.Skipped() {
			span.SetTag("test.status", "SKIP")
		} else {
			span.SetTag("test.status", "PASS")
		}
	}()
	f(t)
}