package platform

type SearchMode uint8

const (
	SearchMode_All SearchMode = iota
	SearchMode_Seek
	SearchMode_Pagination
)

type LastID = ID
type Count = int
type Seek struct {
	LastID LastID
	Count  Count
}

type Limit = int
type Offset = int

type Pagination struct {
	Limit  Limit
	Offset Offset
}

type Order uint8

const (
	Order_ID Order = iota
	Order_Name
)

type Desc = bool

type OrderOption struct {
	Order Order
	Desc  Desc
}

// プラットフォーム検索オプション
type FindOption struct {
	SearchMode  SearchMode
	Seek        Seek
	Pagination  Pagination
	OrderOption OrderOption
}

func NewFindOption() *FindOption {
	return &FindOption{
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
	}
}

func (f *FindOption) SetSeek(lastID LastID, count Count) *FindOption {
	f.SearchMode = SearchMode_Seek
	f.Seek = Seek{
		LastID: lastID,
		Count:  count,
	}
	return f
}

func (f *FindOption) SetPagination(limit Limit, offset Offset) *FindOption {
	f.SearchMode = SearchMode_Pagination
	f.Pagination = Pagination{
		Limit:  limit,
		Offset: offset,
	}
	return f
}

func (f *FindOption) SetOrder(order Order, desc Desc) *FindOption {
	f.OrderOption = OrderOption{
		Order: order,
		Desc:  desc,
	}
	return f
}
