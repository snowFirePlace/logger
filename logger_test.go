package logger

import (
	"errors"
	"testing"
)

func Test(t *testing.T) {

	log := New()

	// fLog, _ := os.OpenFile("./name.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// log.SetOutput(fLog)

	// log.SetLevel("Debug")
	log.Debug("Start with debug ")
	log.Info("asfasfas")
	log.Info("asfasfas\r\nsfasfas\r\nsfafsawfasfas\r\nsdfaw")
	log.Info(4824, 342112)
	log.Info(errors.New("sfasfas"))
	log.Info("sfas", 123332)
	log.Info()
	log.Info("sfas", "123332")
	log.Info()
	log.Debug("123412")
	err := errors.New("sfasfasf")
	log.Error(err)

}
