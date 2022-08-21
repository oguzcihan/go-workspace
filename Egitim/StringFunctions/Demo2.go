package StringFunctions

import (
	"fmt"
	"strings"
)

func Demo2() {
	lastName := "Cihan"
	fmt.Println(strings.HasPrefix(lastName, "Ci")) //ifadede ön ek bulmak için kullanılır /true/false
	fmt.Println(strings.HasSuffix(lastName, "an")) //ifadede son ek bulmak için kullanılır /true/false

	harfler := []string{"foo", "bar", "baz"}
	result := strings.Join(harfler, ",")
	fmt.Println(result) //string dizileri birleştirir
	fmt.Println(strings.Replace(result, "foo", "oğuz", -1))

	result2 := strings.Split(result, ",")
	fmt.Println(result2)
	for _, s := range result2 {
		fmt.Println(s)
	}

	fmt.Println(strings.Repeat(result, 3))

}
