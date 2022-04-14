package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/rodrigoengelberg/imersao-fullcycle-7/codepix/application/usecase"
	"github.com/rodrigoengelberg/imersao-fullcycle-7/codepix/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepostory := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepostory,
	}

	return transactionUseCase
}
