package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Function to return duplicate elements in []string
func duplicateInArray(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	// if result == nil {
	//	return false
	// }
	return result // if result == s [] string then there are no duplicates
}

func main() {
	/*
		Log file format:
		customer_id1 : item_id1 , item_id2 ..... , item_idn
		customer_id2 : item_id1 , item_id2 ..... , item_idn
		...... customer_idn : item_idn......
	*/
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	count := make(map[string]int)
	bdata := bufio.NewScanner(file)
	// bufferedData.Scan()
	// bufferedData.Scan() // Skipping ( for test purposes only)
	for bdata.Scan() {

		//to split string
		data := strings.Split(bdata.Text(), ":")
		//eater_id := data[0]
		food_ids := strings.Split(data[1], " , ")

		// Trim whitespace before and after
		for i := 0; i < len(food_ids); i++ {
			food_id_key := strings.TrimSpace(food_ids[i])
			food_id_arr := strings.Split(food_id_key, ",")
			//fmt.Println(food_id_arr)
			hasDup := duplicateInArray(food_id_arr)
			if len(hasDup) == len(food_id_arr) {
				for key, num := range food_id_arr {
					count[food_id_arr[key]] = count[num] + 1
				}
			} else {
				log.Fatalln("Duplicate items detected")
			}

		}

	}

	//fmt.Println(count)
	keys := make([]string, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return count[keys[i]] > count[keys[j]] })

	var top, second, third string

	for _, v := range keys {
		fmt.Println(v, count[v])
		if top == "" {
			top = v
		} else if top != "" && second == "" {
			second = v
		} else if top != "" && second != "" && third == "" {
			third = v
		}
	}

	fmt.Printf("Top: %v, Second: %v, Third: %v", top, second, third)

}
