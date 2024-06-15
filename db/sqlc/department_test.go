package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/tlbvb/hospital-management/util"
)

func TestGetDepartmentById(t *testing.T) {
	//createDepartmentsAndDoctors(t)
	idtest:=util.RandomInt(1,10)
	department, err := TestQueries.GetDepartmentById(context.Background(), idtest)
	require.NoError(t, err)
	require.NotEmpty(t, department)
	require.Equal(t, idtest, department.ID)
}

func TestGetDepartmentByName(t *testing.T) {
	dep:=util.RandomDep()
	department, err := TestQueries.GetDepartmentByName(context.Background(), dep)
	require.NoError(t, err)
	require.NotEmpty(t, department)
	require.Equal(t,dep, department.Name)
}

func TestGetDepartmentHeadId(t *testing.T) {
	dep:=util.RandomDep()
	headID, err := TestQueries.GetDepartmentHeadId(context.Background(),dep)
	require.NoError(t, err)
	require.NotNil(t, headID)
}

func TestGetDepartmentByHead(t *testing.T) {
	headID:=util.RandomInt(1,10)
	department, err := TestQueries.GetDepartmentByHead(context.Background(), pgtype.Int8{Int64: headID,Valid: true})
	require.NoError(t, err)
	require.NotEmpty(t, department)
	require.Equal(t, headID, department.HeadID.Int64)
	//require.Equal(t, "Orthopedics", department.Name)
}