package debug

import (
    "fmt"
    "net"
    "time"

    "github.com/u6du/zerolog"
    "github.com/u6du/zerolog/log"
)

func Msg(msg string) {
    log.Debug().Out(msg)
}

func Msgf(format string, v ...interface{}) {
    log.Debug().Out(fmt.Sprintf(format, v...))
}

func Fields(fields map[string]interface{}) *zerolog.Event {
    return log.Debug().Fields(fields)
}

func Dict(key string, dict *zerolog.Event) *zerolog.Event {
    return log.Debug().Dict(key, dict)
}

func Array(key string, arr zerolog.LogArrayMarshaler) *zerolog.Event {
    return log.Debug().Array(key, arr)
}

func Object(key string, obj zerolog.LogObjectMarshaler) *zerolog.Event {
    return log.Debug().Object(key, obj)
}

func EmbedObject(obj zerolog.LogObjectMarshaler) *zerolog.Event {
    return log.Debug().EmbedObject(obj)
}

func Str(key, val string) *zerolog.Event {
    return log.Debug().Str(key, val)
}

func Strs(key string, vals []string) *zerolog.Event {
    return log.Debug().Strs(key, vals)
}

func Bytes(key string, val []byte) *zerolog.Event {
    return log.Debug().Bytes(key, val)
}

func Hex(key string, val []byte) *zerolog.Event {
    return log.Debug().Hex(key, val)
}

func RawJSON(key string, b []byte) *zerolog.Event {
    return log.Debug().RawJSON(key, b)
}

func AnErr(key string, err error) *zerolog.Event {
    return log.Debug().AnErr(key, err)
}

func Errs(key string, errs []error) *zerolog.Event {
    return log.Debug().Errs(key, errs)
}

func Err(err error) *zerolog.Event {
    return log.Debug().Err(err)
}

func Bool(key string, b bool) *zerolog.Event {
    return log.Debug().Bool(key, b)
}

func Bools(key string, b []bool) *zerolog.Event {
    return log.Debug().Bools(key, b)
}

func Int(key string, i int) *zerolog.Event {
    return log.Debug().Int(key, i)
}

func Ints(key string, i []int) *zerolog.Event {
    return log.Debug().Ints(key, i)
}

func Int8(key string, i int8) *zerolog.Event {
    return log.Debug().Int8(key, i)
}

func Ints8(key string, i []int8) *zerolog.Event {
    return log.Debug().Ints8(key, i)
}

func Int16(key string, i int16) *zerolog.Event {
    return log.Debug().Int16(key, i)
}

func Ints16(key string, i []int16) *zerolog.Event {
    return log.Debug().Ints16(key, i)
}

func Int32(key string, i int32) *zerolog.Event {
    return log.Debug().Int32(key, i)
}

func Ints32(key string, i []int32) *zerolog.Event {
    return log.Debug().Ints32(key, i)
}

func Int64(key string, i int64) *zerolog.Event {
    return log.Debug().Int64(key, i)
}

func Ints64(key string, i []int64) *zerolog.Event {
    return log.Debug().Ints64(key, i)
}

func Uint(key string, i uint) *zerolog.Event {
    return log.Debug().Uint(key, i)
}

func Uints(key string, i []uint) *zerolog.Event {
    return log.Debug().Uints(key, i)
}

func Uint8(key string, i uint8) *zerolog.Event {
    return log.Debug().Uint8(key, i)
}

func Uints8(key string, i []uint8) *zerolog.Event {
    return log.Debug().Uints8(key, i)
}

func Uint16(key string, i uint16) *zerolog.Event {
    return log.Debug().Uint16(key, i)
}

func Uints16(key string, i []uint16) *zerolog.Event {
    return log.Debug().Uints16(key, i)
}

func Uint32(key string, i uint32) *zerolog.Event {
    return log.Debug().Uint32(key, i)
}

func Uints32(key string, i []uint32) *zerolog.Event {
    return log.Debug().Uints32(key, i)
}

func Uint64(key string, i uint64) *zerolog.Event {
    return log.Debug().Uint64(key, i)
}

func Uints64(key string, i []uint64) *zerolog.Event {
    return log.Debug().Uints64(key, i)
}

func Float32(key string, f float32) *zerolog.Event {
    return log.Debug().Float32(key, f)
}

func Floats32(key string, f []float32) *zerolog.Event {
    return log.Debug().Floats32(key, f)
}

func Float64(key string, f float64) *zerolog.Event {
    return log.Debug().Float64(key, f)
}

func Floats64(key string, f []float64) *zerolog.Event {
    return log.Debug().Floats64(key, f)
}

func Time(key string, t time.Time) *zerolog.Event {
    return log.Debug().Time(key, t)
}

func Times(key string, t []time.Time) *zerolog.Event {
    return log.Debug().Times(key, t)
}

func Dur(key string, d time.Duration) *zerolog.Event {
    return log.Debug().Dur(key, d)
}

func Durs(key string, d []time.Duration) *zerolog.Event {
    return log.Debug().Durs(key, d)
}

func TimeDiff(key string, t time.Time, start time.Time) *zerolog.Event {
    return log.Debug().TimeDiff(key, t, start)
}

func Interface(key string, i interface{}) *zerolog.Event {
    return log.Debug().Interface(key, i)
}

func IPAddr(key string, ip net.IP) *zerolog.Event {
    return log.Debug().IPAddr(key, ip)
}

func IPPrefix(key string, pfx net.IPNet) *zerolog.Event {
    return log.Debug().IPPrefix(key, pfx)
}

func MACAddr(key string, ha net.HardwareAddr) *zerolog.Event {
    return log.Debug().MACAddr(key, ha)
}


