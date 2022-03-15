package database // 0-> move to driver

type SQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Rows, error)
	QueryRow(string, ...interface{}) Row
	// Prepare(string, ...interface{}) (*sql.Stmt, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Err() error
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

// type Tx interface {
// 	Commit() error
// 	Rollback() error
// 	Stmt(stmt *sql.Stmt) *sql.Stmt
// 	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
// }

// type DB interface {
// 	Stmt
// 	Begin() (*sql.Tx, error)
// 	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
// 	Close() error
// 	PingContext(ctx context.Context) error
// 	SetConnMaxLifetime(d time.Duration)
// 	SetMaxIdleConns(n int)
// 	SetMaxOpenConns(n int)
// 	Stats() sql.DBStats
// }
