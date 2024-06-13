package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateTestDoctor(t *testing.T) Patient {
	params := CreateDoctorParams{
		Fullname: "PatientUser",
		Specialty:  "",
		DepartmentID: ,
	}
	patient, err := TestQueries.CreatePatient(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, patient)
	require.Equal(t, "PatientUser", patient.Fullname)
	require.Equal(t, "123 Main St", patient.Address)
	return patient
}

func TestCreatePatient(t *testing.T) {
	CreateTestPatient(t)
}

func TestGetPatientById(t *testing.T) {
	createdPatient := CreateTestPatient(t)

	patient, err := TestQueries.GetPatientById(context.Background(), createdPatient.ID)
	require.NoError(t, err)
	require.NotEmpty(t, patient)
	require.Equal(t, createdPatient.ID, patient.ID)
	require.Equal(t, createdPatient.Fullname, patient.Fullname)
	require.Equal(t, createdPatient.Address, patient.Address)
	require.Equal(t, createdPatient.CreatedAt, patient.CreatedAt)
}

func TestGetPatientByName(t *testing.T) {
	createdPatient := CreateTestPatient(t)
	fmt.Println(createdPatient)
	patient, err := TestQueries.GetPatientByName(context.Background(), "PatientUser")
	fmt.Println(patient, createdPatient)
	require.NoError(t, err)
	require.NotEmpty(t, patient)
	//require.Equal(t, createdPatient.ID, patient.ID)
	require.Equal(t, createdPatient.Fullname, patient.Fullname)
	require.Equal(t, createdPatient.Address, patient.Address)
	require.Equal(t, createdPatient.CreatedAt, patient.CreatedAt)
}
