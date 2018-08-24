package parser

import (
	// #include "./C/c_tokenizer.c"
	"C"
	"unsafe"
)

func Parse(query string) string {
	query_length := C.int(len(query))
	query_c := C.CString(query)
	var first_comment **C.char = nil
	var buf *C.char = nil

	query_digest_c := C.mysql_query_digest_and_first_comment(query_c, query_length, first_comment, buf)
	query_digest := C.GoString(query_digest_c)

	C.free(unsafe.Pointer(query_c))
	C.free(unsafe.Pointer(first_comment))
	C.free(unsafe.Pointer(buf))

	return query_digest
}
