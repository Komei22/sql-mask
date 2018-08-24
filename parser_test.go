package parser

import (
	"strings"
	"testing"
)

func TestPerseValidQuery(t *testing.T) {
	querys := []string{
		"SELECT * FROM user WHERE id = 1",
		"INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES ('test', 'test', '2018-08-23 03:56:44', '2018-08-23 03:56:44')",
		"UPDATE `articles` SET `content` = '12345', `updated_at` = '2018-08-23 03:57:53' WHERE `articles`.`id` = 4",
		"DELETE FROM `articles` WHERE `articles`.`id` = 4",
	}

	expectQueryDigests := []string{
		"SELECT * FROM user WHERE id = ?",
		"INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)",
		"UPDATE `articles` SET `content` = ?, `updated_at` = ? WHERE `articles`.`id` = ?",
		"DELETE FROM `articles` WHERE `articles`.`id` = ?",
	}

	for i := 0; i < len(querys); i++ {
		queryDigest, _ := Parse(querys[i])
		if queryDigest != expectQueryDigests[i] {
			t.Errorf(" Query digest of \"%s\" does not match \"%s\". ", querys[i], expectQueryDigests[i])
			t.Errorf("%s", queryDigest)
		}
	}
}

func TestParseMultiByteQuery(t *testing.T) {
	query := "SELECT * FROM user WHERE name = '太郎'"
	expectQueryDigest := "SELECT * FROM user WHERE name = ?"

	queryDigest, _ := Parse(query)

	if queryDigest != expectQueryDigest {
		t.Errorf(" Query digest of \"%s\" does not match \"%s\". ", query, expectQueryDigest)
		t.Errorf("%s", queryDigest)
	}
}

func TestParseNonValidQuery(t *testing.T) {

}

func TestParseTooLongQuery(t *testing.T) {
	query := strings.Repeat("SELECT * FROM user WHERE id = 1;", 3000)

	_, err := Parse(query)

	if err == nil {
		t.Error("Should be error")
	}
}
