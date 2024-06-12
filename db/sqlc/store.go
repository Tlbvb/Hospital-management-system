package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Store struct{
	*Queries
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) *Store{
	return &Store{
		Queries: New(db),
		db:db,
	}
}

func (store *Store) execTx(ctx context.Context,fn func(*Queries) error) error{
	tx,err:=store.db.BeginTx(context.Background(),pgx.TxOptions{})
	if err!=nil{
		return err
	}
	q:=New(tx)
	err=fn(q)
	if err!=nil{
		txerr:=tx.Rollback(ctx);if txerr!=nil{
			return fmt.Errorf("Transaction failed. Occured error during rollback")
		}
		return err
	}
	return tx.Commit(ctx)
}

type AppointmentParams struct{
	Time time.Time
	Doctor_id int
	Patient_id int
	Visitreason string
	EndTime time.Time
}

type AppointmentRegResult struct{
	Appointment Appointment
	MedRecord MedicalRecord
}
func (store *Store) AppointmentRegistration(arg AppointmentParams) (AppointmentRegResult,error){
	ctx:=context.Background()
	var appResult AppointmentRegResult
	err:=store.execTx(ctx,func(q *Queries) error{
		appointment,err:=q.CreateAppointment(ctx, CreateAppointmentParams{StartTime: arg.Time,EndTime: arg.EndTime,PatientID:pgtype.Int8{Int64: int64(arg.Patient_id), Valid: true}, DoctorID: pgtype.Int8{Int64: int64(arg.Doctor_id), Valid: true}, Visitreason: arg.Visitreason})
		if err!=nil{
			return err
		}
		medRecord,err:=q.CreateMedRecord(ctx, CreateMedRecordParams{PatientID: pgtype.Int8{Int64: int64(arg.Patient_id), Valid: true},DoctorID:  pgtype.Int8{Int64: int64(arg.Doctor_id), Valid: true},Diagnosis: "", Treatment:"", Notes: "",AppointmentID: pgtype.Int8{Int64: appointment.ID,Valid: true} })
		if err!=nil{
			return err
		}
		appResult.Appointment=appointment
		appResult.MedRecord=medRecord
		return nil
	})
	if err!=nil{
		return appResult,err
	}
	return appResult,nil
}

type UpdateMedParams struct{
	Diagnosis string
	Treatment string
	Notes string
	ID int
	Medications []Medication
}

type UpdateMedResult struct{
	MedRec MedicalRecord
	MedRecMed []MedicalRecordMedication
}

func (store *Store)UpdateMedRecordAddMedication(arg UpdateMedParams) (UpdateMedResult,error){
	ctx:=context.Background()
	var MedRes UpdateMedResult
	err:=store.execTx(ctx,func(queries *Queries)error{
		err:=queries.UpdateMedicalRecord(ctx,UpdateMedicalRecordParams{Diagnosis: arg.Diagnosis,ID: int64(arg.ID), Treatment: arg.Treatment,Notes: arg.Notes})
		if err!=nil{
			return err
		}
		recordMedications:=make([]MedicalRecordMedication,len(arg.Medications))
		for _,medic:=range arg.Medications{
			medrmed,err:=queries.CreateMedRecMedic(ctx,CreateMedRecMedicParams{MedicalRecordID: pgtype.Int8{Int64: int64(arg.ID),Valid: true},MedicationID: pgtype.Int8{Int64: int64(medic.ID),Valid: true}})
			recordMedications=append(recordMedications, medrmed)
			return err
		}
		medr,err:=queries.GetMedicalRecord(ctx,int64(arg.ID))
		MedRes.MedRec=medr
		MedRes.MedRecMed=recordMedications
		if err!=nil{
			return err
		}
		
		return nil

	})
	if err!=nil{
		return MedRes,err
	}
	return MedRes,nil

}


//potential transactions for the project: 
//Create an Appointment + Create a medical record
//Updating a medical record + add medical record medication
//if medication is not in db, add medication

//Updating department head id + doctor table department id

//deleting a patient and deleting all related stuff

//Users can register By themselves

//Admin registers doctors -> send them the temporary password and they need to change it with email validation
