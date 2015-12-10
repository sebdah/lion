package models

type Projects struct {
	Projects []*Project `json:"projects"`
}

func NewProjects() *Projects {
	return &Projects{}
}

func (ps *Projects) GetBySlug(slug string) *Project {
	for _, p := range ps.Projects {
		if p.Slug == slug {
			return p
		}
	}

	return nil
}

type Project struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func NewProject() *Project {
	return &Project{}
}
