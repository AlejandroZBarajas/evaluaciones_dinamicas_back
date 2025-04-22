package studentExamApplication

import studentExamDomain "evaluaciones/src/student_exam/domain"

type DeleteStudentExam struct {
	repo studentExamDomain.StudentExamInterface
}

func NewDeleteStudentExam(repo studentExamDomain.StudentExamInterface) *DeleteStudentExam {
	return &DeleteStudentExam{repo: repo}
}

func (d *DeleteStudentExam) Run(studentExamID int32) error {
	return d.repo.DeleteStudentExam(studentExamID)
}
