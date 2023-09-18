package projects

// Create crea un nuevo token y lo almacena en la db
func Create(name string) (*Project, error) {
	project, err := insert(name)
	
	if err != nil {
		return nil, err
	}

	return project, nil
}
