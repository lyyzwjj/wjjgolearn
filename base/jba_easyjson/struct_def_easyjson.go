package jaa_json

/**
 * @author  wjj
 * @date  2020/9/10 1:55 上午
 * @description
 */

type BasicInfo struct {
	Name string
	Age  int
}
type JobInfo struct {
	Skills []string
}
type Employee struct {
	BasicInfo BasicInfo
	JobInfo   JobInfo
}
