package parser

import (
	// #include "./C/c_tokenizer.c"
	"C"
	"errors"
	"unicode/utf8"
	"unsafe"
)

// Parse convert query to query digest that masked literal.
func Parse(query string) (string, error) {
	queryLength := C.int(utf8.RuneCountInString(query))
	queryC := C.CString(query)
	var firstComment **C.char
	var buf *C.char

	defer C.free(unsafe.Pointer(queryC))
	defer C.free(unsafe.Pointer(firstComment))
	defer C.free(unsafe.Pointer(buf))

	if queryLength > C.get_query_digests_max_query_length() {
		return query, errors.New("Query length is over 65000 charactors")
	}

	queryDigestC := C.mysql_query_digest_and_first_comment(queryC, queryLength, firstComment, buf)
	queryDigest := C.GoString(queryDigestC)

	return queryDigest, nil
}
