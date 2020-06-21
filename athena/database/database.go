package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "password"
	dbName     = "athena_db_go"
)

// DB is connection to DB
var DB *sql.DB

func init() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	DB, err = sql.Open("postgres", dbinfo)
	checkErr(err)
	err = DB.Ping()
	checkErr(err)
	prepareDB()
}

func prepareDB() (err error) {
	tables := []string{"departments", "appointments", "patients", "providers", "problems", "provider_types", "races"}
	sql := ""
	for _, t := range tables {
		str := `
			CREATE TABLE IF NOT EXISTS %s(
				id BIGSERIAL,
				remote_id varchar NOT NULL,
				payload jsonb DEFAULT '{}'::jsonb NOT NULL,
				updated_at timestamp without time zone NOT NULL,
				created_at timestamp without time zone NOT NULL
			);
			CREATE UNIQUE INDEX IF NOT EXISTS index_%s_on_remote_id ON public.%s USING btree (remote_id);
		`
		sql += fmt.Sprintf(str, t, t, t)
	}
	DB.QueryRow(sql)
	return nil
}

// Insert insert record
func Insert(toTable string, payload []byte, remoteID string) (err error) {
	preparedPayload := payload // payload.to_json.gsub("'", "''")
	values := fmt.Sprintf("'%s', '%s'::jsonb, NOW(), NOW()", remoteID, preparedPayload)
	rowSQL := `
		INSERT INTO %s (remote_id, payload, updated_at, created_at)
		VALUES (%s)
		ON CONFLICT(remote_id)
		DO
			UPDATE
				SET payload = EXCLUDED.payload, updated_at = NOW();
	`
	DB.QueryRow(fmt.Sprintf(rowSQL, toTable, values))
	return nil
}

// InsertWithPaging insert multiple pages
func InsertWithPaging() {
	// Add
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
