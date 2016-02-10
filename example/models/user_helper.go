//************************************************************************//
// API "congo": Model Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
	"github.com/jinzhu/gorm"
	"time"
)

// MediaType Retrieval Functions
// ListUser returns an array of view: default
func (m *UserDB) ListAppUser(ctx *goa.Context) []*app.User {
	now := time.Now()
	defer ctx.Info("ListUser", "duration", time.Since(now))
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing User", "error", err.Error())
		return objs
	}

	return objs
}

func (m *User) UserToAppUser() *app.User {
	user := &app.User{}
	user.Bio = m.Bio
	user.Country = m.Country
	user.Email = &m.Email
	user.Firstname = &m.Firstname
	user.State = m.State
	user.Lastname = &m.Lastname
	user.ID = &m.ID
	user.City = m.City

	return user
}

// OneAppUser returns an array of view: default
func (m *UserDB) OneUser(ctx *goa.Context, id int) (*app.User, error) {
	now := time.Now()
	var native User
	defer ctx.Info("OneUser", "duration", time.Since(now))
	err := m.Db.Scopes().Table(m.TableName()).Preload("Proposals").Preload("Reviews").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		ctx.Error("error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToAppUser()
	return &view, err

}

// MediaType Retrieval Functions
// ListUserLink returns an array of view: link
func (m *UserDB) ListAppUserLink(ctx *goa.Context) []*app.UserLink {
	now := time.Now()
	defer ctx.Info("ListUserLink", "duration", time.Since(now))
	var objs []*app.UserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&objs).Error

	//	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing User", "error", err.Error())
		return objs
	}

	return objs
}

func (m *User) UserToAppUserLink() *app.UserLink {
	user := &app.UserLink{}
	user.ID = &m.ID
	user.Email = &m.Email

	return user
}

// OneAppUserLink returns an array of view: link
func (m *UserDB) OneUserLink(ctx *goa.Context, id int) (*app.UserLink, error) {
	now := time.Now()
	var native User
	defer ctx.Info("OneUserLink", "duration", time.Since(now))
	err := m.Db.Scopes().Table(m.TableName()).Preload("Proposals").Preload("Reviews").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.RecordNotFound {
		ctx.Error("error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToAppUserLink()
	return &view, err

}