package api

type TeamVisibilityType string

const (
	Public  TeamVisibilityType = "public"
	Private TeamVisibilityType = "private"
)

type NotePublishType string

const (
	Edit  NotePublishType = "edit"
	View  NotePublishType = "view"
	Slide NotePublishType = "slide"
	Book  NotePublishType = "book"
)

type CommentPermissionType string

const (
	Disabled      CommentPermissionType = "disabled"
	Forbidden     CommentPermissionType = "forbidden"
	Owners        CommentPermissionType = "owners"
	SignedInUsers CommentPermissionType = "signed_in_users"
	Everyone      CommentPermissionType = "everyone"
)

type NotePermissionRole string

const (
	Owner    NotePermissionRole = "owner"
	SignedIn NotePermissionRole = "signed_in"
	Guest    NotePermissionRole = "guest"
)

type CreateNoteOptions struct {
	Title             string                `json:"title,omitempty"`
	Content           string                `json:"content,omitempty"`
	ReadPermission    NotePermissionRole    `json:"readPermission,omitempty"`
	WritePermission   NotePermissionRole    `json:"writePermission,omitempty"`
	CommentPermission CommentPermissionType `json:"commentPermission,omitempty"`
	Permalink         string                `json:"permalink,omitempty"`
}

type UpdateNoteOptions struct {
	Content         string             `json:"content,omitempty"`
	ReadPermission  NotePermissionRole `json:"readPermission,omitempty"`
	WritePermission NotePermissionRole `json:"writePermission,omitempty"`
	Permalink       string             `json:"permalink,omitempty"`
}

type Team struct {
	ID          string             `json:"id"`
	OwnerID     string             `json:"ownerId,omitempty"`
	Name        string             `json:"name"`
	Logo        string             `json:"logo"`
	Path        string             `json:"path"`
	Description string             `json:"description"`
	HardBreaks  bool               `json:"hardBreaks"`
	Visibility  TeamVisibilityType `json:"visibility"`
	CreatedAt   int64              `json:"createdAt"`
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email,omitempty"`
	Name     string `json:"name"`
	UserPath string `json:"userPath"`
	Photo    string `json:"photo"`
	Teams    []Team `json:"teams"`
}

type SimpleUserProfile struct {
	Name      string `json:"name"`
	UserPath  string `json:"userPath"`
	Photo     string `json:"photo"`
	Biography string `json:"biography,omitempty"`
}

type Note struct {
	ID             string            `json:"id"`
	Title          string            `json:"title"`
	Tags           []string          `json:"tags,omitempty"`
	LastChangedAt  int64             `json:"lastChangedAt,omitempty"`
	CreatedAt      int64             `json:"createdAt"`
	LastChangeUser SimpleUserProfile `json:"lastChangeUser,omitempty"`
	PublishType    NotePublishType   `json:"publishType"`
	PublishedAt    int64             `json:"publishedAt,omitempty"`
	UserPath       string            `json:"userPath,omitempty"`
	TeamPath       string            `json:"teamPath,omitempty"`
	Permalink      string            `json:"permalink,omitempty"`
	ShortID        string            `json:"shortId"`
	PublishLink    string            `json:"publishLink,omitempty"`

	ReadPermission  NotePermissionRole `json:"readPermission"`
	WritePermission NotePermissionRole `json:"writePermission"`
}

type SingleNote struct {
	Note
	Content string `json:"content,omitempty"`
}
