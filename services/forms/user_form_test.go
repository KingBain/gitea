// Copyright 2018 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package forms

import (
	"testing"

	"code.gitea.io/gitea/modules/setting"

	"github.com/stretchr/testify/assert"
)

func TestRegisterForm_IsDomainAllowed_Empty(t *testing.T) {
	_ = setting.Service

	setting.Service.EmailDomainWhitelist = setting.BuildEmailGlobs([]string{})

	form := RegisterForm{}

	assert.True(t, form.IsEmailDomainAllowed())
}

func TestRegisterForm_IsDomainAllowed_InvalidEmail(t *testing.T) {
	_ = setting.Service

	setting.Service.EmailDomainWhitelist = setting.BuildEmailGlobs([]string{"gitea.io"})

	tt := []struct {
		email string
	}{
		{"securitygieqqq"},
		{"hdudhdd"},
	}

	for _, v := range tt {
		form := RegisterForm{Email: v.email}

		assert.False(t, form.IsEmailDomainAllowed())
	}
}

func TestRegisterForm_IsDomainAllowed_WhitelistedEmail(t *testing.T) {
	_ = setting.Service

	setting.Service.EmailDomainWhitelist = setting.BuildEmailGlobs([]string{"gitea.io", "*.gc.ca"})

	tt := []struct {
		email string
		valid bool
	}{
		{"security@gitea.io", true},
		{"security@gITea.io", true},
		{"hdudhdd", false},
		{"seee@example.com", false},
		{"security@fishsauce.gc.ca", true},
	}

	for _, v := range tt {
		form := RegisterForm{Email: v.email}

		assert.Equal(t, v.valid, form.IsEmailDomainAllowed())
	}
}

func TestRegisterForm_IsDomainAllowed_BlocklistedEmail(t *testing.T) {
	_ = setting.Service

	setting.Service.EmailDomainWhitelist = setting.BuildEmailGlobs([]string{})
	setting.Service.EmailDomainBlocklist = setting.BuildEmailGlobs([]string{"gitea.io", "*.gov"})

	tt := []struct {
		email string
		valid bool
	}{
		{"security@gitea.io", false},
		{"security@gitea.example", true},
		{"hdudhdd", true},
		{"security@fishsauce.gov", false},
	}

	for _, v := range tt {
		form := RegisterForm{Email: v.email}

		assert.Equal(t, v.valid, form.IsEmailDomainAllowed())
	}
}
