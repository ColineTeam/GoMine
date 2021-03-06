package permissions

import (
	"gomine/interfaces"
	"errors"
)

type PermissionManager struct {
	server interfaces.IServer

	defaultGroup interfaces.IPermissionGroup
	permissions map[string]interfaces.IPermission
	groups map[string]interfaces.IPermissionGroup
}

/**
 * Creates a new PermissionManager struct.
 */
func NewPermissionManager(server interfaces.IServer) *PermissionManager {
	var manager = &PermissionManager{server, NewPermissionGroup("visitor", LevelVisitor), make(map[string]interfaces.IPermission), make(map[string]interfaces.IPermissionGroup)}

	manager.AddGroup(NewPermissionGroup("visitor", LevelVisitor))
	var visitorGroup, _ = manager.GetGroup("visitor")

	manager.AddGroup(NewPermissionGroup("member", LevelMember))
	var memberGroup, _ = manager.GetGroup("member")
	memberGroup.InheritGroup(visitorGroup)

	manager.AddGroup(NewPermissionGroup("operator", LevelOperator))
	var operatorGroup, _ = manager.GetGroup("operator")
	operatorGroup.InheritGroup(memberGroup)

	manager.AddGroup(NewPermissionGroup("custom", LevelCustom))

	return manager
}

/**
 * Returns the main server.
 */
func (manager *PermissionManager) GetServer() interfaces.IServer {
	return manager.server
}

/**
 * Returns the default group.
 */
func (manager *PermissionManager) GetDefaultGroup() interfaces.IPermissionGroup {
	return manager.defaultGroup
}

/**
 * Sets the default group.
 */
func (manager *PermissionManager) SetDefaultGroup(group interfaces.IPermissionGroup) {
	manager.defaultGroup = group
}

/**
 * Adds a new permission group.
 * Returns true if a group with the same name was overwritten.
 */
func (manager *PermissionManager) AddGroup(group interfaces.IPermissionGroup) bool {
	var overwritten = manager.GroupExists(group.GetName())

	manager.groups[group.GetName()] = group

	return overwritten
}

/**
 * Returns a group and an error if the group doesn't exist.
 */
func (manager *PermissionManager) GetGroup(name string) (interfaces.IPermissionGroup, error) {
	var group interfaces.IPermissionGroup
	if !manager.GroupExists(name) {
		return group, errors.New("group does not exist")
	}
	return manager.groups[name], nil
}

/**
 * Checks if a group with the given name exists.
 */
func (manager *PermissionManager) GroupExists(name string) bool {
	var _, ok = manager.groups[name]
	return ok
}

/**
 * Removes a permission group.
 * Returns true if the removal was successful.
 */
func (manager *PermissionManager) RemoveGroup(name string) bool {
	var groupExists = manager.GroupExists(name)

	delete(manager.groups, name)

	return groupExists
}

/**
 * Returns a permission with the given name if it exists, otherwise gives an error.
 */
func (manager *PermissionManager) GetPermission(name string) (interfaces.IPermission, error) {
	var perm *Permission
	if !manager.IsPermissionRegistered(name) {
		return perm, errors.New("tried to get an unregistered permission")
	}
	return manager.permissions[name], nil
}

/**
 * Checks if the given permission is registered.
 */
func (manager *PermissionManager) IsPermissionRegistered(name string) bool {
	var _, ok = manager.permissions[name]

	return ok
}

/**
 * Registers a new permission.
 * Returns true if a permission was overwritten.
 */
func (manager *PermissionManager) RegisterPermission(permission interfaces.IPermission) bool {
	var isRegistered = manager.IsPermissionRegistered(permission.GetName())

	manager.permissions[permission.GetName()] = permission

	return isRegistered
}