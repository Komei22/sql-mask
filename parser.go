package parser

import (
	// #include "./C/c_tokenizer.c"
	"C"
	"unicode/utf8"
	"unsafe"
)

func Parse(query string) string {
	query_length := C.int(utf8.RuneCountInString(query))
	query_c := C.CString(query)
	var first_comment **C.char = nil
	var buf *C.char = nil

	defer C.free(unsafe.Pointer(query_c))
	defer C.free(unsafe.Pointer(first_comment))
	defer C.free(unsafe.Pointer(buf))

	query_digest_c := C.mysql_query_digest_and_first_comment(query_c, query_length, first_comment, buf)
	query_digest := C.GoString(query_digest_c)

	return query_digest
}
