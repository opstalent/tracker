package resource

type (
	Field struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
	}

	CustomField struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	Resources interface {
		Sort(field string)
	}
)
