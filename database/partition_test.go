package database

import (
	"testing"

	_ "github.com/Jatin020403/BasaltDB/test_init"
	"github.com/Jatin020403/BasaltDB/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreatePartition(t *testing.T) {
	type args struct {
		partition string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should create partition",
			args: args{"test_partition"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var DIRPATH = "./storage/" + tt.args.partition + ".gob"
			err := CreatePartition(tt.args.partition)
			// assert.Error(t, nil, err)
			assert.NoError(t, err)

			got := utils.CheckFileExists(DIRPATH)
			assert.Equal(t, tt.want, got)

			err = DeletePartition(tt.args.partition)
			assert.NoError(t, err)
		})
	}
}

func TestDeletePartition(t *testing.T) {
	type args struct {
		partition string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should delete partition",
			args: args{"test_partition"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreatePartition(tt.args.partition)
			var DIRPATH = "./storage/" + tt.args.partition + ".gob"
			err := DeletePartition(tt.args.partition)
			// assert.Error(t, nil, err)
			assert.NoError(t, err)
			got := utils.CheckFileExists(DIRPATH)
			assert.Equal(t, tt.want, got)
		})
	}
}

// func TestGetAllPartitions(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want []string
// 	}{
// 		{
// 			name: "should contain newly added partition partition",
// 			want: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			files, err := filepath.Glob("storage/*.gob")

// 			assert.NoError(t, err)

// 			if len(files) == 0 {
// 				return []string{}, errors.New("GetAllPartitions : no partitions")
// 			}

// 			var pt []string

// 			for _, file := range files {
// 				pt = append(pt, strings.Split((strings.Split(file, "/")[1]), ".gob")[0])
// 			}

// 			got := !utils.CheckFileExists(DIRPATH)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }
