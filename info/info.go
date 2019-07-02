package info

import (
	"fmt"
	"net"
	"time"

	"github.com/u6du/zerolog"
	"github.com/u6du/zerolog/log"
)

func Msg(msg string) {
	log.Info().Out(msg)
}

func Msgf(format string, v ...interface{}) {
	log.Info().Out(fmt.Sprintf(format, v...))
}

func Err(err error) *zerolog.Event {
	return log.Info().Err(err)
}
func Dict(key string, dict *zerolog.Event) *zerolog.Event { return log.Info().Dict(key, dict) }
func Array(key string, arr zerolog.LogArrayMarshaler) *zerolog.Event {
	return log.Info().Array(key, arr)
}
func EmbedObject(obj zerolog.LogObjectMarshaler) *zerolog.Event {
	return log.Info().EmbedObject(obj)
}
func Object(key string, obj zerolog.LogObjectMarshaler) *zerolog.Event {
	return log.Info().Object(key, obj)
}
func Str(key string, val string) *zerolog.Event         { return log.Info().Str(key, val) }
func Strs(key string, vals []string) *zerolog.Event     { return log.Info().Strs(key, vals) }
func Bytes(key string, val []byte) *zerolog.Event       { return log.Info().Bytes(key, val) }
func Hex(key string, val []byte) *zerolog.Event         { return log.Info().Hex(key, val) }
func RawJSON(key string, b []byte) *zerolog.Event       { return log.Info().RawJSON(key, b) }
func AnErr(key string, err error) *zerolog.Event        { return log.Info().AnErr(key, err) }
func Errs(key string, errs []error) *zerolog.Event      { return log.Info().Errs(key, errs) }
func Bool(key string, b bool) *zerolog.Event            { return log.Info().Bool(key, b) }
func Bools(key string, b []bool) *zerolog.Event         { return log.Info().Bools(key, b) }
func Int(key string, i int) *zerolog.Event              { return log.Info().Int(key, i) }
func Ints(key string, i []int) *zerolog.Event           { return log.Info().Ints(key, i) }
func Int8(key string, i int8) *zerolog.Event            { return log.Info().Int8(key, i) }
func Ints8(key string, i []int8) *zerolog.Event         { return log.Info().Ints8(key, i) }
func Int16(key string, i int16) *zerolog.Event          { return log.Info().Int16(key, i) }
func Ints16(key string, i []int16) *zerolog.Event       { return log.Info().Ints16(key, i) }
func Int32(key string, i int32) *zerolog.Event          { return log.Info().Int32(key, i) }
func Ints32(key string, i []int32) *zerolog.Event       { return log.Info().Ints32(key, i) }
func Int64(key string, i int64) *zerolog.Event          { return log.Info().Int64(key, i) }
func Ints64(key string, i []int64) *zerolog.Event       { return log.Info().Ints64(key, i) }
func Uint(key string, i uint) *zerolog.Event            { return log.Info().Uint(key, i) }
func Uints(key string, i []uint) *zerolog.Event         { return log.Info().Uints(key, i) }
func Uint8(key string, i uint8) *zerolog.Event          { return log.Info().Uint8(key, i) }
func Uints8(key string, i []uint8) *zerolog.Event       { return log.Info().Uints8(key, i) }
func Uint16(key string, i uint16) *zerolog.Event        { return log.Info().Uint16(key, i) }
func Uints16(key string, i []uint16) *zerolog.Event     { return log.Info().Uints16(key, i) }
func Uint32(key string, i uint32) *zerolog.Event        { return log.Info().Uint32(key, i) }
func Uints32(key string, i []uint32) *zerolog.Event     { return log.Info().Uints32(key, i) }
func Uint64(key string, i uint64) *zerolog.Event        { return log.Info().Uint64(key, i) }
func Uints64(key string, i []uint64) *zerolog.Event     { return log.Info().Uints64(key, i) }
func Float32(key string, f float32) *zerolog.Event      { return log.Info().Float32(key, f) }
func Floats32(key string, f []float32) *zerolog.Event   { return log.Info().Floats32(key, f) }
func Float64(key string, f float64) *zerolog.Event      { return log.Info().Float64(key, f) }
func Floats64(key string, f []float64) *zerolog.Event   { return log.Info().Floats64(key, f) }
func Time(key string, t time.Time) *zerolog.Event       { return log.Info().Time(key, t) }
func Times(key string, t []time.Time) *zerolog.Event    { return log.Info().Times(key, t) }
func Dur(key string, d time.Duration) *zerolog.Event    { return log.Info().Dur(key, d) }
func Durs(key string, d []time.Duration) *zerolog.Event { return log.Info().Durs(key, d) }
func TimeDiff(key string, t time.Time, start time.Time) *zerolog.Event {
	return log.Info().TimeDiff(key, t, start)
}
func Interface(key string, i interface{}) *zerolog.Event     { return log.Info().Interface(key, i) }
func IPAddr(key string, ip net.IP) *zerolog.Event            { return log.Info().IPAddr(key, ip) }
func IPPrefix(key string, pfx net.IPNet) *zerolog.Event      { return log.Info().IPPrefix(key, pfx) }
func MACAddr(key string, ha net.HardwareAddr) *zerolog.Event { return log.Info().MACAddr(key, ha) }
