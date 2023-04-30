package main
import (
	"fmt"
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
type NotePublishType string
type CommentPermissionType string
type NotePermissionRole string

const (
	TeamVisibilityPublic  TeamVisibilityType = "public"
	TeamVisibilityPrivate TeamVisibilityType = "private"

	NotePublishTypeEdit   NotePublishType = "edit"
	NotePublishTypeView   NotePublishType = "view"
	NotePublishTypeSlide  NotePublishType = "slide"
	NotePublishTypeBook   NotePublishType = "book"

	CommentPermissionTypeDisabled       CommentPermissionType = "disabled"
	CommentPermissionTypeForbidden      CommentPermissionType = "forbidden"
	CommentPermissionTypeOwners         CommentPermissionType = "owners"
	CommentPermissionTypeSignedInUsers  CommentPermissionType = "signed_in_users"
	CommentPermissionTypeEveryone       CommentPermissionType = "everyone"

	NotePermissionRoleOwner    NotePermissionRole = "owner"
	NotePermissionRoleSignedIn NotePermissionRole = "signed_in"
	NotePermissionRoleGuest    NotePermissionRole = "guest"
)

type CreateNoteOptions struct {
	Title           string              `json:"title,omitempty"`
	Content         string              `json:"content,omitempty"`
	ReadPermission  NotePermissionRole  `json:"readPermission,omitempty"`
	WritePermission NotePermissionRole  `json:"writePermission,omitempty"`
	CommentPermission CommentPermissionType `json:"commentPermission,omitempty"`
	Permalink       string              `json:"permalink,omitempty"`
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
			SetTimeout(10 * time.Second).
			SetHeader("Authorization", "Bearer "+accessToken).
			SetHeader("Content-Type", "application/json"),
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
