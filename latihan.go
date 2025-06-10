import (
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
func main() {
	tests := []string{
		"123", "-123", "3.14", "+3.14", "2e10", " 3 ", "abc", "1a", "e3", ".", "3.", ".5", "3e", "3e+2", "-.9",
	}

	for _, s := range tests {
		fmt.Printf("'%s' is number? %v\n", s, isNumber(s))
	}
}
