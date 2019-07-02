package warn

import (
    "fmt"
    "net"
    "time"

    "github.com/u6du/zerolog"
    "github.com/u6du/zerolog/log"
)

func Msg(msg string) {
    log.Warn().Out(msg)
}

func Msgf(format string, v ...interface{}) {
    log.Warn().Out(fmt.Sprintf(format, v...))
}

func Fields(fields map[string]interface{}) *zerolog.Event {
    return log.Warn().Fields(fields)
}

func Dict(key string, dict *zerolog.Event) *zerolog.Event {
    return log.Warn().Dict(key, dict)
}

func Array(key string, arr zerolog.LogArrayMarshaler) *zerolog.Event {
    return log.Warn().Array(key, arr)
}

func Object(key string, obj zerolog.LogObjectMarshaler) *zerolog.Event {
    return log.Warn().Object(key, obj)
}

func EmbedObject(obj zerolog.LogObjectMarshaler) *zerolog.Event {
    return log.Warn().EmbedObject(obj)
}

func Str(key, val string) *zerolog.Event {
    return log.Warn().Str(key, val)
}

func Strs(key string, vals []string) *zerolog.Event {
    return log.Warn().Strs(key, vals)
}

func Bytes(key string, val []byte) *zerolog.Event {
    return log.Warn().Bytes(key, val)
}

func Hex(key string, val []byte) *zerolog.Event {
    return log.Warn().Hex(key, val)
}

func RawJSON(key string, b []byte) *zerolog.Event {
    return log.Warn().RawJSON(key, b)
}

func AnErr(key string, err error) *zerolog.Event {
    return log.Warn().AnErr(key, err)
}

func Errs(key string, errs []error) *zerolog.Event {
    return log.Warn().Errs(key, errs)
}

func Err(err error) *zerolog.Event {
    return log.Warn().Err(err)
}

func Bool(key string, b bool) *zerolog.Event {
    return log.Warn().Bool(key, b)
}

func Bools(key string, b []bool) *zerolog.Event {
    return log.Warn().Bools(key, b)
}

func Int(key string, i int) *zerolog.Event {
    return log.Warn().Int(key, i)
}

func Ints(key string, i []int) *zerolog.Event {
    return log.Warn().Ints(key, i)
}

func Int8(key string, i int8) *zerolog.Event {
    return log.Warn().Int8(key, i)
}

func Ints8(key string, i []int8) *zerolog.Event {
    return log.Warn().Ints8(key, i)
}

func Int16(key string, i int16) *zerolog.Event {
    return log.Warn().Int16(key, i)
}

func Ints16(key string, i []int16) *zerolog.Event {
    return log.Warn().Ints16(key, i)
}

func Int32(key string, i int32) *zerolog.Event {
    return log.Warn().Int32(key, i)
}

func Ints32(key string, i []int32) *zerolog.Event {
    return log.Warn().Ints32(key, i)
}

func Int64(key string, i int64) *zerolog.Event {
    return log.Warn().Int64(key, i)
}

func Ints64(key string, i []int64) *zerolog.Event {
    return log.Warn().Ints64(key, i)
}

func Uint(key string, i uint) *zerolog.Event {
    return log.Warn().Uint(key, i)
}

func Uints(key string, i []uint) *zerolog.Event {
    return log.Warn().Uints(key, i)
}

func Uint8(key string, i uint8) *zerolog.Event {
    return log.Warn().Uint8(key, i)
}

func Uints8(key string, i []uint8) *zerolog.Event {
    return log.Warn().Uints8(key, i)
}

func Uint16(key string, i uint16) *zerolog.Event {
    return log.Warn().Uint16(key, i)
}

func Uints16(key string, i []uint16) *zerolog.Event {
    return log.Warn().Uints16(key, i)
}

func Uint32(key string, i uint32) *zerolog.Event {
    return log.Warn().Uint32(key, i)
}

func Uints32(key string, i []uint32) *zerolog.Event {
    return log.Warn().Uints32(key, i)
}

func Uint64(key string, i uint64) *zerolog.Event {
    return log.Warn().Uint64(key, i)
}

func Uints64(key string, i []uint64) *zerolog.Event {
    return log.Warn().Uints64(key, i)
}

func Float32(key string, f float32) *zerolog.Event {
    return log.Warn().Float32(key, f)
}

func Floats32(key string, f []float32) *zerolog.Event {
    return log.Warn().Floats32(key, f)
}

func Float64(key string, f float64) *zerolog.Event {
    return log.Warn().Float64(key, f)
}

func Floats64(key string, f []float64) *zerolog.Event {
    return log.Warn().Floats64(key, f)
}

func Time(key string, t time.Time) *zerolog.Event {
    return log.Warn().Time(key, t)
}

func Times(key string, t []time.Time) *zerolog.Event {
    return log.Warn().Times(key, t)
}

func Dur(key string, d time.Duration) *zerolog.Event {
    return log.Warn().Dur(key, d)
}

func Durs(key string, d []time.Duration) *zerolog.Event {
    return log.Warn().Durs(key, d)
}

func TimeDiff(key string, t time.Time, start time.Time) *zerolog.Event {
    return log.Warn().TimeDiff(key, t, start)
}

func Interface(key string, i interface{}) *zerolog.Event {
    return log.Warn().Interface(key, i)
}

func IPAddr(key string, ip net.IP) *zerolog.Event {
    return log.Warn().IPAddr(key, ip)
}

func IPPrefix(key string, pfx net.IPNet) *zerolog.Event {
    return log.Warn().IPPrefix(key, pfx)
}

func MACAddr(key string, ha net.HardwareAddr) *zerolog.Event {
    return log.Warn().MACAddr(key, ha)
}


