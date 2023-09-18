package projects

// Find busca un token en la db
func Find(id string) (*Project, error) {
	return findByID(id)
}
