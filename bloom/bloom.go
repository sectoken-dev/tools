package bloom_filter

import (
	"github.com/garyburd/redigo/redis"
	"math/big"
)



type Bloom struct {
	r redis.Conn
}

var hashFuncs = []func(string) *big.Int{
	rs_hash, js_hash, pjw_hash, elf_hash, bkdr_hash, sdbm_hash, djb_hash, dek_hash,
}

func random_generator(hash_value *big.Int) *big.Int {
	return hash_value.Mod(hash_value, big.NewInt(int64(1<<30)))
}

func (b *Bloom) Update(key, item string) error {
	// 检查是否是新的条目，是新条目则更新bitmap并返回True，是重复条目则返回False
	for _, _func := range hashFuncs {
		hash_value := _func(item)
		real_value := random_generator(hash_value)
		res, err := redis.Int(b.r.Do("GETBIT", key, real_value))
		if err != nil {
			return err
		}
		if res == 1 {
			continue
		}
		_, err = b.r.Do("SETBIT", key, real_value, 1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bloom) IsExist(key, item string) (bool, error) {
	for _, _func := range hashFuncs {
		hash_value := _func(item)
		real_value := random_generator(hash_value)
		res, err := redis.Int(b.r.Do("GETBIT", key, real_value))
		if err != nil {
			return false, err
		}
		if res == 0 {
			return false, nil
		}
	}
	return true, nil
}

func NewBloom(r redis.Conn) *Bloom {
	b := new(Bloom)
	b.r = r
	return b
}

