package nonrepeatingstr

import "fmt"

func searchMaxlong(str string) int {

	strmap := make(map[rune]int)

	maxlength := 0
	start := 0

	for i, s := range []rune(str) {

		if v, ok := strmap[s]; ok && v >= start {

			start = v + 1
		}

		if (i - start + 1) > maxlength {
			maxlength = i - start + 1

		}

		strmap[s] = i

	}

	return maxlength

}

func main() {
	m := map[string]string{

		"a": "adada",
		"b": "dsds",
		"c": "dfsfs",
		"d": "hahah",
		"e": "dsfsfs",
		"f": "daqwrterrdfsfdsg",
		"g": "dddd",
		"h": "有没有重复呢",
		"m": "上海的自来水来自海上",
		"o": "一二三二一",
	}

	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)

		fmt.Println(searchMaxlong(v))
	}

	delete(m, "a")
	fmt.Println(m)
	m["f"] = "fasdfdsaf"
	fmt.Println("add:", m)

	if v, ok := m["d"]; ok {

		fmt.Println(v)
	} else {
		fmt.Println("没有这个key")
	}

	if _, ok := m["df"]; ok {

	} else {
		fmt.Println("没有这个key")
	}
}
