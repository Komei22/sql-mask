package masker

import (
	"testing"
)

func TestMaskValidPgQuery(t *testing.T) {
	queries := []string{
		`SELECT * FROM "user" WHERE "id" = 1;`,
		`INSERT INTO "articles" ("title", "content", "created_at", "updated_at") VALUES ('test', 'test', '2018-08-23 03:56:44', '2018-08-23 03:56:44') RETURNING *;`,
		`UPDATE "articles" SET "content" = '12345', "updated_at" = '2018-08-23 03:57:53' WHERE "id" = 4;`,
		`DELETE FROM "articles" WHERE "id" IN (SELECT "id" FROM "articles" WHERE "id" = '4' LIMIT 1)`,
	}

	expectQueryDigests := []string{
		`SELECT * FROM "user" WHERE "id" = ?;`,
		`INSERT INTO "articles" ("title", "content", "created_at", "updated_at") VALUES (?, ?, ?, ?) RETURNING *;`,
		`UPDATE "articles" SET "content" = ?, "updated_at" = ? WHERE "id" = ?;`,
		`DELETE FROM "articles" WHERE "id" IN (SELECT "id" FROM "articles" WHERE "id" = ? LIMIT ?)`,
	}

	m := &PgMasker{}
	for i := 0; i < len(queries); i++ {
		queryDigest, _ := Mask(m, queries[i])
		if queryDigest != expectQueryDigests[i] {
			t.Errorf(" Query digest of \"%s\" does not match \"%s\". ", queries[i], expectQueryDigests[i])
			t.Errorf("QueryDigest is \"%s\"", queryDigest)
		}
	}
}

func TestMaskMultiVytePgQuery(t *testing.T) {
	query := `SELECT * FROM "user" WHERE "name" = '太郎'`
	expectQueryDigest := `SELECT * FROM "user" WHERE "name" = ?`

	m := &PgMasker{}
	queryDigest, _ := Mask(m, query)

	if queryDigest != expectQueryDigest {
		t.Errorf(" Query digest of \"%s\" does not match \"%s\". ", query, expectQueryDigest)
		t.Errorf("QueryDigest is \"%s\"", queryDigest)
	}
}

func TestMaskInvalidPgQuery(t *testing.T) {
	query := `INSERT INTO "articles" ("title", "content", "created_at", "updated_at") VALUES (test, test, '2018-08-23 03:56:44', '2018-08-23 03:56:44') RETURNING *;`
	expectQueryDigest := `INSERT INTO "articles" ("title", "content", "created_at", "updated_at") VALUES (test, test, ?, ?) RETURNING *;`

	m := &PgMasker{}
	queryDigest, _ := Mask(m, query)

	if queryDigest != expectQueryDigest {
		t.Errorf(" Query digest of \"%s\" does not match \"%s\". ", query, expectQueryDigest)
		t.Errorf("QueryDigest is \"%s\"", queryDigest)
	}
}
