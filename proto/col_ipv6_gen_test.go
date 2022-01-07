// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-faster/ch/internal/gold"
)

func TestColIPv6_DecodeColumn(t *testing.T) {
	const rows = 50
	var data ColIPv6
	for i := 0; i < rows; i++ {
		v := IPv6FromInt(i)
		data.Append(v)
		require.Equal(t, v, data.Row(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)
	t.Run("Golden", func(t *testing.T) {
		gold.Bytes(t, buf.Buf, "col_ipv6")
	})
	t.Run("Ok", func(t *testing.T) {
		br := bytes.NewReader(buf.Buf)
		r := NewReader(br)

		var dec ColIPv6
		require.NoError(t, dec.DecodeColumn(r, rows))
		require.Equal(t, data, dec)
		require.Equal(t, rows, dec.Rows())
		dec.Reset()
		require.Equal(t, 0, dec.Rows())
		require.Equal(t, ColumnTypeIPv6, dec.Type())
	})
	t.Run("ZeroRows", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		var dec ColIPv6
		require.NoError(t, dec.DecodeColumn(r, 0))
	})
	t.Run("ErrUnexpectedEOF", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		var dec ColIPv6
		require.ErrorIs(t, dec.DecodeColumn(r, rows), io.ErrUnexpectedEOF)
	})
	t.Run("NoShortRead", func(t *testing.T) {
		var dec ColIPv6
		requireNoShortRead(t, buf.Buf, colAware(&dec, rows))
	})
	t.Run("ZeroRowsEncode", func(t *testing.T) {
		var v ColIPv6
		v.EncodeColumn(nil) // should be no-op
	})
}

func TestColIPv6Array(t *testing.T) {
	const rows = 50
	data := NewArrIPv6()
	for i := 0; i < rows; i++ {
		data.AppendIPv6([]IPv6{
			IPv6FromInt(i),
			IPv6FromInt(i + 1),
			IPv6FromInt(i + 2),
		})
	}

	var buf Buffer
	data.EncodeColumn(&buf)
	t.Run("Golden", func(t *testing.T) {
		gold.Bytes(t, buf.Buf, "col_arr_ipv6")
	})
	t.Run("Ok", func(t *testing.T) {
		br := bytes.NewReader(buf.Buf)
		r := NewReader(br)

		dec := NewArrIPv6()
		require.NoError(t, dec.DecodeColumn(r, rows))
		require.Equal(t, data, dec)
		require.Equal(t, rows, dec.Rows())
		dec.Reset()
		require.Equal(t, 0, dec.Rows())
		require.Equal(t, ColumnTypeIPv6.Array(), dec.Type())
	})
	t.Run("ErrUnexpectedEOF", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		dec := NewArrIPv6()
		require.ErrorIs(t, dec.DecodeColumn(r, rows), io.ErrUnexpectedEOF)
	})
}

func BenchmarkColIPv6_DecodeColumn(b *testing.B) {
	const rows = 1_000
	var data ColIPv6
	for i := 0; i < rows; i++ {
		data = append(data, IPv6FromInt(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	br := bytes.NewReader(buf.Buf)
	r := NewReader(br)

	var dec ColIPv6
	if err := dec.DecodeColumn(r, rows); err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		br.Reset(buf.Buf)
		r.raw.Reset(br)
		dec.Reset()

		if err := dec.DecodeColumn(r, rows); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkColIPv6_EncodeColumn(b *testing.B) {
	const rows = 1_000
	var data ColIPv6
	for i := 0; i < rows; i++ {
		data = append(data, IPv6FromInt(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		data.EncodeColumn(&buf)
	}
}
