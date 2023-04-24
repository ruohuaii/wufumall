package builder

import "testing"

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/4/24
* Time: 12:10
**/

func TestIDBuilder(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(UUID(), DigitID(), StringID(), KsUID())
	}
}
