package concat_slices

import "bytes"

func concatCopyPreAllocate(slices [][]byte) []byte {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	tmp := make([]byte, totalLen)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}

func concatWriteBuffer(slices [][]byte) []byte {
	var buf bytes.Buffer
	for _, s := range slices {
		buf.Write(s)
	}
	return buf.Bytes()
}

func concatAppend(slices [][]byte) []byte {
	var tmp []byte
	for _, s := range slices {
		tmp = append(tmp, s...)
	}
	return tmp
}
