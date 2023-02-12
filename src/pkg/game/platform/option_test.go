package platform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFindOption(t *testing.T) {
	tests := []struct {
		name string
		want *FindOption
	}{
		{
			name: "new",
			want: &FindOption{
				SearchMode: SearchMode_All,
				Seek: Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: OrderOption{
					Order: Order_ID,
					Desc:  false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewFindOption(), tt.want)
		})
	}
}

func TestFindOption_SetSeek(t *testing.T) {
	type fields struct {
		SearchMode  SearchMode
		Seek        Seek
		Pagination  Pagination
		OrderOption OrderOption
	}
	type args struct {
		lastID LastID
		count  Count
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FindOption
	}{
		{
			name: "set ok",
			fields: fields{
				SearchMode: SearchMode_All,
				Seek: Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: OrderOption{
					Order: Order_ID,
					Desc:  false,
				},
			},
			args: args{
				lastID: 999,
				count:  888,
			},
			want: &FindOption{
				SearchMode: SearchMode_Seek,
				Seek: Seek{
					LastID: 999,
					Count:  888,
				},
				Pagination: Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: OrderOption{
					Order: Order_ID,
					Desc:  false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FindOption{
				SearchMode:  tt.fields.SearchMode,
				Seek:        tt.fields.Seek,
				Pagination:  tt.fields.Pagination,
				OrderOption: tt.fields.OrderOption,
			}
			got := f.SetSeek(tt.args.lastID, tt.args.count)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestFindOption_SetPagination(t *testing.T) {
	type fields struct {
		SearchMode  SearchMode
		Seek        Seek
		Pagination  Pagination
		OrderOption OrderOption
	}
	type args struct {
		limit  Limit
		offset Offset
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FindOption
	}{
		{
			name: "set ok",
			fields: fields{
				SearchMode: SearchMode_Pagination,
				Seek: Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: OrderOption{
					Order: Order_ID,
					Desc:  false,
				},
			},
			args: args{
				limit:  999,
				offset: 888,
			},
			want: &FindOption{
				SearchMode: SearchMode_Pagination,
				Seek: Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: Pagination{
					Limit:  999,
					Offset: 888,
				},
				OrderOption: OrderOption{
					Order: Order_ID,
					Desc:  false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FindOption{
				SearchMode:  tt.fields.SearchMode,
				Seek:        tt.fields.Seek,
				Pagination:  tt.fields.Pagination,
				OrderOption: tt.fields.OrderOption,
			}
			got := f.SetPagination(tt.args.limit, tt.args.offset)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestFindOption_SetOrder(t *testing.T) {
	type fields struct {
		SearchMode  SearchMode
		Seek        Seek
		Pagination  Pagination
		OrderOption OrderOption
	}
	type args struct {
		order Order
		desc  Desc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FindOption
	}{
		{
			name: "set ok",
			fields: fields{
				SearchMode: SearchMode_All,
				Seek: Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: OrderOption{
					Order: Order_ID,
					Desc:  false,
				},
			},
			args: args{
				order: Order_Name,
				desc:  true,
			},
			want: &FindOption{
				SearchMode: SearchMode_All,
				Seek: Seek{
					LastID: 0,
					Count:  30,
				},
				Pagination: Pagination{
					Limit:  30,
					Offset: 0,
				},
				OrderOption: OrderOption{
					Order: Order_Name,
					Desc:  true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FindOption{
				SearchMode:  tt.fields.SearchMode,
				Seek:        tt.fields.Seek,
				Pagination:  tt.fields.Pagination,
				OrderOption: tt.fields.OrderOption,
			}
			got := f.SetOrder(tt.args.order, tt.args.desc)
			assert.Equal(t, got, tt.want)
		})
	}
}
