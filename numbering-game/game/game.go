package game

// 숫자 입력
// 숫자 비교. 맞으면 끝, 틀리면 다시
// 틀리면 카운트 +1

type Game struct {
	Count  int
	Number int
}

func New() *Game{
	return &Game{
		0,
		20,
	}
}
