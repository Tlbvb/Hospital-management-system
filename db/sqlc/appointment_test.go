package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createTestAppointment(t *testing.T) Appointment {
	startTime := time.Now().Add(time.Hour)
	endTime := startTime.Add(time.Hour)
	fmt.Println(endTime)
	pat:=CreateTestPatient(t)
	doc:=CreateTestDoctor(t)
	params := CreateAppointmentParams{
		StartTime:  startTime,
		EndTime:    endTime,
		PatientID:  pat.ID,
		DoctorID:   doc.ID,
		Visitreason: "Checkup",
		Status:     "Scheduled",
	}
	appointment, err := TestQueries.CreateAppointment(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, appointment)
	require.WithinDuration(t, startTime, appointment.StartTime,time.Millisecond)
	require.WithinDuration(t, endTime, appointment.EndTime,time.Millisecond)
	require.Equal(t, pat.ID, appointment.PatientID)
	require.Equal(t, doc.ID, appointment.DoctorID)
	require.Equal(t, "Checkup", appointment.Visitreason)
	require.Equal(t, "Scheduled", appointment.Status)
	return appointment
}

func TestCreateAppointment(t *testing.T) {
	createTestAppointment(t)
}

func TestGetAppointment(t *testing.T) {
	createdAppointment := createTestAppointment(t)
	appointment, err := TestQueries.GetAppointment(context.Background(), createdAppointment.ID)
	require.NoError(t, err)
	require.NotEmpty(t, appointment)
	require.Equal(t, createdAppointment.ID, appointment.ID)
	require.Equal(t, createdAppointment.StartTime, appointment.StartTime)
	require.Equal(t, createdAppointment.EndTime, appointment.EndTime)
	require.Equal(t, createdAppointment.PatientID, appointment.PatientID)
	require.Equal(t, createdAppointment.DoctorID, appointment.DoctorID)
	require.Equal(t, "Checkup", appointment.Visitreason)
	require.Equal(t, createdAppointment.Status, appointment.Status)
}

func TestListAppointmentsPatient(t *testing.T) {

	app:=createTestAppointment(t)

	appointments, err := TestQueries.ListAppointmentsPatient(context.Background(), app.PatientID)
	require.NoError(t, err)
	require.NotEmpty(t, appointments)

	for _, appointment := range appointments {
		require.Equal(t, app.PatientID, appointment.PatientID)
	}
}

func TestListAppointmentsDoctor(t *testing.T) {
	app:=createTestAppointment(t)
	appointments, err := TestQueries.ListAppointmentsDoctor(context.Background(), app.DoctorID)
	require.NoError(t, err)
	require.NotEmpty(t, appointments)

	for _, appointment := range appointments {
		require.Equal(t, app.DoctorID, appointment.DoctorID)
	}
}

func TestListAppointmentsDoctorStatus(t *testing.T) {

	app:=createTestAppointment(t)

	appointments, err := TestQueries.ListAppointmentsDoctorStatus(context.Background(), ListAppointmentsDoctorStatusParams{DoctorID: app.DoctorID,Status:app.Status })
	require.NoError(t, err)
	require.NotEmpty(t, appointments)

	for _, appointment := range appointments {
		require.Equal(t, app.DoctorID, appointment.DoctorID)
		require.Equal(t, app.Status, appointment.Status)
	}
}

func TestListAppointmentsPatientStatus(t *testing.T) {
	app:=createTestAppointment(t)
	appointments, err := TestQueries.ListAppointmentsPatientStatus(context.Background(), ListAppointmentsPatientStatusParams{PatientID: app.PatientID,Status: app.Status})
	require.NoError(t, err)
	require.NotEmpty(t, appointments)

	for _, appointment := range appointments {
		require.Equal(t, app.PatientID, appointment.PatientID)
		require.Equal(t, app.Status, appointment.Status)
	}
}

func TestListAppointmentsPatientDoctorStatus(t *testing.T) {
	app:=createTestAppointment(t)
	appointments, err := TestQueries.ListAppointmentsPatientDoctorStatus(context.Background(), ListAppointmentsPatientDoctorStatusParams{DoctorID: app.DoctorID,Status: app.Status,PatientID: app.PatientID})
	require.NoError(t, err)
	require.NotEmpty(t, appointments)

	for _, appointment := range appointments {
		require.Equal(t, app.PatientID, appointment.PatientID)
		require.Equal(t, app.DoctorID, appointment.DoctorID)
		require.Equal(t, app.Status, appointment.Status)
	}
}

func TestUpdateAppointmentStatus(t *testing.T) {
	createdAppointment := createTestAppointment(t)
	newStatus := "Completed"

	err := TestQueries.UpdateAppointmentStatus(context.Background(), UpdateAppointmentStatusParams{ID: createdAppointment.ID,Status: newStatus})
	require.NoError(t, err)

	updatedAppointment, err := TestQueries.GetAppointment(context.Background(), createdAppointment.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAppointment)
	require.Equal(t, newStatus, updatedAppointment.Status)
}
