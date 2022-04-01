package models

type Password struct {
	Old_Password string `json:"old_password,omitempty"`
	New_Password string `json:"new_password,omitempty"`
}
