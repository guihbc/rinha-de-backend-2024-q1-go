package response

import "encoding/json"

func GetBytes(response any) []byte {
	b, _ := json.Marshal(&response)
	return b
}
