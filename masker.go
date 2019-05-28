package masker

import (
	// #include "./C/c_tokenizer.c"
	"C"
	"fmt"
	"unicode/utf8"
	"unsafe"
)

// Masker is masker interface
type Masker interface {
	mask(string) (string, error)
}

// MysqlMasker is masker for mysql
type MysqlMasker struct{}

func (m *MysqlMasker) mask(query string) (string, error) {
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

// Mask mask literal values in a query
func Mask(m Masker, query string) (string, error) {
	return m.mask(query)
}
