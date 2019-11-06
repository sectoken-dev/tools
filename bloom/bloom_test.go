package bloom_filter

import (
	"github.com/garyburd/redigo/redis"
	"testing"
)

const (
	REDIS_SERVER = "127.0.0.1"
	REDIS_PASS = "aaaaaaa"
)

func TestBloom_Update(t *testing.T) {
	conn, err := redis.Dial("tcp", REDIS_SERVER, redis.DialPassword(REDIS_PASS), redis.DialDatabase(0))
	if err != nil {
		return
	}

	bl := NewBloom(conn)
	bl.Update("k", "v")
}

func TestBloom_IsExist(t *testing.T) {
	conn, err := redis.Dial("tcp", REDIS_SERVER, redis.DialPassword(REDIS_PASS), redis.DialDatabase(0))
	if err != nil {
		return
	}

	bl := NewBloom(conn)
	bl.IsExist("k", "v")
}

