package httputil

import (
	"context"
	"crypto/tls"
	"net/http/httptrace"
	"time"

	"github.com/lovego/tracer"
)

func httpTrace(ctx context.Context) (context.Context, *time.Time) {
	var getConn = time.Now() // set it because GetConn is not called sometimes
	var start, gotFirstResponseByte time.Time

	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			getConn = time.Now()
		},

		DNSStart: func(info httptrace.DNSStartInfo) {
			start = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			logTimeSpent(ctx, "DNS", start)
		},

		ConnectStart: func(network, addr string) {
			start = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			logTimeSpent(ctx, "Connect", start)
		},

		TLSHandshakeStart: func() {
			start = time.Now()
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			logTimeSpent(ctx, "TLS", start)
		},

		GotConn: func(info httptrace.GotConnInfo) {
			logTimeSpent(ctx, "GotConn", getConn)
			start = time.Now()
		},

		WroteRequest: func(info httptrace.WroteRequestInfo) {
			logTimeSpent(ctx, "Write", start)
			start = time.Now()
		},

		GotFirstResponseByte: func() {
			logTimeSpent(ctx, "Wait", start)
			gotFirstResponseByte = time.Now()
		},
	}
	return httptrace.WithClientTrace(ctx, trace), &gotFirstResponseByte
}

func logTimeSpent(ctx context.Context, name string, start time.Time) {
	if start.IsZero() {
		return
	}
	spent := time.Since(start).Round(time.Millisecond)
	if spent == 0 {
		tracer.Log(ctx, name+": ", "0ms")
	} else {
		tracer.Log(ctx, name+": ", spent.String())
	}
}
