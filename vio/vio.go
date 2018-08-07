package vio

import (
	"fmt"
	"io/ioutil"
)

type Vio struct {
	data []byte
	cur  int
}

func (v *Vio) Open(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	v.data = data
}

func (v *Vio) ReadSingle() byte {
	out := v.data[v.cur]
	v.cur += 1
	return out
}

func (v *Vio) Read(n int) []byte {
	out := make([]byte, n)
	copy(out, v.data[v.cur:(v.cur+n)])
	v.cur += n
	return out
}

func (v *Vio) WriteSingle(b byte) {
	v.data = append(v.data, b)
	v.cur++
}

func (v *Vio) Write(b []byte) {
	v.data = append(v.data, b...)
	v.cur += len(b)
}

func (v *Vio) Seek(offset int) {
	if len(v.data) < offset {
		v.cur = len(v.data)
	} else {
		v.cur = offset
	}
}

func (v *Vio) Finalize(max int) {
	if len(v.data) < max {
		b := make([]byte, max-len(v.data))
		v.Write(b)
	}
}

func (v *Vio) Get() []byte {
	return v.data
}

func (v *Vio) Cur() int {
	return v.cur
}
