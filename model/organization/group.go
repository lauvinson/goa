// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import "goa/model"

// Group 组
type Group struct {
	model.BaseModel
	Name string `gorm:"not null;comment:'组名称'"`
	Code string `gorm:"not null;comment:'组编码'"`
}

// GroupUserMapping 用户组关联
type GroupUserMapping struct {
	model.BaseModel
	GroupID uint `gorm:"not null;comment:'组ID'"`
	UserID  uint `gorm:"not null;comment:'用户ID'"`
}

// GroupRoleMapping 组权限关联
type GroupRoleMapping struct {
	model.BaseModel
	GroupID uint `gorm:"not null;comment:'组ID'"`
	RoleID  uint `gorm:"not null;comment:'角色ID'"`
}

// GroupModuleMapping 组模块关联
type GroupModuleMapping struct {
	model.BaseModel
	GroupID  uint `gorm:"not null;comment:'组ID'"`
	ModuleID uint `gorm:"not null;comment:'模块ID'"`
}