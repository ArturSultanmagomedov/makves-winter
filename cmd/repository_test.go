package main

import (
	"encoding/csv"
	"os"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestRepository_FindItemById(t *testing.T) {
	type fields struct {
		table [][]string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedItem  []string
		expectedFound bool
	}{
		{
			name:   "Ok",
			fields: fields{getTable()},
			args:   args{"1913"},
			expectedItem: []string{"1261", "1913", "S-1-5-21-3686381713-1037878038-1682765610-2925", "dev.makves.ru",
				"Зоя Хованска", "ИТО", "Специалист технической поддержки АРМ", "Z.Khovanska", "78", "0", "0", "0", "0", "0",
				"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "1", "0", "0", "0",
				"0", "0", "0", "0", "0", "5", "5", "2;8;10;11;19", "9;25;26;43;44", "1", "0"},
			expectedFound: true,
		},
		{
			name:          "Not exist",
			fields:        fields{getTable()},
			args:          args{"3000"},
			expectedItem:  nil,
			expectedFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				table: tt.fields.table,
			}
			gotItem, gotFound := r.FindItemById(tt.args.id)
			if !reflect.DeepEqual(gotItem, tt.expectedItem) {
				t.Errorf("FindItemById() gotItem = %v, expected %v", gotItem, tt.expectedItem)
			}
			if gotFound != tt.expectedFound {
				t.Errorf("FindItemById() gotFound = %v, expected %v", gotFound, tt.expectedFound)
			}
		})
	}
}

func getTable() [][]string {
	file, err := os.Open("ueba.csv")
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 50
	t, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	t = t[1:]

	sort.Slice(t, func(i, j int) bool {
		a, _ := strconv.Atoi(t[i][1])
		b, _ := strconv.Atoi(t[j][1])
		return a < b
	})
	return t
}
