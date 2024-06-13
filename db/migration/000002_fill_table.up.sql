-- Insert 10 departments
INSERT INTO "Department" ("name", "description", "head_id")
VALUES
  ('Emergency', 'Provides immediate treatment for acute illnesses and trauma.', NULL),
  ('Intensive Care Unit', 'Specialized treatment and monitoring for critically ill patients.', NULL),
  ('Cardiology', 'Diagnoses and treats heart and blood vessel conditions.', NULL),
  ('Neurology', 'Diagnoses and treats disorders of the nervous system.', NULL),
  ('Pediatrics', 'Provides medical care for infants, children, and adolescents.', NULL),
  ('Obstetrics and Gynecology', 'Care of women giving birth and health of the female reproductive system.', NULL),
  ('Orthopedics', 'Diagnoses and treats skeletal deformities and injuries.', NULL),
  ('Radiology', 'Uses imaging technologies to diagnose and treat diseases.', NULL),
  ('Oncology', 'Diagnoses and treats cancer.', NULL),
  ('Surgery', 'Provides operative treatments of diseases and injuries.', NULL);

-- Insert 10 doctors and set them as heads of their respective departments
WITH inserted_doctors AS (
  INSERT INTO "Doctor" ("fullname", "specialty", "department_id")
  VALUES
    ('Dr. John Smith', 'Emergency Medicine', 1),
    ('Dr. Emily Johnson', 'Critical Care', 2),
    ('Dr. Michael Brown', 'Cardiology', 3),
    ('Dr. Sarah Davis', 'Neurology', 4),
    ('Dr. David Wilson', 'Pediatrics', 5),
    ('Dr. Laura Moore', 'Obstetrics and Gynecology', 6),
    ('Dr. Robert Taylor', 'Orthopedics', 7),
    ('Dr. Jessica Anderson', 'Radiology', 8),
    ('Dr. Thomas Jackson', 'Oncology', 9),
    ('Dr. Karen White', 'General Surgery', 10)
  RETURNING id, department_id
)
UPDATE "Department"
SET head_id = d.id
FROM inserted_doctors d
WHERE "Department".id = d.department_id;
