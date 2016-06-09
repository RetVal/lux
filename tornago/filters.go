package tornago

// All the possible base groups there are.
const (
	GroupNone uint16 = 0x0
	Group1    uint16 = 1 << iota
	Group2
	Group3
	Group4
	Group5
	Group6
	Group7
	Group8
	Group9
	Group10
	Group11
	Group12
	Group13
	Group14
	Group15
	GroupAll uint16 = 0xffff
)

// Group is a quick function to get a uint16 filter from an int
//    Group(i<0) = 0
//    Group(i), i E [0-15] = 0x1<<i
//    Group(i>15) = 0xFFFF
func Group(i int) uint16 {
	if i < 0 {
		return 0
	}
	if i > 15 {
		return 0xFFFF
	}
	return 0x1 << uint(i)
}

// Mask returns the same as Group
func Mask(i int) uint16 {
	return Group(i)
}
