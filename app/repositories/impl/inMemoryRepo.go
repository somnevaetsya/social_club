package repositories_impl

import (
	"social_club/app/models"
	"social_club/app/repositories"
	"strconv"
	"sync"
)

type InMemoryRepository struct {
	adjMatr map[uint][]*models.Node
	mutex   sync.RWMutex
}

func MakeInMemoryRepository() repositories.Repository {
	return &InMemoryRepository{adjMatr: make(map[uint][]*models.Node)}
}

func (repo *InMemoryRepository) contains(n1 *models.Node, n2 *models.Node) {
	flag := false
	if _, ok := repo.adjMatr[n1.Id]; ok {
		flag = false
		for _, item := range repo.adjMatr[n1.Id] {
			if item.Id == n2.Id {
				item.Weight++
				flag = true
			}
		}
		if flag == false {
			n2.Weight = 1
			repo.adjMatr[n1.Id] = append(repo.adjMatr[n1.Id], n2)
		}
	} else {
		n2.Weight = 1
		repo.adjMatr[n1.Id] = append(repo.adjMatr[n1.Id], n2)
	}
}

func (repo *InMemoryRepository) findStartPoint() uint {
	var max uint
	for n := range repo.adjMatr {
		if n > max {
			max = n
		}
	}
	return max
}

func (repo *InMemoryRepository) IsEmpty() (bool, error) {
	if len(repo.adjMatr) == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (repo *InMemoryRepository) Add(n1 *models.Node, n2 *models.Node) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.contains(n1, n2)
	repo.contains(n2, n1)
	return nil
}

func (repo *InMemoryRepository) GetInfo() (min, max uint, avg float32, err error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	used := make(map[uint]bool)
	startPoint := repo.findStartPoint()

	used[startPoint] = true

	var queue []uint
	queue = append(queue, startPoint)
	var countUsers uint
	var allWeights uint
	min = repo.adjMatr[startPoint][0].Weight
	for len(queue) > 0 {
		u := queue[0]     //front
		queue = queue[1:] //pop
		countUsers++
		for i := 0; i < len(repo.adjMatr[u]); i++ {
			v := repo.adjMatr[u][i]
			allWeights += v.Weight
			if v.Weight < min {
				min = v.Weight
			}
			if v.Weight > max {
				max = v.Weight
			}
			if !used[v.Id] {

				used[v.Id] = true
				queue = append(queue, v.Id)
			}
		}
	}
	avg = float32(allWeights) / (float32(countUsers) * 2)
	return
}

func (repo *InMemoryRepository) GetGraph() (models.Info, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	s := ""
	var messages []models.UserInfo
	for node, nodes := range repo.adjMatr {
		for _, item := range nodes {
			s = "{User: " + strconv.FormatUint(uint64(item.Id), 10) + "; Messages in dialog: " + strconv.FormatUint(uint64(item.Weight), 10) + "}; "
		}
		messages = append(messages, models.UserInfo{Id: node, Messages: s})
	}
	min, max, avg, err := repo.GetInfo()
	if err != nil {
		return models.Info{}, err
	}
	info := models.Info{
		Graph:    messages,
		MinValue: min,
		AvgValue: avg,
		MaxValue: max,
		IsEmpty:  false,
	}
	return info, nil
}
