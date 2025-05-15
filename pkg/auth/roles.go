package auth

import (
	"fmt"
	"sync"
)

// Role represents a user role in the system
type Role string

const (
	RoleAdmin    Role = "ADMIN"
	RoleMember   Role = "MEMBER"
	RoleCreator  Role = "CREATOR"
	RoleClaimant Role = "CLAIMANT"
)

// RoleManager manages user roles
type RoleManager struct {
	mu    sync.RWMutex
	roles map[string][]Role // map of address to roles
}

// NewRoleManager creates a new role manager
func NewRoleManager() *RoleManager {
	return &RoleManager{
		roles: make(map[string][]Role),
	}
}

// AssignRole assigns a role to an address
func (rm *RoleManager) AssignRole(address string, role Role) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if roles, exists := rm.roles[address]; exists {
		// Check if role already assigned
		for _, r := range roles {
			if r == role {
				return nil
			}
		}
		rm.roles[address] = append(roles, role)
	} else {
		rm.roles[address] = []Role{role}
	}

	return nil
}

// HasRole checks if an address has a specific role
func (rm *RoleManager) HasRole(address string, role Role) bool {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	roles, exists := rm.roles[address]
	if !exists {
		return false
	}

	for _, r := range roles {
		if r == role {
			return true
		}
	}

	return false
}

// RemoveRole removes a role from an address
func (rm *RoleManager) RemoveRole(address string, role Role) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	roles, exists := rm.roles[address]
	if !exists {
		return fmt.Errorf("address not found")
	}

	// Filter out the role
	newRoles := make([]Role, 0)
	for _, r := range roles {
		if r != role {
			newRoles = append(newRoles, r)
		}
	}

	rm.roles[address] = newRoles
	return nil
}

// GetRoles returns all roles for an address
func (rm *RoleManager) GetRoles(address string) []Role {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	if roles, exists := rm.roles[address]; exists {
		return append([]Role{}, roles...)
	}
	return nil
}
