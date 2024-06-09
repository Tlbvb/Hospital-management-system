package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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

type AllAppointmentsOfPatientParams struct{

}

type AllAppointmentsOfPatientResult struct{

}


func (store *Store) AllAppointmentsOfPatient() []Appointment{
	store.execTx(context.Background(),func(q *Queries)error{
		
	})

}
//




//potential transactions to be done
//-> medical record and medication -> addition together