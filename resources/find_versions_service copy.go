package resources

func GetVersions(project string, language string) ([]string, error) {
	return findVersions(project, language)
}
