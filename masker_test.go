package masker

import (
	"strings"
	"testing"
)

func TestMysqlMasker_mask(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Valid query : SELECT",
			args:    args{query: "SELECT * FROM user WHERE id = 1"},
			want:    "SELECT * FROM user WHERE id = ?",
			wantErr: false,
		},
		{
			name:    "Valid query : INSERT",
			args:    args{query: "INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES ('test', 'test', '2018-08-23 03:56:44', '2018-08-23 03:56:44')"},
			want:    "INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)",
			wantErr: false,
		},
		{
			name:    "Valid query : UPDATE",
			args:    args{query: "UPDATE `articles` SET `content` = '12345', `updated_at` = '2018-08-23 03:57:53' WHERE `articles`.`id` = 4"},
			want:    "UPDATE `articles` SET `content` = ?, `updated_at` = ? WHERE `articles`.`id` = ?",
			wantErr: false,
		},
		{
			name:    "Valid query : DELETE",
			args:    args{query: "DELETE FROM `articles` WHERE `articles`.`id` = 4"},
			want:    "DELETE FROM `articles` WHERE `articles`.`id` = ?",
			wantErr: false,
		},
		{
			name:    "Multi byte query",
			args:    args{query: "SELECT * FROM user WHERE name = '太郎'"},
			want:    "SELECT * FROM user WHERE name = ?",
			wantErr: false,
		},
		{
			name:    "Invalid query",
			args:    args{query: "INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES (test, test, '2018-08-23 03:56:44', '2018-08-23 03:56:44')"},
			want:    "INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES (test, test, ?, ?)",
			wantErr: false,
		},
		{
			name:    "Too long invalid query",
			args:    args{query: strings.Repeat("SELECT * FROM user WHERE id = 1;", 3000)},
			want:    strings.Repeat("SELECT * FROM user WHERE id = 1;", 3000),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MysqlMasker{}
			got, err := m.mask(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMasker.mask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MysqlMasker.mask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgMasker_mask(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Valid query : SELECT",
			args:    args{query: `SELECT * FROM "user" WHERE "id" = 1;`},
			want:    `SELECT * FROM "user" WHERE "id" = ?;`,
			wantErr: false,
		},
		{
			name:    "Valid query : INSERT",
			args:    args{query: `INSERT INTO "articles" ("title", "content", "created_at", "updated_at") VALUES ('test', 'test', '2018-08-23 03:56:44', '2018-08-23 03:56:44') RETURNING *;`},
			want:    `INSERT INTO "articles" ("title", "content", "created_at", "updated_at") VALUES (?, ?, ?, ?) RETURNING *;`,
			wantErr: false,
		},
		{
			name:    "Valid query : UPDATE",
			args:    args{query: `UPDATE "articles" SET "content" = '12345', "updated_at" = '2018-08-23 03:57:53' WHERE "id" = 4;`},
			want:    `UPDATE "articles" SET "content" = ?, "updated_at" = ? WHERE "id" = ?;`,
			wantErr: false,
		},
		{
			name:    "Valid query : DELETE",
			args:    args{query: `DELETE FROM "articles" WHERE "id" IN (SELECT "id" FROM "articles" WHERE "id" = '4' LIMIT 1)`},
			want:    `DELETE FROM "articles" WHERE "id" IN (SELECT "id" FROM "articles" WHERE "id" = ? LIMIT ?)`,
			wantErr: false,
		},
		{
			name:    "Multi byte query",
			args:    args{query: `SELECT * FROM "user" WHERE "name" = '太郎'`},
			want:    `SELECT * FROM "user" WHERE "name" = ?`,
			wantErr: false,
		},
		{
			name:    "Invalid query",
			args:    args{query: `INSERT INTO "articles" ("title", "content", "created_at", "updated_at") VALUES (test, test, '2018-08-23 03:56:44', '2018-08-23 03:56:44') RETURNING *;`},
			want:    `INSERT INTO "articles" ("title", "content", "created_at", "updated_at") VALUES (test, test, ?, ?) RETURNING *;`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PgMasker{}
			got, err := p.mask(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgMasker.mask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PgMasker.mask() = %v, want %v", got, tt.want)
			}
		})
	}
}
