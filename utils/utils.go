package utils

import "document-service/models"

func Response(error bool, message string, data interface{}) models.Response {
	var res models.Response

	res.SetError(error)
	res.SetMessage(message)
	res.SetData(data)

	return res
}
