package repositories_impl

import (
	"github.com/jackc/pgx"
	"social_club/app/models"
	"social_club/app/repositories"
	"strconv"
)

type PostgresRepository struct {
	db      *pgx.ConnPool
	adjMatr map[uint][]*models.Node
}

type Result struct {
	FirstUser  int
	SecondUser int
	Weight     int
}

func MakePostgresRepository(db *pgx.ConnPool) repositories.Repository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) Add(n1 *models.Node, n2 *models.Node) error {
	_, err := repo.db.Exec("insert into messages (first_user,second_user,weight) values ($1, $2, $3) on conflict(first_user,second_user) do update set weight=messages.weight+1;", n1.Id, n2.Id, 1)
	if err != nil {
		return err
	}
	_, err = repo.db.Exec("insert into messages (first_user,second_user,weight) values ($1, $2, $3) on conflict(first_user,second_user) do update set weight=messages.weight+1;", n2.Id, n1.Id, 1)
	if err != nil {
		return err
	}
	return err
}

func (repo *PostgresRepository) findStartPoint() uint {
	var max uint
	for n := range repo.adjMatr {
		if n > max {
			max = n
		}
	}
	return max
}

func (repo *PostgresRepository) IsEmpty() (bool, error) {
	var rows int
	err := repo.db.QueryRow("select count(*) from messages;").Scan(&rows)
	if err != nil {
		return false, err
	} else if rows != 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (repo *PostgresRepository) GetInfo() (min, max uint, avg float32, err error) {
	var minMsg int
	err = repo.db.QueryRow("select min(weight) from messages;").Scan(&minMsg)
	if err != nil {
		return
	}
	min = uint(minMsg)
	var maxMsg int
	err = repo.db.QueryRow("select max(weight) from messages;").Scan(&maxMsg)
	if err != nil {
		return
	}
	max = uint(maxMsg)

	used := make(map[uint]bool)
	startPoint := repo.findStartPoint()

	used[startPoint] = true

	var queue []uint
	queue = append(queue, startPoint)
	var countUsers uint
	var allWeights uint
	for len(queue) > 0 {
		u := queue[0]     //front
		queue = queue[1:] //pop
		countUsers++
		for i := 0; i < len(repo.adjMatr[u]); i++ {
			v := repo.adjMatr[u][i]
			allWeights += v.Weight
			if !used[v.Id] {
				used[v.Id] = true
				queue = append(queue, v.Id)
			}
		}
	}
	avg = float32(allWeights) / (float32(countUsers) * 2)
	return
}
func (repo *PostgresRepository) GetGraph() (models.Info, error) {
	rows, err := repo.db.Query("select * from messages;")
	defer rows.Close()
	var data []Result
	for rows.Next() {
		res := Result{}
		err = rows.Scan(&res.FirstUser, &res.SecondUser, &res.Weight)
		if err != nil {
			return models.Info{}, err
		}
		data = append(data, res)
	}
	s := ""
	repo.adjMatr = make(map[uint][]*models.Node)
	var messages []models.UserInfo
	for _, user := range data {
		repo.adjMatr[uint(user.FirstUser)] = append(repo.adjMatr[uint(user.FirstUser)], &models.Node{Id: uint(user.SecondUser), Weight: uint(user.Weight)})
	}
	for node, nodes := range repo.adjMatr {
		s = ""
		for _, item := range nodes {
			s += "{User: " + strconv.FormatUint(uint64(item.Id), 10) + "; Messages in dialog: " + strconv.FormatUint(uint64(item.Weight), 10) + "}; "
		}
		messages = append(messages, models.UserInfo{Id: node, Messages: s})
	}
	min, max, avg, err := repo.GetInfo()
	if err != nil {
		return models.Info{}, err
	}
	return models.Info{Graph: messages, MaxValue: max, MinValue: min, AvgValue: avg, IsEmpty: false}, nil
}
