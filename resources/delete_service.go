package resources

func Delete(proejct string, language string, semver string) {
	if resource, err := findBy(proejct, language, semver); err == nil {
		delete(resource.ID)
	}
}
