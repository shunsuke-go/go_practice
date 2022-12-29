package outputport

// output (プレゼンター) のinterface
type InitialPresenter struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type IInitailPresenter interface {
	ToInitialPresenter() (InitialPresenter)	
}