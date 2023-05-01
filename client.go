package client

import (
	"errors"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

const defaultHackmdAPIEndpointURL = "https://api.hackmd.io/v1"

type APIClient struct {
	accessToken          string
	hackmdAPIEndpointURL string
	client               *req.Client
}

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

func NewAPIClient(accessToken string, options ...Option) *APIClient {
	if accessToken == "" {
		panic("Missing access token when creating HackMD client")
	}

	client := &APIClient{
		accessToken:          accessToken,
		hackmdAPIEndpointURL: defaultHackmdAPIEndpointURL,
		client: req.C().
			SetBaseURL(defaultHackmdAPIEndpointURL).
			SetTimeout(10*time.Second).
			SetCommonHeader("Authorization", "Bearer "+accessToken).
			SetCommonHeader("Content-Type", "application/json"),
	}

	for _, opt := range options {
		opt(client)
	}

	return client
}

type Option func(*APIClient)

func WithAPIEndpointURL(url string) Option {
	return func(c *APIClient) {
		c.hackmdAPIEndpointURL = url
		c.client.SetBaseURL(url)
	}
}

// get /me
func (c *APIClient) GetMe() (*User, error) {
	var user User

	resp, err := c.client.
		R().SetSuccessResult(&user).Get(c.hackmdAPIEndpointURL + "/me")

	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get /me")
	}

	return &user, nil
}

func (c *APIClient) GetHistory() (*[]Note, error) {
	var notes []Note

	resp, err := c.client.
		R().SetSuccessResult(&notes).Get(c.hackmdAPIEndpointURL + "/history")

	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get /history")
	}

	return &notes, nil
}

func (c *APIClient) GetNoteList() (*[]Note, error) {
	var notes []Note

	resp, err := c.client.
		R().SetSuccessResult(&notes).Get(c.hackmdAPIEndpointURL + "/notes")

	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get /notes")
	}

	return &notes, nil
}

func (c *APIClient) GetNote(noteID string) (*SingleNote, error) {
	var note SingleNote

	resp, err := c.client.
		R().SetSuccessResult(&note).Get(c.hackmdAPIEndpointURL + "/notes/" + noteID)

	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get /notes/" + noteID)
	}

	return &note, nil
}

func (c *APIClient) CreateNote(options *CreateNoteOptions) (*SingleNote, error) {
	var note SingleNote

	resp, err := c.client.
		R().
		SetBody(options).
		SetSuccessResult(&note).
		Post(c.hackmdAPIEndpointURL + "/notes")

	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusCreated {
		return nil, errors.New("Failed to create note")
	}

	return &note, nil
}

// Update note api return 202 Accepted with empty body
func (c *APIClient) UpdateNoteContent(noteID string, content string) error {
	resp, err := c.client.
		R().
		SetBody(map[string]string{"content": content}).
		Put(c.hackmdAPIEndpointURL + "/notes/" + noteID)

	if err != nil {
		return err
	}

	if resp.Response.StatusCode != http.StatusAccepted {
		return errors.New("Failed to update note content")
	}

	return nil
}

func (c *APIClient) UpdateNote(noteID string, options *UpdateNoteOptions) error {
	resp, err := c.client.
		R().
		SetBody(options).
		Patch(c.hackmdAPIEndpointURL + "/notes/" + noteID)

	if err != nil {
		return err
	}

	if resp.Response.StatusCode != http.StatusAccepted {
		return errors.New("Failed to update note")
	}

	return nil
}

func (c *APIClient) DeleteNote(noteID string) error {
	resp, err := c.client.
		R().
		Delete(c.hackmdAPIEndpointURL + "/notes/" + noteID)

	if err != nil {
		return err
	}

	// 204 No Content
	if resp.Response.StatusCode != http.StatusNoContent {
		return errors.New("Failed to delete note")
	}

	return nil
}

func (c *APIClient) GetTeams() ([]Team, error) {
	var teams []Team

	resp, err := c.client.
		R().SetSuccessResult(&teams).Get(c.hackmdAPIEndpointURL + "/teams")

	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get /teams")
	}

	return teams, nil
}

func (c *APIClient) GetTeamNotes(teamPath string) ([]Note, error) {
	var notes []Note

	resp, err := c.client.
		R().SetSuccessResult(&notes).Get(c.hackmdAPIEndpointURL + "/teams/" + teamPath + "/notes")

	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get /teams/" + teamPath + "/notes")
	}

	return notes, nil
}

func (c *APIClient) CreateTeamNote(teamPath string, options *CreateNoteOptions) (*SingleNote, error) {
	var note SingleNote

	resp, err := c.client.
		R().
		SetBody(options).
		SetSuccessResult(&note).
		Post(c.hackmdAPIEndpointURL + "/teams/" + teamPath + "/notes")

	if err != nil {
		return nil, err
	}

	if resp.Response.StatusCode != http.StatusCreated {
		return nil, errors.New("Failed to create team note")
	}

	return &note, nil
}

func (c *APIClient) UpdateTeamNoteContent(teamPath string, noteID string, content string) error {
	resp, err := c.client.
		R().
		SetBody(map[string]string{"content": content}).
		Put(c.hackmdAPIEndpointURL + "/teams/" + teamPath + "/notes/" + noteID)

	if err != nil {
		return err
	}

	if resp.Response.StatusCode != http.StatusAccepted {
		return errors.New("Failed to update team note content")
	}

	return nil
}

func (c *APIClient) UpdateTeamNote(teamPath string, noteID string, options *UpdateNoteOptions) error {
	resp, err := c.client.
		R().
		SetBody(options).
		Patch(c.hackmdAPIEndpointURL + "/teams/" + teamPath + "/notes/" + noteID)

	if err != nil {
		return err
	}

	if resp.Response.StatusCode != http.StatusAccepted {
		return errors.New("Failed to update team note")
	}

	return nil
}

func (c *APIClient) DeleteTeamNote(teamPath string, noteID string) error {
	resp, err := c.client.
		R().
		Delete(c.hackmdAPIEndpointURL + "/teams/" + teamPath + "/notes/" + noteID)

	if err != nil {
		return err
	}

	if resp.Response.StatusCode != http.StatusNoContent {
		return errors.New("Failed to delete team note")
	}

	return nil
}
