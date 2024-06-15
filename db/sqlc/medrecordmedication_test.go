package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
)

func createTestMedRecMedic(t *testing.T) MedicalRecordMedication {
	rec:=createTestMedicalRecord(t)
	medic:=createTestMedication(t)
	params := CreateMedRecMedicParams{
		MedicalRecordID: rec.ID,
		MedicationID:    medic.ID,
	}
	medRecMedic, err := TestQueries.CreateMedRecMedic(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, medRecMedic)
	require.Equal(t, params.MedicalRecordID, medRecMedic.MedicalRecordID)
	require.Equal(t, params.MedicationID, medRecMedic.MedicationID)
	return medRecMedic
}

func TestCreateMedRecMedic(t *testing.T) {

	createTestMedRecMedic(t)
}

func TestGetMedicationByRecord(t *testing.T) {
	createdMedRecMedic := createTestMedRecMedic(t)
	// Fetch the medication by medical record ID
	medRecMedications, err := TestQueries.GetMedicationByRecord(context.Background(), createdMedRecMedic.MedicalRecordID)
	require.NoError(t, err)
	require.NotEmpty(t, medRecMedications)

	// Verify that the fetched records match the created one
	require.Equal(t, createdMedRecMedic.ID, medRecMedications[0].ID)
	require.Equal(t, createdMedRecMedic.MedicalRecordID, medRecMedications[0].MedicalRecordID)
	require.Equal(t, createdMedRecMedic.MedicationID, medRecMedications[0].MedicationID)
}
