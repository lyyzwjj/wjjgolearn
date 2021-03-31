package lba_optimization

/**
 * @author  wjj
 * @date  2020/9/13 7:00 下午
 * @description
 */

type Request struct {
	TransactionID string `json:"transaction_id"`
	PayLoad       []int  `json:"payload"`
}

type Response struct {
	TransactionID string `json:"transaction_id"`
	Expression    string `json:"exp"`
}
