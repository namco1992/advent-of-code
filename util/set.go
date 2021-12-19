package util

type CharSet map[int32]struct{}

func NewSet(s string) CharSet {
	set := make(CharSet)
	for _, i := range s {
		set[i] = struct{}{}
	}
	return set
}

func Intersection(s1, s2 CharSet) CharSet {
	ret := make(CharSet)
	for i := range s2 {
		_, ok := s1[i]
		if ok {
			ret[i] = struct{}{}
		}
	}
	return ret
}
