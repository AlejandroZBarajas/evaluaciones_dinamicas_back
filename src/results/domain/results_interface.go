package resultsDomain

import resultsEntity "evaluaciones/src/results/domain/entity"

type ResultsInterface interface {
	CreateResult(result *resultsEntity.ResultsEntity) error

	GetAllResultsByExam(examid int32) ([]*resultsEntity.ResultsEntity, error)

	DeleteAllResultsByExam(examid int32) error
}
