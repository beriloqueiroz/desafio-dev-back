package usecase

type SyncUseCase struct {
}

func NewSyncUseCase() *SyncUseCase {
	return &SyncUseCase{}
}

func (u *SyncUseCase) Execute() {
	// captura locations
	// captura mensagens
	// popula cache
}
