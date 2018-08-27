package parser

import (
	// #include "./C/c_tokenizer.c"
	"C"
	"fmt"
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

	queryDigestsMaxQueryLength := C.get_query_digests_max_query_length()
	if queryLength > queryDigestsMaxQueryLength {
		return query, fmt.Errorf("Query length is over %d charactors", queryDigestsMaxQueryLength)
	}

	queryDigestC := C.mysql_query_digest_and_first_comment(queryC, queryLength, firstComment, buf)
	queryDigest := C.GoString(queryDigestC)

	return queryDigest, nil
}
