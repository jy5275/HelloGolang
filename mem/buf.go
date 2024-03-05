package mem

import (
	"unsafe"

	"github.com/jy5275/HelloGolang/c"
)

/*
	https://cdn.learnku.com/uploads/images/202205/23/58489/4FVN22kCa6.png

	Data:
	-----------------------------------------------------------
	|   invalid    |    valid       |          invalid        |
	-----------------------------------------------------------
	^0             ^head            ^length                   ^capacity
*/

type Buf struct {
	Next     *Buf           // 多个 buffer 用链表链接
	Capacity int            // 缓存容量大小
	length   int            // 有效数据长度
	head     int            // 未处理数据的头部位置索引
	data     unsafe.Pointer // 保存的数据地址
}

func NewBuf(size int) *Buf {
	return &Buf{
		Capacity: size,
		data:     c.Malloc(size),
	}
}

func (b *Buf) SetBytes(src []byte) {
	c.Memcpy(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), src, len(src))
	b.length += len(src)
}
