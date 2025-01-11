package id_generator

import (
	"github.com/sony/sonyflake"
	"time"
)

func InitIDGenerator() *sonyflake.Sonyflake {
	return sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now(),
	})
}
