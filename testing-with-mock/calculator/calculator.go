package calculator

type DBService interface {
	GetSum() int
}

func GetSumFromDB(db DBService) int {
	return db.GetSum()
}
