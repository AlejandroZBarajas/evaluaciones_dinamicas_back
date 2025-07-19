package resultsApplication

import (
	resultsDomain "evaluaciones/src/results/domain"
	resultsEntity "evaluaciones/src/results/domain/entity"
)

type GetAllResultsByExam struct {
	repo resultsDomain.ResultsInterface
}

func NewGetAllResultsByExam(repo resultsDomain.ResultsInterface) *GetAllResultsByExam {
	return &GetAllResultsByExam{repo: repo}
}

func (g *GetAllResultsByExam) Run(examID int32) ([]resultsEntity.ResultsEntity, error) {
	return g.repo.GetAllResultsByExam(examID)
}
