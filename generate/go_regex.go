package generate

import (
	"fmt"
	"regexp"
	"strings"
)

// Example function
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// }
// Extracting: function name, function parameters (name + types) and function return type
var funcRegexp = regexp.MustCompile(`func (.*?)\((.*?)\) (.*?) {`)

type GoSignature struct {
	lowerName  string
	upperName  string
	parameters []string
	out        string
}

func GetGoSignature(code string) GoSignature {
	gs := GoSignature{}
	matches := funcRegexp.FindStringSubmatch(code)
	if len(matches) == 4 {
		gs.lowerName = matches[1]

		if len(matches[1]) > 1 {
			gs.upperName = strings.ToUpper(string(matches[1][0])) + string(matches[1][1:])
		}

		gs.parameters = strings.Split(matches[2], ", ")
		gs.out = matches[3]
	} else {
		fmt.Println("len not equal to 4")
	}
	return gs
}
