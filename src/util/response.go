package util

type MessageResponse struct {
	Message interface{}  `json:"message"`
	Detail  interface{} `json:"detail,omitempty"`
}

func ErrorResponse (message string ,err error)MessageResponse{
	return MessageResponse{
		Message:message,
		Detail:err.Error(),
	}
}

func GetErrorResponse(err error)MessageResponse{
	return ErrorResponse("取得に失敗しました",err)
}

func CreateErrorResponse(err error)MessageResponse{
	return ErrorResponse("データに不備があります",err)
}

func DeleterrorResponse(err error)MessageResponse{
	return ErrorResponse("削除に失敗しました",err)
}

func DBCreateErrorResponse(err error)MessageResponse{
	return ErrorResponse("データベースエラー",err)
}

func NomalErrorResponse(err error)MessageResponse{
	return ErrorResponse("エラー",err)
}