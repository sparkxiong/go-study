package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

/**
应该抛给上层，有完整的堆栈信息跟踪，利于业务排查
*/

func dao() error {
	return errors.Wrap(sql.ErrNoRows, "no result")
}

func query() error {
	return errors.WithMessage(dao(), "query failure ")
}

func main() {
	err := query()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}
