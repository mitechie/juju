// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package common

import (
	"launchpad.net/juju-core/state"
	"launchpad.net/juju-core/state/api/params"
)

// PasswordChanger implements a common SetPasswords method for use by
// various facades.
type PasswordChanger struct {
	st           AgentEntityGetter
	getCanChange GetAuthFunc
}

type AgentEntityGetter interface {
	AgentEntity(tag string) (state.AgentEntity, error)
}

// NewPasswordChanger returns a new PasswordChanger. The GetAuthFunc will be
// used on each invocation of SetPasswords to determine current permissions.
func NewPasswordChanger(st AgentEntityGetter, getCanChange GetAuthFunc) *PasswordChanger {
	return &PasswordChanger{
		st:           st,
		getCanChange: getCanChange,
	}
}

// SetPasswords sets the given password for each supplied entity, if possible.
func (pc *PasswordChanger) SetPasswords(args params.PasswordChanges) (params.ErrorResults, error) {
	result := params.ErrorResults{
		Results: make([]params.ErrorResult, len(args.Changes)),
	}
	if len(args.Changes) == 0 {
		return result, nil
	}
	canChange, err := pc.getCanChange()
	if err != nil {
		return params.ErrorResults{}, err
	}
	for i, param := range args.Changes {
		if !canChange(param.Tag) {
			result.Results[i].Error = ServerError(ErrPerm)
			continue
		}
		if err := pc.setPassword(param.Tag, param.Password); err != nil {
			result.Results[i].Error = ServerError(err)
		}
	}
	return result, nil
}

func (pc *PasswordChanger) setPassword(tag, password string) error {
	entity, err := pc.st.AgentEntity(tag)
	if err != nil {
		return err
	}
	// We set the mongo password first on the grounds that
	// if it fails, the agent in question should still be able
	// to authenticate to another API server and ask it to change
	// its password.

	// TODO(rog) when the API is universal, check that the entity is a
	// machine with jobs that imply it needs access to the mongo state.
	if err := entity.SetMongoPassword(password); err != nil {
		return err
	}
	return entity.SetPassword(password)
}
