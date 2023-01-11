package main

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"
)

const idColumn = 1

const filepath = "ueba.csv" // TODO: получить из конфига

type Repository struct {
	table [][]string // Т.к. table не редактируется после инициализации, Mutex не требуется
}

func NewRepository() (*Repository, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 50 // TODO: получить из конфига
	t, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	t = t[1:]

	// Т.к. бизнес-логика предполагает получение записей по колонке ID
	// таблица сортируется для более быстрого поиска.
	sort.Slice(t, func(i, j int) bool {
		a, _ := strconv.Atoi(t[i][1])
		b, _ := strconv.Atoi(t[j][1])
		return a < b
	})

	return &Repository{table: t}, nil
}

func (r Repository) FindItemById(id string) (item []string, found bool) {
	i := sort.Search(len(r.table), func(i int) bool {
		return r.table[i][idColumn] >= id
	})
	if i < len(r.table) && r.table[i][idColumn] == id {
		return r.table[i], true
	}
	return nil, false
}
