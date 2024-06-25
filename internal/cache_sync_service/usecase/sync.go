package usecase

type SyncUseCase struct {
	locations
}

func NewSyncUseCase() *SyncUseCase {
	return &SyncUseCase{}
}

func (u *SyncUseCase) Execute() {
	// captura locations
	//
}
