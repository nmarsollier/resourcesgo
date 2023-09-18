package projects

func Delete(id string) error {
	err := delete(id)
	if err != nil {
		return err
	}

	return nil
}
