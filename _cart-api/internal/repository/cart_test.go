package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartRepository_Add(t *testing.T) {
	type args struct {
		userID string
		item   Item
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				userID: "user-id",
				item: Item{
					ID: "some ID",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewCartRepository()
			if err := r.Add(tt.args.userID, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			carts, err := r.Get(tt.args.userID)
			if err != nil {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, 1, len(carts))
		})
	}
}
