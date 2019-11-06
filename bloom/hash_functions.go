package bloom_filter

import (
	"math/big"
)

func rsHash(key string) *big.Int {
	a := big.NewInt(378551)
	b := big.NewInt(63689)
	hash_value := big.NewInt(0)
	for _, i := range key {
		hash_value.Mul(hash_value, a).Add(hash_value, big.NewInt(int64(rune(i))))
		a.Mul(a, b)
	}
	return hash_value
}

func jsHash(key string) *big.Int {
	hash_value := big.NewInt(1315423911)
	for _, i := range key {
		// part 1, 2, 3对应(hash_value << 5)，int(rune(i))， (hash_value >> 2)
		// final对应等号右面括号
		part_1 := big.NewInt(0)
		part_2 := big.NewInt(int64(rune(i)))
		part_3 := big.NewInt(0)
		final := big.NewInt(0)
		hash_value.Xor(hash_value, final.Add(part_1.Lsh(hash_value, 5), part_2).Add(final,
			part_3.Rsh(hash_value, 2)))
	}
	return hash_value
}

func pjwHash(key string) *big.Int {
	high_bits := big.NewInt(0)
	hex_flag_1, _ := new(big.Int).SetString("FFFFFFFF", 16)
	high_bits.Lsh(hex_flag_1, 28)
	hash_value := big.NewInt(0)
	test := big.NewInt(0)
	for _, i := range key {
		hash_value.Lsh(hash_value, 4).Add(hash_value, big.NewInt(int64(rune(i))))
		test.And(hash_value, high_bits)
	}
	if test != big.NewInt(0) {
		hash_value.And(hash_value.Xor(hash_value, test.Rsh(test, 24)), high_bits.Not(high_bits))
	}
	hex_flag_2, _ := new(big.Int).SetString("7FFFFFFF", 16)
	return hash_value.And(hash_value, hex_flag_2)
}

func elfHash(key string) *big.Int {
	hash_value := big.NewInt(0)
	for _, i := range key {
		hash_value.Add(hash_value.Lsh(hash_value, 4), big.NewInt(int64(rune(i))))
		x := big.NewInt(0)
		hex_flag, _ := new(big.Int).SetString("F0000000", 16)
		x.And(hash_value, hex_flag)
		cmp := x.Cmp(big.NewInt(0))
		if cmp != 0 {
			x_temp := big.NewInt(0)
			hash_value.Xor(hash_value, x_temp.Rsh(x, 24))
		}
		hash_value.And(hash_value, x.Not(x))
	}
	return hash_value
}

func bkdrHash(key string) *big.Int {
	seed := big.NewInt(int64(131))
	hash_value := big.NewInt(0)
	for _, i := range key {
		hash_value.Add(hash_value.Mul(hash_value, seed), big.NewInt(int64(rune(i))))
	}
	return hash_value
}

func sdbmHash(key string) *big.Int {
	hash_value := big.NewInt(0)
	for _, i := range key {
		// hash_value = int(rune(i)) + (hash_value << 6) + (hash_value << 16) - hash_value
		part_1 := big.NewInt(0) // hash_value << 6
		part_2 := big.NewInt(0) // hash_value << 16
		add_1 := big.NewInt(0)
		add_2 := big.NewInt(0)
		add_1.Add(big.NewInt(int64(rune(i))), part_1.Lsh(hash_value, 6))
		add_2.Sub(part_2.Lsh(hash_value, 16), hash_value)
		hash_value.Add(add_1, add_2)
	}
	return hash_value
}

func djbHash(key string) *big.Int {
	hash_value := big.NewInt(int64(5381))
	for _, i := range key {
		part_1 := big.NewInt(0)
		part_2 := big.NewInt(0)
		hash_value.Add(part_2.Add(part_1.Lsh(hash_value, 5), hash_value), big.NewInt(int64(rune(i))))
	}
	return hash_value
}

func dekHash(key string) *big.Int {
	hash_value := big.NewInt(int64(len(key)))
	for _, i := range key {
		part_1 := big.NewInt(0)
		part_2 := big.NewInt(0)
		part_3 := big.NewInt(0)
		hash_value.Xor(part_3.Xor(part_1.Lsh(hash_value, 5), part_2.Rsh(hash_value, 27)),
			big.NewInt(int64(rune(i))))
	}
	// fmt.Println(hash_value)
	return hash_value
}

