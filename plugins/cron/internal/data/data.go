package data

import (
	"encoding/json"
	"os"
)

type Schedule struct {
	Channel string
	Fields  string
	Command string
}

type ScheduleStore struct {
	list []*Schedule
}

func LoadScheduleStore(path string) (*ScheduleStore, error) {
	if _, err := os.Stat(path); err != nil {
		// ファイルが存在しない
		return &ScheduleStore{}, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var list []*Schedule
	err = json.NewDecoder(file).Decode(&list)
	if err != nil {
		return nil, err
	}

	return &ScheduleStore{
		list: list,
	}, nil
}

func SaveScheduleStore(path string, store *ScheduleStore) error {
	if store == nil {
		panic("nil store.")
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(store.list)
}

func (s *ScheduleStore) Add(schedule *Schedule) {
	if schedule == nil {
		panic("nil schedule")
	}

	s.list = append(s.list, schedule)
}

func (s *ScheduleStore) Remove(i int) bool {
	if i < 0 || i >= len(s.list) {
		return false
	}

	s.list = append(s.list[:i], s.list[i+1:]...)

	return true
}

func (s *ScheduleStore) List() []*Schedule {
	list := make([]*Schedule, len(s.list))

	copy(list, s.list)

	return list
}
