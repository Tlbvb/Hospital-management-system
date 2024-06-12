CREATE TABLE "Appointment" (
  "id" bigserial PRIMARY KEY,
  "start_time" timestamptz NOT NULL UNIQUE,
  "end_time"  timestamptz NOT NULL UNIQUE,
  "patient_id" bigint,
  "doctor_id" bigint,
  "visitreason" varchar NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "MedicalRecord" (
  "id" bigserial PRIMARY KEY,
  "appointment_id" bigint,
  "patient_id" bigint,
  "doctor_id" bigint,
  "diagnosis" varchar NOT NULL,
  "treatment" varchar NOT NULL,
  "notes" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "Medication" (
  "id" bigserial PRIMARY KEY,
  "medicationame" varchar UNIQUE,
  "description" varchar NOT NULL,
  "sideeffects" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "MedicalRecordMedication" (
  "id" bigserial PRIMARY KEY,
  "medical_record_id" bigint,
  "medication_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "Doctor" (
  "id" bigserial PRIMARY KEY,
  "fullname" varchar NOT NULL,
  "specialty" varchar NOT NULL,
  "department_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "Patient" (
  "id" bigserial PRIMARY KEY,
  "fullname" varchar NOT NULL,
  "address" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "Admin" (
  "id" bigserial PRIMARY KEY,
  "fullname" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "Department" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "head_id" bigint UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "Appointment" ADD FOREIGN KEY ("patient_id") REFERENCES "Patient" ("id");

ALTER TABLE "Appointment" ADD FOREIGN KEY ("doctor_id") REFERENCES "Doctor" ("id");

ALTER TABLE "MedicalRecord" ADD FOREIGN KEY ("appointment_id") REFERENCES "Appointment" ("id");

ALTER TABLE "MedicalRecord" ADD FOREIGN KEY ("patient_id") REFERENCES "Patient" ("id");

ALTER TABLE "MedicalRecord" ADD FOREIGN KEY ("doctor_id") REFERENCES "Doctor" ("id");

ALTER TABLE "MedicalRecordMedication" ADD FOREIGN KEY ("medical_record_id") REFERENCES "MedicalRecord" ("id");

ALTER TABLE "MedicalRecordMedication" ADD FOREIGN KEY ("medication_id") REFERENCES "Medication" ("id");

ALTER TABLE "Doctor" ADD FOREIGN KEY ("department_id") REFERENCES "Department" ("id");

ALTER TABLE "Department" ADD FOREIGN KEY ("head_id") REFERENCES "Doctor" ("id");
