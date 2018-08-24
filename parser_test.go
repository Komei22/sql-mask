package parser

import "testing"

func TestPerseValidQuery(t *testing.T) {
	querys := []string{
		"SELECT * FROM user WHERE id = 1",
		"INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES ('test', 'test', '2018-08-23 03:56:44', '2018-08-23 03:56:44')",
		"UPDATE `articles` SET `content` = '12345', `updated_at` = '2018-08-23 03:57:53' WHERE `articles`.`id` = 4",
		"DELETE FROM `articles` WHERE `articles`.`id` = 4",
	}

	expect_query_digests := []string{
		"SELECT * FROM user WHERE id = ?",
		"INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)",
		"UPDATE `articles` SET `content` = ?, `updated_at` = ? WHERE `articles`.`id` = ?",
		"DELETE FROM `articles` WHERE `articles`.`id` = ?",
	}

	for i := 0; i < len(querys); i++ {
		query_digest := Parse(querys[i])
		if query_digest != expect_query_digests[i] {
			t.Errorf(" Query digest of \"%s\" does not match \"%s\". ", querys[i], expect_query_digests[i])
			t.Errorf("%s", query_digest)
		}
	}

}
