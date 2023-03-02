package oidstorage

import (
	"fmt"
	"log"
	"os"

	"github.com/xujiajun/nutsdb"
)

type DB struct {
	db *nutsdb.DB
}

func InitializeDb() *DB {
	opt := nutsdb.DefaultOptions

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dirname + string(os.PathSeparator) + ".willow")
	opt.Dir = "./tmp/nutsdb"
	opt.SegmentSize = 8 * nutsdb.MB
	client, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatalln(err)
	}

	return &DB{
		db: client,
	}
}

func (d *DB) CloseDb() {
	d.db.Close()
}

func (d *DB) Test() *nutsdb.DB {
	return d.db
}

func (d *DB) Update(fn func(tx *nutsdb.Tx) error) error {
	return d.db.Update(fn)
}
