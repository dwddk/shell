package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_printProperMessage(t *testing.T) {
	messages := &Messages{
		secMsg:  "tick",
		minMsg:  "tock",
		hourMsg: "bong",
	}

	type args struct {
		messages *Messages
		i        int
	}

	expectedForSeconds := messages.secMsg + "\n"
	expectedForMinutes := messages.minMsg + "\n"
	expectedForHours := messages.hourMsg + "\n"
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{"seconds 1 sec", args{messages, 1}, expectedForSeconds},
		{"seconds 2 secs", args{messages, 2}, expectedForSeconds},
		{"seconds 3 secs", args{messages, 3}, expectedForSeconds},
		{"seconds 1 hour and a second", args{messages, 3601}, expectedForSeconds},
		{"minutes 1 minute", args{messages, 60}, expectedForMinutes},
		{"minutes 2 minutes", args{messages, 120}, expectedForMinutes},
		{"minutes 4 minutes", args{messages, 240}, expectedForMinutes},

		{"hours 1 hour", args{messages, 3600}, expectedForHours},
		{"hours 2 hours", args{messages, 2 * 3600}, expectedForHours},
		{"hours 10 hours", args{messages, 10 * 3600}, expectedForHours},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			printProperMessage(tt.args.messages, tt.args.i, &buf)
			assert.Equal(t, tt.expected, buf.String())
		})
	}
}

//
//func Test_readConfig(t *testing.T) {
//	type args struct {
//		reader io.Reader
//	}
//	tests := []struct {
//		name string
//		args args
//		want Messages
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := readConfig(tt.args.reader); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("readConfig() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_tickClock(t *testing.T) {
//	type args struct {
//		messages *Messages
//		locking  *Locking
//		duration time.Duration
//	}
//	tests := []struct {
//		name       string
//		args       args
//		wantWriter string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			writer := &bytes.Buffer{}
//			tickClock(tt.args.messages, tt.args.locking, writer, tt.args.duration)
//			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
//				t.Errorf("tickClock() = %v, want %v", gotWriter, tt.wantWriter)
//			}
//		})
//	}
//}
//
//func Test_updateConf(t *testing.T) {
//	type args struct {
//		messages *Messages
//		locking  *Locking
//		filename string
//		interval time.Duration
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}