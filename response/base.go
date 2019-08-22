package response

type BaseResponse struct {
	Err error
}

func (b *BaseResponse) IsError() bool {
	return b.Err == nil
}