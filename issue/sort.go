package issue

type SortedIssues map[string]map[string]map[string][]Issue

func Sort(issues *Issues) SortedIssues {
	sorted := make(SortedIssues)
	for _, issue := range issues.Resources {
		if len(issue.AssignedTo.Name) != 0 {
			sorted.Add(issue.AssignedTo.Name, issue.Status.Name, issue.Project.Name, issue)
		}
	}

	return sorted
}

func (m SortedIssues) Add(user, status, project string, value Issue) {
	_, ok := m[user]
	if !ok {
		makeUserMap(m, user)
	}

	_, ok = m[user][status]
	if !ok {
		makeStatusMap(m[user], status)
	}

	m[user][status][project] = append(m[user][status][project], value)
}

func makeUserMap(m SortedIssues, path string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]map[string][]Issue)
		m[path] = mm
	}
}

func makeStatusMap(m map[string]map[string][]Issue, path string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string][]Issue)
		m[path] = mm
	}
}
