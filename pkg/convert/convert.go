package convert

import "strconv"

//响应处理
type StrTo string

func (s StrTo) String () string {
	return string(s)
}

func (s StrTo) Int () (int, error) {
	return strconv.Atoi(s.String())
}

func (s StrTo) MustInt() int {
	i, _ := s.Int()
	return i
}
func (s StrTo) UInt32() (uint32, error){
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
