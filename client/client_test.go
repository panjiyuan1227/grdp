// client_test.go
package client

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	c := NewClient("10.45.18.65:3389", "wxh", "wangxh88@Hj", TC_RDP, nil)
	err := c.Login()
	if err != nil {
		fmt.Println("Login:", err)
	}
	c.OnReady(func() {
		fmt.Println("ready")
	})
	time.Sleep(100 * time.Second)
}
