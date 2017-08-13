package user

import (
	osuser "os/user"
)

// User is the structure of a user
type User struct {
	UID       string
	Gid       string
	Username  string
	Groupname string
	Groupids  []string
	HomeDir   string
}

// GetCurrentUser retrieves the current
func GetCurrentUser() (User, error) {
	var user User
	var err error

	theUser, err := osuser.Current()
	if err != nil {
		return user, err
	}
	theGroup, err := osuser.LookupGroupId(theUser.Gid)
	if err != nil {
		return user, err
	}
	groupids, err := theUser.GroupIds()
	if err != nil {
		return user, err
	}
	homeDir := theUser.HomeDir
	user = User{
		UID:       theUser.Uid,
		Gid:       theUser.Gid,
		Username:  theUser.Username,
		Groupname: theGroup.Name,
		Groupids:  groupids,
		HomeDir:   homeDir,
	}
	return user, err
}
