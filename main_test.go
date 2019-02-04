package tdl_test

import (
	"os"
	"testing"

	"github.com/joematpal/go-tdl"
)

var Learner tdl.Student

func TestMain(m *testing.M) {

	Learner = tdl.ParseQuestions()

	exitCode := m.Run()
	os.Exit(exitCode)
}
