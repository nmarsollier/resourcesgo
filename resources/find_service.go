package resources

import (
	"sort"
	"strconv"
	"strings"
)

func GetLastResource(project string, language string, semver string) (*Resource, error) {
	version, err := getLastVersion(project, language, semver)
	if err != nil {
		return nil, err
	}

	return findBy(project, language, version)
}

func getLastVersion(project string, language string, semver string) (string, error) {
	versions, err := findVersions(project, language)
	if err != nil {
		return "", err
	}

	var valids []string
	for i := 0; i < len(versions); i++ {
		if isValid(versions[i], semver) {
			valids = append(valids, versions[i])
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(valids)))

	if len(valids) > 0 {
		return valids[0], nil
	}

	return "", err
}

func isValid(version string, semVer string) bool {
	if strings.HasSuffix(semVer, "+") {
		newSemVer := strings.ReplaceAll(semVer, "+", "")
		return strings.HasPrefix(version, newSemVer)
	}

	versionArray := strings.Split(version, ".")
	semVerArray := strings.Split(semVer, ".")

	if len(versionArray) != 3 || len(semVerArray) != 3 {
		return false
	}

	v1, err := strconv.Atoi(versionArray[0])
	if err != nil {
		return false
	}
	s1, err := strconv.Atoi(semVerArray[0])
	if err != nil {
		return false
	}
	if v1 > s1 {
		return false
	}

	v2, err := strconv.Atoi(versionArray[1])
	if err != nil {
		return false
	}
	s2, err := strconv.Atoi(semVerArray[1])
	if err != nil {
		return false
	}
	if v2 > s2 {
		return false
	}

	v3, err := strconv.Atoi(versionArray[2])
	if err != nil {
		return false
	}
	s3, err := strconv.Atoi(semVerArray[2])
	if err != nil {
		return false
	}
	if v3 > s3 {
		return false
	}

	return true
}
