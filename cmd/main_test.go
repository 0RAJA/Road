package main

import (
	"fmt"
	"github.com/0RAJA/Road/internal/global"
	"testing"
)

func TestSetupSetting(t *testing.T) {
	err := SetupSetting()
	if err != nil {
		fmt.Println("Setup err:", err)
	}
	fmt.Println(global.AllSetting)
}
