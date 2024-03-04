package tracker

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type BasicUsers []BasicUser

// User
// Basic user structure in Yandex.Tracker
type BasicUser struct {
	Self    string `json:"self"`
	ID      string `json:"id"`
	Display string `json:"display"`
}

// Id
// Get user id
func (u *BasicUser) Id() string {
	if u != nil {
		return u.ID
	}

	return ""
}

// Name
// Get username
func (u *BasicUser) Name() string {
	if u != nil {
		return u.Display
	}

	return ""
}

type Users []User

// User structure in Yandex.Tracker
// https://cloud.yandex.ru/en/docs/tracker/get-user-info
type User struct {
	// Address of the API resource with information about the user account
	Self string `json:"self"`

	// Unique ID of the user Tracker account
	UID int `json:"uid"`

	// User's login
	Login string `json:"login"`

	// Unique ID of the user Tracker account
	TrackerUid int `json:"trackerUid"`

	// Unique ID of the user account in the Yandex 360 for Business organization and Yandex ID
	PassportUid int `json:"passportUid"`

	// User unique ID in Yandex Cloud Organization
	CloudUid string `json:"cloudUid"`

	// User's first name
	FirstName string `json:"firstName"`

	// User's last name
	LastName string `json:"lastName"`

	// Displayed user name
	Display string `json:"display"`

	// User email address
	Email string `json:"email"`

	// Service parameter
	External bool `json:"external"`

	// Flag indicating whether the user has full access to Tracker:
	// true: Full access
	// false: Read-only access
	HasLicense bool `json:"hasLicense"`

	// User status in the organization:
	// true: User is dismissed
	// false: User is a current employee
	Dismissed bool `json:"dismissed"`

	// Service parameter
	UseNewFilters bool `json:"useNewFilters"`

	// Flag indicating whether user notifications are forced disabled:
	// true: Disabled
	// false: Enabled
	DisableNotifications bool `json:"disableNotifications"`

	// Date and time of the user's first authentication, in the YYYY-MM-DDThh:mm:ss.sss±hhmm format
	FirstLoginDate string `json:"firstLoginDate"`

	// Date and time of the user's last authentication, in the YYYY-MM-DDThh:mm:ss.sss±hhmm format
	LastLoginDate string `json:"lastLoginDate"`

	// Method of adding a user:
	// true: By sending an invitation by email
	// false: By other means
	WelcomeMailSent bool `json:"welcomeMailSent"`
}

func (t *trackerClient) Myself() (*User, *resty.Response, error) {
	req := t.NewRequest(resty.MethodPost, "/v2/myself", nil)
	result := new(User)
	resp, err := t.Do(req, result)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get user info: %w", err)
	}
	return result, resp, nil
}
