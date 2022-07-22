package security

import (
	"api/data"
	"api/db"
)

// GetUserPermissions
func GetUserPermissions(userId *string, access *db.Access) (data.Permissions, error) {
	// Create data-access
	da := data.NewPermissionDA(access)
	temp := "USER"
	permissions, err := da.FindPermission(&data.Permission{PermissionId: userId, PermissionIdType: &temp}, &data.Permissions{data.CreateGenericPermission("VIEW", "PERMISSION", "USER")})
	if err != nil {
		return nil, err
	}

	// Get user roles
	dr := data.NewRoleDA(access)
	roles, err := dr.FindUserRole(&data.UserRole{UserId: userId}, &data.Permissions{data.CreateGenericPermission("VIEW", "USER", "ROLE")})
	if err != nil {
		return nil, err
	}

	// Get role permissions
	temp = "ROLE"
	for _, role := range roles {
		permission, err := da.FindPermission(&data.Permission{PermissionId: role.RoleId, PermissionIdType: &temp}, &data.Permissions{data.CreateGenericPermission("VIEW", "PERMISSION", "ROLE")})
		if err != nil {
			return nil, err
		}
		for _, entry := range permission {
			permissions = append(permissions, entry)
		}
	}

	// Convert Role Id permissions to User Id permissions
	for _, permission := range permissions {
		if *permission.PermissionTenant == "ROLE" {
			roleUsers, err := dr.FindUserRole(&data.UserRole{RoleId: permission.PermissionTenantId}, &data.Permissions{data.CreateGenericPermission("VIEW", "ROLE", "USER")})
			if err != nil {
				return nil, err
			}
			for _, roleUser := range roleUsers {
				permissionTenant := "USER"
				permissions = append(permissions, &data.Permission{Id: permission.Id, PermissionType: permission.PermissionType,
					PermissionCategory: permission.PermissionCategory, PermissionTenant: &permissionTenant, PermissionTenantId: roleUser.UserId, DateAdded: permission.DateAdded})
			}
		}
	}

	return permissions, nil
}

//RemoveRolePermissions removes role permissions from array
func RemoveRolePermissions(permissions *data.Permissions) *data.Permissions {
	var result data.Permissions
	for _, permission := range *permissions {
		if *permission.PermissionTenant != "ROLE" {
			result = append(result, permission)
		}
	}
	return &result
}

//RemoveRolePermissions removes user permissions from array
func RemoveUserPermissions(permissions *data.Permissions) *data.Permissions {
	var result data.Permissions
	for _, permission := range *permissions {
		if *permission.PermissionTenant != "USER" {
			result = append(result, permission)
		}
	}
	return &result
}
