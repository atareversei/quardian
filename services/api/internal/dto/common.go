package dto

type ListMeta struct {
	CurrentPage        int `json:"current_page"`
	PerPage            int `json:"per_page"`
	LastPage           int `json:"last_page"`
	TotalNumberOfItems int `json:"total_number_of_items"`
}
