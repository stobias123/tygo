package simple

import "github.com/google/uuid"

// Comments are kept :)
type ComplexType map[string]map[uint16]*uint32

type UserRole = string

const (
	UserRoleDefault UserRole = "viewer"
	UserRoleEditor  UserRole = "editor" // Line comments are also kept
)

type UserEntry struct {
	// Instead of specifying `tstype` we could also declare the typing
	// for uuid.NullUUID in the config file.
	ID uuid.NullUUID `json:"id" tstype:"string | null"`

	Preferences map[string]struct {
		Foo uint32 `json:"foo"`
		// An unknown type without a `tstype` tag or mapping in the config file
		// becomes `any`
		Bar uuid.UUID `json:"bar"`
	} `json:"prefs"`

	MaybeFieldWithStar *string  `json:"address"`
	Nickname           string   `json:"nickname,omitempty"`
	Role               UserRole `json:"role"`

	Complex    ComplexType `json:"complex"`
	unexported bool        // Unexported fields won't be in the output
}

type ListUsersResponse struct {
	Users []UserEntry `json:"users"`
}

type DumbExample struct {
	PrimaryUser UserEntry   `jsonapi:"attr,user"`
	Users       []UserEntry `jsonapi:"relation,users"`
}
