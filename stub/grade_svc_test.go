package stub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var _ IGradebook = (*StubGradebook)(nil)

type StubGradebook struct {
	mock.Mock
}

func (gb *StubGradebook) GradesFor(student *Student) map[string]float32 {
	args := gb.Called(student)
	return args.Get(0).(map[string]float32)
}

func Test_GradeService_AverageGrades(t *testing.T) {
	student := &Student{}
	var stubgb StubGradebook
	stubgb.On("GradesFor", student).Return(map[string]float32{
		"OOP": 8,
		"FP":  6,
		"DB":  10,
	})
	svc := NewGradeService(&stubgb)
	avg := svc.AverageGrades(student)
	assert.EqualValues(t, 8.0, avg)
}
