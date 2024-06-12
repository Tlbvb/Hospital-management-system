package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateTestPatient(t *testing.T) Patient {
	params := CreatePatientParams{
		Fullname: "PatientUser",
		Address:  "123 Main St",
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

	patient, err := TestQueries.GetPatientByName(context.Background(), "PatientUser")
	require.NoError(t, err)
	require.NotEmpty(t, patient)
	require.Equal(t, createdPatient.ID, patient.ID)
	require.Equal(t, createdPatient.Fullname, patient.Fullname)
	require.Equal(t, createdPatient.Address, patient.Address)
	require.Equal(t, createdPatient.CreatedAt, patient.CreatedAt)
}