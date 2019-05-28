package masker

import (
	"strings"
	"testing"
)

func TestMaskValidMysqlQuery(t *testing.T) {
	queries := []string{
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

	m := &MysqlMasker{}
	for i := 0; i < len(queries); i++ {
		queryDigest, _ := Mask(m, queries[i])
		if queryDigest != expectQueryDigests[i] {
			t.Errorf(" Query digest of \"%s\" does not match \"%s\". ", queries[i], expectQueryDigests[i])
			t.Errorf("QueryDigest is \"%s\"", queryDigest)
		}
	}
}

func TestMaskMultiByteMysqlQuery(t *testing.T) {
	query := "SELECT * FROM user WHERE name = '太郎'"
	expectQueryDigest := "SELECT * FROM user WHERE name = ?"

	m := &MysqlMasker{}
	queryDigest, _ := Mask(m, query)

	if queryDigest != expectQueryDigest {
		t.Errorf(" Query digest of \"%s\" does not match \"%s\". ", query, expectQueryDigest)
		t.Errorf("QueryDigest is \"%s\"", queryDigest)
	}
}

func TestMaskInvalidMysqlQuery(t *testing.T) {
	query := "INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES (test, test, '2018-08-23 03:56:44', '2018-08-23 03:56:44')"
	expectQueryDigest := "INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES (test, test, ?, ?)"

	m := &MysqlMasker{}
	queryDigest, _ := Mask(m, query)

	if queryDigest != expectQueryDigest {
		t.Errorf(" Query digest of \"%s\" does not match \"%s\". ", query, expectQueryDigest)
		t.Errorf("QueryDigest is \"%s\"", queryDigest)
	}
}

func TestMaskTooLongMysqlQuery(t *testing.T) {
	query := strings.Repeat("SELECT * FROM user WHERE id = 1;", 3000)

	m := &MysqlMasker{}
	_, err := Mask(m, query)

	if err == nil {
		t.Error("Should be error")
	}
}
