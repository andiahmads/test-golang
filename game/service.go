package game

type Service interface {
	InserData(input InputDataGame) (DataGame, error)
	ShowDataGame() ([]DataGame, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) InserData(input InputDataGame) (DataGame, error) {
	data := DataGame{}

	data.Mdate = input.Mdate
	data.Stadium = input.Stadium
	data.Team1 = input.Team1
	data.Team2 = input.Team2

	newData, err := s.repository.Insert(data)
	if err != nil {
		return newData, nil
	}

	return newData, nil

}

func (s *service) ShowDataGame() ([]DataGame, error) {
	return s.repository.ShowDataGame()
}
