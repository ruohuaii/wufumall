package builder

import (
	guid "github.com/google/uuid"
	"github.com/segmentio/ksuid"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/4/24
* Time: 11:51
**/

func DigitID() uint32 {
	return guid.New().ID()
}

func StringID() string {
	return KsUID()
}

func UUID() string {
	return guid.New().String()
}

func KsUID() string {
	return ksuid.New().String()
}
