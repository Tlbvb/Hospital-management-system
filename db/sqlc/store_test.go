package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tlbvb/hospital-management/util"
)



func TestAppointmentReg(t *testing.T){
	doctor:=CreateTestDoctor(t)
	pat:=CreateTestPatient(t)
	appTime:=util.RandomTime()
	params:=AppointmentParams{
		Time: appTime,
		Doctor_id: int(doctor.ID),
		Patient_id: int(pat.ID),
		Visitreason: util.RandomString(10),
		EndTime: appTime.Add(time.Hour),
	}

	res,err:=TestStore.AppointmentRegistration(params)
	require.NoError(t,err)
	require.Equal(t,res.Appointment.DoctorID,int64(params.Doctor_id))
	require.Equal(t,res.Appointment.PatientID,int64(params.Patient_id))
	require.Equal(t, params.Visitreason, res.Appointment.Visitreason, "Visit reason mismatch")
	require.WithinDuration(t, params.Time, res.Appointment.StartTime, time.Millisecond,"Start time mismatch")
	require.WithinDuration(t, params.EndTime, res.Appointment.EndTime,time.Millisecond ,"End time mismatch")

	

}