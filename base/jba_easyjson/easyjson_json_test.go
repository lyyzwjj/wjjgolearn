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
	e := Employee{}
	_ = e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)
	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(v))
	}
}
func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err = json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}
func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Error(err)
		}
		if _, err = e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}
