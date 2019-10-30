package CO2Struct

type Project struct {
	ID					int64				`json:"id"`
	Description			string				`json:"description"`
	Name				string				`json:"name"`
	NameWithNamespace	string				`json:"name_with_namespace"`
	Path				string				`json:"path"`
	PathWithNamespace	string				`json:"path_with_namespace"`
	CreatedAt			string				`json:"created_at"`
	DefaultBranch		string				`json:"default_branch"`
	TagList				string				`json:"tag_list"`
	SSHURLToRepo		string				`json:"ssh_url_to_repo"`
	HTTPURLToRepo		string				`json:"http_url_to_repo"`
	WebURL				string				`json:"web_url"`
	ReadmeURL			string				`json:"readme_url"`
	AvatarURL			string				`json:"avatar_url"`
	StarCount			int64				`json:"star_count"`
	ForksCount			int64				`json:"forks_count"`
	LastActivityAt		string				`json:"last_activity_at"`
	Namespace			ProjectNamespace	`json:"namespace"`
}

type ProjectNamespace struct {
	ID					int64				`json:"id"`
	Name				string				`json:"name"`
	Path				string				`json:"path"`
	Kind				string				`json:"kind"`
	FullPath			string				`json:"full_path"`
	ParentID			int64				`json:"parent_id"`
	AvatarURL			string				`json:"avatar_url"`
	WebURL				string				`json:"web_url"`
}