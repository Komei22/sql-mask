package parser

import (
	// #include "./C/c_tokenizer.c"
	"C"
	"errors"
	"unicode/utf8"
	"unsafe"
)

func Parse(query string) (string, error) {
	query_length := C.int(utf8.RuneCountInString(query))
	query_c := C.CString(query)
	var first_comment **C.char = nil
	var buf *C.char = nil

	defer C.free(unsafe.Pointer(query_c))
	defer C.free(unsafe.Pointer(first_comment))
	defer C.free(unsafe.Pointer(buf))

	if query_length > C.get_query_digests_max_query_length() {
		return query, errors.New("Query length is over 65000 charactors")
	}

	query_digest_c := C.mysql_query_digest_and_first_comment(query_c, query_length, first_comment, buf)
	query_digest := C.GoString(query_digest_c)

	return query_digest, nil
}
