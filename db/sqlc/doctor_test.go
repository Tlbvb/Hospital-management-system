package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tlbvb/hospital-management/util"
)

func CreateTestDoctor(t *testing.T) Doctor {
	params := CreateDoctorParams{
		Fullname: util.RandomFullName(),
		Specialty:  util.RandomSpecialization(),
		DepartmentID: util.RandomInt(1,10),
	}
	doctor, err := TestQueries.CreateDoctor(context.Background(),params)
	fmt.Println(err)
	require.NoError(t, err)
	require.NotEmpty(t, doctor)
	require.Equal(t, params.Fullname, doctor.Fullname)
	require.Equal(t, params.Specialty, doctor.Specialty)
	require.Equal(t,int64(params.DepartmentID),doctor.DepartmentID)
	return doctor
}

func TestCreateDoctor(t *testing.T) {
	CreateTestDoctor(t)
}

func TestGetDoctorById(t *testing.T) {
	createdDoctor := CreateTestDoctor(t)
	doctor, err := TestQueries.GetDoctorById(context.Background(), createdDoctor.ID)
	require.NoError(t, err)
	require.NotEmpty(t, doctor)
	require.Equal(t, createdDoctor.ID, doctor.ID)
	require.Equal(t, createdDoctor.Fullname, doctor.Fullname)
	require.Equal(t, createdDoctor.Specialty, doctor.Specialty)
	require.Equal(t, createdDoctor.DepartmentID, doctor.DepartmentID)
}

func TestGetDoctorByName(t *testing.T) {
	createdDoctor := CreateTestDoctor(t)
	doctor, err := TestQueries.GetDoctorByName(context.Background(), createdDoctor.Fullname)
	require.NoError(t, err)
	require.NotEmpty(t, doctor)
	//require.Equal(t, createdDoctor.ID, doctor.ID)
	require.Equal(t, createdDoctor.Fullname, doctor.Fullname)
	require.Equal(t, createdDoctor.Specialty, doctor.Specialty)
	require.Equal(t, createdDoctor.DepartmentID, doctor.DepartmentID)
}

func TestGetDoctorsOfDepartment(t *testing.T) {
	departmentID := int64(8)
	CreateTestDoctor(t) // Ensure at least one doctor exists in the department

	doctors, err := TestQueries.GetDoctorsOfDepartment(context.Background(), departmentID)
	require.NoError(t, err)
	require.NotEmpty(t, doctors)

	for _, doctor := range doctors {
		require.Equal(t, departmentID, doctor.DepartmentID)
	}
}
