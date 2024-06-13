package lib

type ByteToStringConverter struct {
	bytesStack []byte
}

func (b *ByteToStringConverter) Add(byte byte) {
	b.bytesStack = append(b.bytesStack, byte)
}

func (b *ByteToStringConverter) Convert() []rune {
	return b.convert()
}

func (b *ByteToStringConverter) ConvertAll() []rune {
	res := b.convert()
	res = append(res, []rune(string(b.bytesStack))...)
	b.bytesStack = []byte{}

	return res
}

func (b *ByteToStringConverter) convert() []rune {
	var result []rune
	var stack []byte

	globalOffset := 0
	offset := 0
	for i := 0; i < len(b.bytesStack); i++ {
		curBt := b.bytesStack[i]

		offset++

		if (curBt & 0b11010000) == 0b11010000 {
			globalOffset += len(stack)
			result = append(result, []rune(string(stack))...)
			stack = []byte{}
			offset = 0
		}

		stack = append(stack, curBt)

		if offset == 8 {
			globalOffset += len(stack)
			result = append(result, []rune(string(stack))...)
			stack = []byte{}
			offset = 0
			continue
		}
	}

	b.bytesStack = b.bytesStack[globalOffset:]

	return result
}
