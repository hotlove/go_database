package test

import (
	"go_database/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	logger.Info("测试日子打印")
	t.Logf("cehi")
}
