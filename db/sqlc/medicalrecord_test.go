package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tlbvb/hospital-management/util"
)

// func createTestAppointment(t *testing.T, patientID, doctorID int64, startTime, endTime time.Time, status string) Appointment {
// 	params := CreateAppointmentParams{
// 		StartTime:   startTime,
// 		EndTime:     endTime,
// 		PatientID:   patientID,
// 		DoctorID:    doctorID,
// 		VisitReason: "Checkup",
// 		Status:      status,
// 	}
// 	appointment, err := TestQueries.CreateAppointment(context.Background(), params)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, appointment)
// 	return appointment
// }

func createTestMedicalRecord(t *testing.T) MedicalRecord {
	app:=createTestAppointment(t)
	params := CreateMedRecordParams{
		AppointmentID: app.ID,
		Diagnosis:     util.RandomString(5),
		Treatment:     util.RandomString(10),
		Notes:         util.RandomString(10),
		PatientID:    app.PatientID,
		DoctorID:      app.DoctorID,
	}
	record, err := TestQueries.CreateMedRecord(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, record)
	require.Equal(t, app.ID, record.AppointmentID)
	require.Equal(t, app.PatientID, record.PatientID)
	require.Equal(t, app.DoctorID, record.DoctorID)
	require.Equal(t, params.Diagnosis, record.Diagnosis)
	require.Equal(t, params.Treatment, record.Treatment)
	require.Equal(t, params.Notes, record.Notes)
	return record
}

func TestCreateMedRecord(t *testing.T) {
	createTestMedicalRecord(t)
}

func TestGetMedicalRecord(t *testing.T) {
	record := createTestMedicalRecord(t)

	fetchedRecord, err := TestQueries.GetMedicalRecord(context.Background(), record.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedRecord)
	require.Equal(t, record.ID, fetchedRecord.ID)
	require.Equal(t, record.AppointmentID, fetchedRecord.AppointmentID)
	require.Equal(t, record.PatientID, fetchedRecord.PatientID)
	require.Equal(t, record.DoctorID, fetchedRecord.DoctorID)
	require.Equal(t, record.Diagnosis, fetchedRecord.Diagnosis)
	require.Equal(t, record.Treatment, fetchedRecord.Treatment)
	require.Equal(t, record.Notes, fetchedRecord.Notes)
}

func TestGetMedicalRecordAppointment(t *testing.T) {
	record := createTestMedicalRecord(t)

	fetchedRecord, err := TestQueries.GetMedicalRecordAppointment(context.Background(), record.AppointmentID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedRecord)
	require.Equal(t, record.ID, fetchedRecord.ID)
	require.Equal(t, record.AppointmentID, fetchedRecord.AppointmentID)
	require.Equal(t, record.PatientID, fetchedRecord.PatientID)
	require.Equal(t, record.DoctorID, fetchedRecord.DoctorID)
	require.Equal(t, record.Diagnosis, fetchedRecord.Diagnosis)
	require.Equal(t, record.Treatment, fetchedRecord.Treatment)
	require.Equal(t, record.Notes, fetchedRecord.Notes)
}

func TestGetMedicalRecordsDoctor(t *testing.T) {
	record:=createTestMedicalRecord(t)
	app,_:=TestQueries.GetAppointment(context.Background(),record.AppointmentID)
	records, err := TestQueries.GetMedicalRecordsDoctor(context.Background(), app.DoctorID)
	require.NoError(t, err)
	require.NotEmpty(t, records)

	for _, record := range records {
		require.Equal(t, app.DoctorID, record.DoctorID)
	}
}

func TestGetMedicalRecordsPatient(t *testing.T) {
	rec:=createTestMedicalRecord(t)
	app,_:=TestQueries.GetAppointment(context.Background(),rec.AppointmentID)
	records, err := TestQueries.GetMedicalRecordsPatient(context.Background(), app.PatientID)
	require.NoError(t, err)
	require.NotEmpty(t, records)

	for _, record := range records {
		require.Equal(t, app.PatientID, record.PatientID)
	}
}

func TestUpdateMedicalRecord(t *testing.T) {
	record := createTestMedicalRecord(t)

	newDiagnosis := "Common Cold"
	newTreatment := "Rest, fluids, and over-the-counter medication"
	newNotes := "No follow-up needed unless symptoms worsen"

	params := UpdateMedicalRecordParams{
		ID:        record.ID,
		Diagnosis: newDiagnosis,
		Treatment: newTreatment,
		Notes:     newNotes,
	}

	err := TestQueries.UpdateMedicalRecord(context.Background(), params)
	require.NoError(t, err)

	updatedRecord, err := TestQueries.GetMedicalRecord(context.Background(), record.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRecord)
	require.Equal(t, newDiagnosis, updatedRecord.Diagnosis)
	require.Equal(t, newTreatment, updatedRecord.Treatment)
	require.Equal(t, newNotes, updatedRecord.Notes)
}
