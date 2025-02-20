package main

// TC: O(1)
// SC: O(1)
func reverseBits1(num uint32) uint32 {
	var result, power uint32
	power = 31
	for num != 0 {
		result += (num & 1) << power
		power--
		num >>= 1
	}

	return result
}

// TC: O(1)
// SC: O(1)
func reverseBits2(num uint32) uint32 {
	result := uint64(0)
	power := uint64(24)
	cache := make(map[uint32]uint64)

	for num != 0 {
		result += reverseByte(num&0xff, cache) << power
		num >>= 8
		power -= 8
	}

	return uint32(result)
}

func reverseByte(b uint32, cache map[uint32]uint64) uint64 {
	v, ok := cache[b]
	if ok {
		return v
	}

	v = (uint64(b) * 0x010884422010 & 0x010884422010) % 1023
	cache[b] = v

	return v
}

// TC: O(1)
// SC: O(1)
func reverseBits3(num uint32) uint32 {
	num = (num >> 16) | (num << 16)                             // break the original 32-bit into 2 blocks and switch them
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8) // break the 16 bit blocks
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)

	return num
}
