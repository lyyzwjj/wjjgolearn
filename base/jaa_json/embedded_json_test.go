package jaa_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/9/10 1:58 上午
 * @description
 */
var jsonStr = `{
		"basic_info":{
			"name":"wjj",
			"age":25
		},
		"job_info":{
			"skills":[
				"Java",
				"Go",
				"C"
			]
		}
	}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e) // str ->对象
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil { // 对象 -> str
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}
