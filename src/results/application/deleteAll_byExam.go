package resultsApplication

import resultsDomain "evaluaciones/src/results/domain"

type DeleteAllResultsByExam struct {
	repo resultsDomain.ResultsInterface
}

func NewDeleteAllResultsByExam(repo resultsDomain.ResultsInterface) *DeleteAllResultsByExam {
	return &DeleteAllResultsByExam{repo: repo}
}

func (d *DeleteAllResultsByExam) Run(examID int32) error {
	return d.repo.DeleteAllResultsByExam(examID)
}
