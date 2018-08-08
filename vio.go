package liblsdj

import (
	"fmt"
	"io/ioutil"
)

type vio struct {
	data []byte
	cur  int
}

func (v *vio) open(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	v.data = data
}

func (v *vio) readByte() byte {
	out := v.data[v.cur]
	v.cur += 1
	return out
}

func (v *vio) read(n int) []byte {
	out := v.data[v.cur:(v.cur + n)]
	v.cur += n
	return out
}

func (v *vio) writeByte(b byte) {
	v.data = append(v.data, b)
	v.cur++
}

func (v *vio) write(b []byte) {
	v.data = append(v.data, b...)
	v.cur += len(b)
}

func (v *vio) seek(offset int) {
	if len(v.data) < offset {
		v.cur = len(v.data)
	} else {
		v.cur = offset
	}
}

func (v *vio) seekCur(offset int) {
	if len(v.data) < offset {
		v.cur = len(v.data)
	} else {
		v.cur += offset
	}
}

func (v *vio) get() []byte {
	return v.data
}

func (v *vio) getCur() int {
	return v.cur
}
