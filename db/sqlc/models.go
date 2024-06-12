// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Admin struct {
	ID        int64     `json:"id"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"created_at"`
}

type Appointment struct {
	ID          int64       `json:"id"`
	StartTime   time.Time   `json:"start_time"`
	EndTime     time.Time   `json:"end_time"`
	PatientID   pgtype.Int8 `json:"patient_id"`
	DoctorID    pgtype.Int8 `json:"doctor_id"`
	Visitreason string      `json:"visitreason"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Department struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	HeadID      pgtype.Int8 `json:"head_id"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Doctor struct {
	ID           int64       `json:"id"`
	Fullname     string      `json:"fullname"`
	Specialty    string      `json:"specialty"`
	DepartmentID pgtype.Int8 `json:"department_id"`
	CreatedAt    time.Time   `json:"created_at"`
}

type MedicalRecord struct {
	ID            int64       `json:"id"`
	AppointmentID pgtype.Int8 `json:"appointment_id"`
	PatientID     pgtype.Int8 `json:"patient_id"`
	DoctorID      pgtype.Int8 `json:"doctor_id"`
	Diagnosis     string      `json:"diagnosis"`
	Treatment     string      `json:"treatment"`
	Notes         string      `json:"notes"`
	CreatedAt     time.Time   `json:"created_at"`
}

type MedicalRecordMedication struct {
	ID              int64       `json:"id"`
	MedicalRecordID pgtype.Int8 `json:"medical_record_id"`
	MedicationID    pgtype.Int8 `json:"medication_id"`
	CreatedAt       time.Time   `json:"created_at"`
}

type Medication struct {
	ID            int64       `json:"id"`
	Medicationame pgtype.Text `json:"medicationame"`
	Description   string      `json:"description"`
	Sideeffects   string      `json:"sideeffects"`
	CreatedAt     time.Time   `json:"created_at"`
}

type Patient struct {
	ID        int64     `json:"id"`
	Fullname  string    `json:"fullname"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
