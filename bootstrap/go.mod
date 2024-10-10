module COREMOD

go 1.23.1

replace github.com/tabajara-coder/backend-base-structure => ../

require github.com/lib/pq v1.10.9

require github.com/mattn/go-sqlite3 v1.14.22 // indirect

require (
	github.com/anthdm/superkit v0.0.0-20240701091803-e7f8e0aad3e9
	github.com/go-chi/chi/v5 v5.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/tabajara-coder/backend-base-structure v0.0.0-20241001034707-e2ea6a9ac521 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gorm.io/driver/postgres v1.5.9 // indirect
	gorm.io/driver/sqlite v1.5.6
	gorm.io/gorm v1.25.10 // indirect
)
