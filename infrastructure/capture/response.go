package capture

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 17:31
**/

type Response struct {
	*WareError
	Rlt any `json:"rlt"`
}
