package resultsApplication

import (
	resultsDomain "evaluaciones/src/results/domain"
	resultsEntity "evaluaciones/src/results/domain/entity"
)

type CreateResult struct {
	repo resultsDomain.ResultsInterface
}

func NewCreateResult(repo resultsDomain.ResultsInterface) *CreateResult {
	return &CreateResult{repo: repo}
}

func (c *CreateResult) Run(result *resultsEntity.ResultsEntity) error {
	return c.repo.CreateResult(result)
}
