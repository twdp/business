package business

type Response struct {
	Success bool `json:"success"`
	Err     error `json:"err"`
}

func (r *Response) Error(err error) error {
	r.Success = false
	r.Err = err
	return nil
}
