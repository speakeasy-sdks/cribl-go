// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GitSettings - a list of string objects
type GitSettings struct {
	AuthType                 *string     `json:"authType,omitempty"`
	AutoAction               *string     `json:"autoAction,omitempty"`
	AutoActionMessage        *string     `json:"autoActionMessage,omitempty"`
	AutoActionSchedule       *string     `json:"autoActionSchedule,omitempty"`
	Branch                   *string     `json:"branch,omitempty"`
	CommitDeploySingleAction *bool       `json:"commitDeploySingleAction,omitempty"`
	DefaultCommitMessage     *string     `json:"defaultCommitMessage,omitempty"`
	GitOps                   *GitOpsType `json:"gitOps,omitempty"`
	Password                 *string     `json:"password,omitempty"`
	Remote                   *string     `json:"remote,omitempty"`
	SSHKey                   *string     `json:"sshKey,omitempty"`
	StrictHostKeyChecking    *bool       `json:"strictHostKeyChecking,omitempty"`
	Timeout                  *int64      `json:"timeout,omitempty"`
	User                     *string     `json:"user,omitempty"`
}

func (o *GitSettings) GetAuthType() *string {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *GitSettings) GetAutoAction() *string {
	if o == nil {
		return nil
	}
	return o.AutoAction
}

func (o *GitSettings) GetAutoActionMessage() *string {
	if o == nil {
		return nil
	}
	return o.AutoActionMessage
}

func (o *GitSettings) GetAutoActionSchedule() *string {
	if o == nil {
		return nil
	}
	return o.AutoActionSchedule
}

func (o *GitSettings) GetBranch() *string {
	if o == nil {
		return nil
	}
	return o.Branch
}

func (o *GitSettings) GetCommitDeploySingleAction() *bool {
	if o == nil {
		return nil
	}
	return o.CommitDeploySingleAction
}

func (o *GitSettings) GetDefaultCommitMessage() *string {
	if o == nil {
		return nil
	}
	return o.DefaultCommitMessage
}

func (o *GitSettings) GetGitOps() *GitOpsType {
	if o == nil {
		return nil
	}
	return o.GitOps
}

func (o *GitSettings) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *GitSettings) GetRemote() *string {
	if o == nil {
		return nil
	}
	return o.Remote
}

func (o *GitSettings) GetSSHKey() *string {
	if o == nil {
		return nil
	}
	return o.SSHKey
}

func (o *GitSettings) GetStrictHostKeyChecking() *bool {
	if o == nil {
		return nil
	}
	return o.StrictHostKeyChecking
}

func (o *GitSettings) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *GitSettings) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
