package golib

import "strings"

func StringsEqual(s1, s2 string) bool {
	return strings.Compare(s1, s2) == 0
}

func StringsNotEqual(s1, s2 string) bool {
	return !StringsEqual(s1, s2)
}


	
