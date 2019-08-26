package httputil

import (
	"context"
	"fmt"

	"github.com/lovego/tracer"
)

func ExampleTrace() {
	runRequest()
	println()
	runRequest()
	// Output:
}

func runRequest() {
	ctx := tracer.Start(context.Background(), "")
	_, err := GetCtx(ctx, "http", "https://united.hztl3.com/", nil, nil)
	if err != nil {
		println(err.Error())
	}

	t := tracer.Get(ctx).Children[0]
	println(fmt.Sprintf("total: %dms", int(t.Duration)))
	for _, log := range t.Logs {
		println(log)
	}
}
