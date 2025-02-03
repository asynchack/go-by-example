package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(intB)
	fmt.Println(string(intB))

	strB, _ := json.Marshal("good")
	fmt.Println(string(strB))

	// marshal slice and map
	sliB, _ := json.Marshal([]string{"apple", "orange"})
	fmt.Println(string(sliB))

	mapB, _ := json.Marshal(map[string]int{"aa": 1, "bb": 2})
	fmt.Println(string(mapB))

	// custom data type
	resp1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "orange"},
	}
	respB, _ := json.Marshal(resp1)
	fmt.Println(string(respB))

	resp2 := &response2{
		Page:   2,
		Fruits: []string{"good", "upup"},
	}
	resp2B, _ := json.Marshal(resp2)
	fmt.Println(string(resp2B))

	// 反序列化，json字节流转为  go数据类型
	byt := []byte(`{"num": 1, "strs": ["aa", "bb"]}`)

	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

	num := dat["num"].(float64) // 上面的1为float类型
	fmt.Println(num)
	tmp := dat["strs"].([]interface{})
	aaVal := tmp[0].(string)
	fmt.Println(aaVal)

	//
	str11 := `{"page": 1, "fruits": ["apple", "peach"]}`
	fmt.Printf("%T\n", str11)
	resp11 := response2{}
	if err := json.Unmarshal([]byte(str11), &resp11); err != nil {
		panic(err)
	}
	fmt.Println(resp11)
	fmt.Println(resp11.Fruits[0])

	//
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"k1": 1, "k2": 2}
	if err := enc.Encode(d); err != nil {
		panic(err)
	}

	// decode to os.Stdout
	dec := json.NewDecoder(strings.NewReader(str11))
	resp22 := response2{}
	dec.Decode(&resp22)
	fmt.Println(resp22)
}
