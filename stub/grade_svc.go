package stub

type Student struct{}

type IGradebook interface {
	GradesFor(s *Student) map[string]float32 // map[string]int, key is subject, value is score
}

var _ IGradebook = (*RealGradebook)(nil)

type RealGradebook struct {
}

func (gb *RealGradebook) GradesFor(s *Student) map[string]float32 {
	// real implements
	return nil
}

type GradeService struct {
	gradebook IGradebook
}

func NewGradeService(gb IGradebook) *GradeService {
	return &GradeService{
		gradebook: gb,
	}
}

func (s *GradeService) AverageGrades(student *Student) float32 {
	grades := s.gradebook.GradesFor(student)
	sum := float32(0.0)
	for _, score := range grades {
		sum += score
	}
	if len(grades) == 0 {
		return 0
	}
	return sum / float32(len(grades))
}
