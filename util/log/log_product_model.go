package smitlog

type categoryNames string

func (CategoryNames categoryNames) isProduct() categoryNames {
	return CategoryNames
}

type LogCategoryNames interface {
	isProduct() categoryNames
}

const (
	Arcade   = categoryNames("Arcade")
	Action   = categoryNames("Action")
	Strategy = categoryNames("Strategy")
	ShowAll  = categoryNames("ShowAll")
	All      = categoryNames("All")
)
