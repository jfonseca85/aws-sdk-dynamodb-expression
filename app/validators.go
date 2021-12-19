package app

func CreateAppParams() []*Param {

	return []*Param{
		{
			Name:     argId,
			Type:     "string",
			Required: true,
		},
		{
			Name:     argDocument,
			Type:     "yaml",
			Required: true,
		},
	}
}
