package iaa_pipe_filter

/**
 * @author  wjj
 * @date  2020/9/8 1:35 上午
 * @description
 */
// Request is the input of filter
type Request interface {
}

// Response is the output of filter
type Response interface {
}

// Filter interface is the definition of the data processing components
// Pipe-Filter structure
type Filter interface {
	Process(data Request) (Response, error)
}
