package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
)

func GetFileByte(filepath string, offset, length int64) (data []byte, err error) {
	var file *os.File
	if file, err = os.Open(filepath); err == nil {
		buf := make([]byte, length)
		var n int
		if n, err = file.ReadAt(buf, offset); err == nil {
			data = buf[:n]
		}
		_ = file.Close()
	}
	return
}

func BytesReverse(data []byte) []byte {
	for from, to := 0, len(data)-1; from < to; from, to = from+1, to-1 {
		data[from], data[to] = data[to], data[from]
	}
	return data
}

func MergeBytes(pBytes ...[]byte) (data []byte) {
	l := len(pBytes)
	s := make([][]byte, l)
	for index := 0; index < l; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	data = bytes.Join(s, sep)
	return
}

type Octets struct {
	Data []byte
}

func (e *Octets) AddBytes(val []byte, addLen bool, fill int) {
	l := len(val)
	if fill > 0 && fill-l > 0 {
		f := make([]byte, fill-l)
		val = MergeBytes(val, f)
	}
	if addLen {
		e.Data = MergeBytes(e.Data, []byte{byte(len(val))}, val)
	} else {
		e.Data = MergeBytes(e.Data, val)
	}
}

func (e *Octets) GetBytes() []byte {
	return e.Data
}

func main() {
	b1, _ := GetFileByte("./gamed/config/elements.data", 0, 4)
	b2, _ := GetFileByte("./gamed/config/tasks.data", 0, 1)
	b3, _ := GetFileByte("./gamed/config/gshop.data", 0, 4)
	b4, _ := GetFileByte("./gamed/config/gshop1.data", 0, 4)

	b1 = BytesReverse(b1)
	b3 = BytesReverse(b3)
	b4 = BytesReverse(b4)

	d := new(Octets)
	d.AddBytes(b1, false, 0)
	d.AddBytes(b2, false, 0)
	d.AddBytes(b3, false, 0)
	d.AddBytes(b4, false, 0)

	h := []byte(hex.EncodeToString(d.GetBytes()))

	fmt.Println(hex.EncodeToString(h))
}
