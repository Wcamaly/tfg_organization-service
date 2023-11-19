INSERT INTO user_schema.role (id, name, description) VALUES (gen_random_uuid(), 'PATIENT', 'This role is from patient');
INSERT INTO user_schema.role (id, name, description) VALUES (gen_random_uuid(), 'DOCTOR', 'This role is from doctors and extend patientRole');
INSERT INTO user_schema.role (id, name, description) VALUES (gen_random_uuid(), 'MEDICAL_CENTER', 'This role is from medical center');