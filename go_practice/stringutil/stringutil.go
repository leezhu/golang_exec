//Package stringutil contains string operation
package stringutil

var a = 1

const b = 1

//Reverse function is reverse string
func Reverse(str string) string {
	r := []int32(str)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
