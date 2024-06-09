-- Step 1: Disable Foreign Key Constraints
ALTER TABLE "MedicalRecordMedication" DROP CONSTRAINT IF EXISTS "MedicalRecordMedication_medical_record_id_fkey";
ALTER TABLE "MedicalRecordMedication" DROP CONSTRAINT IF EXISTS "MedicalRecordMedication_medication_id_fkey";
ALTER TABLE "MedicalRecord" DROP CONSTRAINT IF EXISTS "MedicalRecord_appointment_id_fkey";
ALTER TABLE "MedicalRecord" DROP CONSTRAINT IF EXISTS "MedicalRecord_patient_id_fkey";
ALTER TABLE "MedicalRecord" DROP CONSTRAINT IF EXISTS "MedicalRecord_doctor_id_fkey";
ALTER TABLE "Appointment" DROP CONSTRAINT IF EXISTS "Appointment_patient_id_fkey";
ALTER TABLE "Appointment" DROP CONSTRAINT IF EXISTS "Appointment_doctor_id_fkey";
ALTER TABLE "Doctor" DROP CONSTRAINT IF EXISTS "Doctor_department_id_fkey";
ALTER TABLE "Department" DROP CONSTRAINT IF EXISTS "Department_head_id_fkey";

-- Step 2: Drop Tables
DROP TABLE IF EXISTS "MedicalRecordMedication";
DROP TABLE IF EXISTS "Medication";
DROP TABLE IF EXISTS "MedicalRecord";
DROP TABLE IF EXISTS "Appointment";
DROP TABLE IF EXISTS "Doctor";
DROP TABLE IF EXISTS "Patient";
DROP TABLE IF EXISTS "Admin";
DROP TABLE IF EXISTS "Department";

-- Step 3: Enable Foreign Key Constraints (if needed)
-- If you plan to recreate the tables and reinstate the foreign key constraints,
-- you can skip this step.

-- Enable Foreign Key Constraints (if needed)
-- ALTER TABLE "MedicalRecordMedication" ADD CONSTRAINT "MedicalRecordMedication_medical_record_id_fkey" FOREIGN KEY ("medical_record_id") REFERENCES "MedicalRecord" ("id");
-- ALTER TABLE "MedicalRecordMedication" ADD CONSTRAINT "MedicalRecordMedication_medication_id_fkey" FOREIGN KEY ("medication_id") REFERENCES "Medication" ("id");
-- ALTER TABLE "MedicalRecord" ADD CONSTRAINT "MedicalRecord_appointment_id_fkey" FOREIGN KEY ("appointment_id") REFERENCES "Appointment" ("id");
-- ALTER TABLE "MedicalRecord" ADD CONSTRAINT "MedicalRecord_patient_id_fkey" FOREIGN KEY ("patient_id") REFERENCES "Patient" ("id");
-- ALTER TABLE "MedicalRecord" ADD CONSTRAINT "MedicalRecord_doctor_id_fkey" FOREIGN KEY ("doctor_id") REFERENCES "Doctor" ("id");
-- ALTER TABLE "Appointment" ADD CONSTRAINT "Appointment_patient_id_fkey" FOREIGN KEY ("patient_id") REFERENCES "Patient" ("id");
-- ALTER TABLE "Appointment" ADD CONSTRAINT "Appointment_doctor_id_fkey" FOREIGN KEY ("doctor_id") REFERENCES "Doctor" ("id");
-- ALTER TABLE "Doctor" ADD CONSTRAINT "Doctor_department_id_fkey" FOREIGN KEY ("department_id") REFERENCES "Department" ("id");
-- ALTER TABLE "Department" ADD CONSTRAINT "Department_head_id_fkey" FOREIGN KEY ("head_id") REFERENCES "Doctor" ("id");
