package business

import "errors"

type Response struct {
	Success bool `json:"success"`
	//Err     error  `json:"err"`
	Err string `json:"error"`
}

func (r *Response) Error(err error) error {
	r.Success = false
	r.Err = err.Error()
	//r.Errr =
	return nil
}

func CheckError(r *Response, err error) {
	if err != nil {
		panic(err)
	} else if r.Err != "" {
		panic(errors.New(r.Err))
	}

}
