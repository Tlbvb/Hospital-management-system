package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tlbvb/hospital-management/util"
)

func createTestMedication(t *testing.T) Medication {
	params := CreateMedicationParams{
		Medicationame: util.RandomString(7),
		Description:   util.RandomString(15),
		Sideeffects:   util.RandomString(9),
	}
	medication, err := TestQueries.CreateMedication(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, medication)
	require.Equal(t, params.Medicationame, medication.Medicationame)
	require.Equal(t, params.Description, medication.Description)
	require.Equal(t, params.Sideeffects, medication.Sideeffects)
	return medication
}

func TestCreateMedication(t *testing.T) {
	createTestMedication(t)
}

func TestGetMedication(t *testing.T) {
	createdMedication := createTestMedication(t)
	fetchedMedication, err := TestQueries.GetMedication(context.Background(), createdMedication.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedMedication)
	require.Equal(t, createdMedication.ID, fetchedMedication.ID)
	require.Equal(t, createdMedication.Medicationame, fetchedMedication.Medicationame)
	require.Equal(t, createdMedication.Description, fetchedMedication.Description)
	require.Equal(t, createdMedication.Sideeffects, fetchedMedication.Sideeffects)
}

func TestGetMedicationByName(t *testing.T) {
	createdMedication := createTestMedication(t)
	fetchedMedication, err := TestQueries.GetMedicationByName(context.Background(), createdMedication.Medicationame)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedMedication)
	require.Equal(t, createdMedication.ID, fetchedMedication.ID)
	require.Equal(t, createdMedication.Medicationame, fetchedMedication.Medicationame)
	require.Equal(t, createdMedication.Description, fetchedMedication.Description)
	require.Equal(t, createdMedication.Sideeffects, fetchedMedication.Sideeffects)
}
