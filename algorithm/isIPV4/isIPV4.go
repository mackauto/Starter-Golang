package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func IsValidIPV4(ip string) bool {
	strs := strings.Split(ip, ".")
	if len(strs) != 4 {
		return false
	}
	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if num < 0 || num > math.MaxUint8 {
			return false
		}
	}
	return true
}

func main() {
	type t struct {
		ip  string
		ans bool
	}
	ts := []t{
		{ip: "127.0.0.1", ans: true},
		{ip: "127.0.0.0", ans: true},
		{ip: "127.0.0.256", ans: false},
		{ip: "127.0.0.-1", ans: false},
		{ip: "127.0.0", ans: false},
		{ip: "127.0.0.1.1", ans: false},
	}
	for _, v := range ts {
		ans := IsValidIPV4(v.ip)
		if ans != v.ans {
			fmt.Println(v.ip, "exp:", v.ans, "res:", ans)
			return
		}
	}
	fmt.Println("pass")
}
