// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"Gf-Vben/app/dao/internal"
)

// appUserDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type appUserDao struct {
	internal.AppUserDao
}

var (
	// AppUser is globally public accessible object for table app_user operations.
	AppUser = appUserDao{
		internal.AppUser,
	}
)

// Fill with you ideas below.
