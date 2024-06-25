package usecase

type GetMsgsUseCase struct {
}

func (u *GetMsgsUseCase) Execute(locations []string) (map[string]string, error) {
	//TODO implement me
	res := make(map[string]string)
	for _, loc := range locations {
		res[loc] = loc + " teste calor quintura e mormaço"
	}
	return res, nil
}
