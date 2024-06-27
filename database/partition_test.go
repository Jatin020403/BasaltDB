package database

import (
	"testing"

	_ "github.com/Jatin020403/BasaltDB/test_init"
	"github.com/stretchr/testify/assert"
)

func TestCreatePartition(t *testing.T) {
	type args struct {
		partition string
		n         int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should create partition",
			args: args{"test_partition", 5},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// var DIRPATH = "./" + tt.args.partition + "/"
			// var FILEPATH = DIRPATH + tt.args.partition + ".gob"
			err := CreateTemplate(tt.args.partition, tt.args.n)
			assert.NoError(t, err)
			err = CreatePartition(tt.args.partition)
			// assert.Error(t, nil, err)
			assert.NoError(t, err)

			assert.DirExists(t, tt.args.partition)

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
			err := CreatePartition(tt.args.partition)
			assert.NoError(t, err)
			assert.DirExists(t, tt.args.partition)
			err = DeletePartition(tt.args.partition)
			assert.NoError(t, err)
			assert.NoDirExists(t, tt.args.partition)
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

// 			files, err := filepath.Glob(""+tt.args.partition+"/*.gob")

// 			assert.NoError(t, err)

// 			if len(files) == 0 {
// 				return []string{}, errors.New("GetAllPartitions : no partitions")
// 			}

// 			var pt []string

// 			for _, file := range files {
// 				pt = append(pt, strings.Split((strings.Split(file, "/")[1]), ".gob")[0])
// 			}

// 			got := !utils.CheckPathExists(DIRPATH)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }
