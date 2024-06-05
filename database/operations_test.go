package database

import (
	"errors"
	"testing"

	// does a cd ..
	_ "github.com/Jatin020403/BasaltDB/test_init"
	"github.com/stretchr/testify/assert"
)

func TestInsertOne(t *testing.T) {
	type args struct {
		partition string
		key       string
		value     string
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "should insert value",
			args: args{"test11", "k1", "v1"},
			want: nil,
		},
		{
			name: "should throw error (no partition)",
			args: args{"", "k1", "v1"},
			want: errors.New("InsertOne : invalid partition"),
		},
		{
			name: "should throw error (no key)",
			args: args{"test11", "", "v1"},
			want: errors.New("InsertOne : invalid key"),
		},
		{
			name: "should throw error (no value)",
			args: args{"test11", "k1", ""},
			want: errors.New("InsertOne : invalid value"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreatePartition(tt.args.partition)
			got := InsertOne(tt.args.partition, tt.args.key, tt.args.value)
			assert.Equal(t, tt.want, got)
			DeletePartition(tt.args.partition)
		})
	}
}

func TestDeleteOne(t *testing.T) {
	type args struct {
		partition string
		key       string
		value     string
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "should delete value",
			args: args{"test11", "k1", "v1"},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreatePartition(tt.args.partition)
			
			err := InsertOne(tt.args.partition, tt.args.key, tt.args.value)
			assert.Equal(t, nil, err)

			got := DeleteOne(tt.args.partition, tt.args.key)
			assert.Equal(t, tt.want, got)
			DeletePartition(tt.args.partition)
		})
	}
}

func TestGetOne(t *testing.T) {
	type args struct {
		partition string
		key       string
		value     string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should get value",
			args: args{"test11", "k1", "v1"},
			want: "v1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreatePartition(tt.args.partition)
			
			err := InsertOne(tt.args.partition, tt.args.key, tt.args.value)
			assert.Equal(t, nil, err)

			got,err := GetOne(tt.args.partition, tt.args.key)
			assert.Equal(t, nil, err)
			assert.Equal(t, tt.want, got)
			DeletePartition(tt.args.partition)
		})
	}
}
