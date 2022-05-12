package db

type Database struct {
	Username        string
	Db              string
	Host            string
	Password        string
	Port            int
	ReadTimeoutSec  int
	WriteTimeoutSec int
	TimeoutSec      int
}
