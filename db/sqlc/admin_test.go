package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)
func CreateAdmin(t *testing.T) Admin{
	admin,err:=TestQueries.CreateAdmin(context.Background(),"AdminUser")
	fmt.Println(admin)
	fmt.Printf("%T",admin.Fullname)
	require.NoError(t,err)
	require.Equal(t,"AdminUser",admin.Fullname)
	return admin
}

func TestAdminCreate(t *testing.T) {
	CreateAdmin(t)
}

func TestGetAdminById(t *testing.T) {
	createdAdmin:=CreateAdmin(t)

	admin, err := TestQueries.GetAdminById(context.Background(), createdAdmin.ID)
	require.NoError(t, err)
	require.NotEmpty(t, admin)
	require.Equal(t, createdAdmin.ID, admin.ID)
	require.Equal(t, createdAdmin.Fullname, admin.Fullname)
	require.Equal(t, createdAdmin.CreatedAt, admin.CreatedAt)
}

func TestGetAdminByName(t *testing.T) {
	createdAdmin:=CreateAdmin(t)
	admin, err := TestQueries.GetAdminByName(context.Background(), "AdminUser")
	require.NoError(t, err)
	require.NotEmpty(t, admin)
	//require.Equal(t, createdAdmin.ID, admin.ID)
	require.Equal(t, createdAdmin.Fullname, admin.Fullname)
	require.Equal(t, createdAdmin.CreatedAt, admin.CreatedAt)
}

