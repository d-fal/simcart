package logger_test

import (
	"simcart/pkg/logger"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {

	t.Run("logger test", func(t *testing.T) {

		lo := logger.NewPrototype(logger.WithPath("."), logger.WithDebug(true))

		lo.Add("name", "davood").Add("item1 ", "item 2").Level(zapcore.InfoLevel)

		lo.Commit("done")

		lo.Add("lastname", "Ahmadi")

		lo.Commit("new name")
	})

}
