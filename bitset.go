package porcupine

type bitset []uint64

// data layout:
// bits 0-63 are in data[0], the next are in data[1], etc.

func newBitset(bits uint) bitset { _ = "STUB: not implemented"; return *new(bitset) }

func (b bitset) clone() bitset { _ = "STUB: not implemented"; return *new(bitset) }

func bitsetIndex(pos uint) (uint, uint) { _ = "STUB: not implemented"; return 0, 0 }

func (b bitset) set(pos uint) bitset { _ = "STUB: not implemented"; return *new(bitset) }

func (b bitset) clear(pos uint) bitset { _ = "STUB: not implemented"; return *new(bitset) }

func (b bitset) popcnt() uint { _ = "STUB: not implemented"; return 0 }

func (b bitset) hash() uint64 { _ = "STUB: not implemented"; return 0 }

func (b bitset) equals(b2 bitset) bool { _ = "STUB: not implemented"; return false }
