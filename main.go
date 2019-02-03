package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Вход: ")
	text, _ := reader.ReadString('\n')
	tmp := strings.Split(strings.Replace(text, "\n", "", 1), " ")
	nums := make([]int, 4)
	for i := range nums {
		nums[i], _ = strconv.Atoi(tmp[i])
	}

	if willTheyMeet(float64(nums[0]), float64(nums[1]), float64(nums[2]), float64(nums[3])) {
		fmt.Println("Результат: YES")
	}else{
		fmt.Println("Результат: NO")
	}
}


func willTheyMeet(x1, v1, x2, v2 float64) bool {
	var t float64
	x := (x1*v2 - x2*v1)/(v2-v1)
	t = 1/v1 * x - x1/v1

	return t > 0 && (math.Mod(t, float64(1))) == 0
}