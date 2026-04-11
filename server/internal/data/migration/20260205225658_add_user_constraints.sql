-- +goose Up
-- +goose StatementBegin

-- case-insensitive unique email
CREATE UNIQUE INDEX IF NOT EXISTS users_email_lower_idx ON users (LOWER(email));

-- check date of birth
ALTER TABLE IF EXISTS customers
ADD CONSTRAINT IF NOT EXISTS customers_dob_check
CHECK (date_of_birth IS NULL OR date_of_birth <= CURRENT_DATE);
ALTER TABLE IF EXISTS employees
ADD CONSTRAINT IF NOT EXISTS employees_dob_check
CHECK (date_of_birth IS NULL OR date_of_birth <= CURRENT_DATE);

-- check employee salary
ALTER TABLE IF EXISTS employees 
ADD CONSTRAINT IF NOT EXISTS employees_salary_check
CHECK (salary >= 0);

-- role profile validation
CREATE OR REPLACE FUNCTION validate_user_role_profile()
RETURNS TRIGGER AS $$
BEGIN
    -- when inserting/updating user
    IF TG_TABLE_NAME = 'users' THEN
        IF NEW.role = 'customer' THEN
            IF EXISTS (SELECT 1 FROM employees WHERE user_id = NEW.id) THEN
                RAISE EXCEPTION 'Customer cannot have employee profile';
            END IF;
        ELSE
            IF EXISTS (SELECT 1 FROM customers WHERE user_id = NEW.id) THEN
                RAISE EXCEPTION 'Employee cannot have customer profile';
            END IF;
        END IF;
    END IF;

    -- when inserting/updating customer
    IF TG_TABLE_NAME = 'customers' THEN
        IF EXISTS (
            SELECT 1 FROM users
            WHERE id = NEW.user_id
            AND role <> 'customer'
        ) THEN
            RAISE EXCEPTION 'Only customer role can have customer profile';
        END IF;
    END IF;

    -- when inserting/updating employee
    IF TG_TABLE_NAME = 'employees' THEN
        IF EXISTS (
            SELECT 1 FROM users
            WHERE id = NEW.user_id
            AND role = 'customer'
        ) THEN
            RAISE EXCEPTION 'Customer role cannot have employee profile';
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- triggers
CREATE TRIGGER trg_users_role_profile
BEFORE INSERT OR UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION validate_user_role_profile();

CREATE TRIGGER trg_customers_role_profile
BEFORE INSERT OR UPDATE ON customers
FOR EACH ROW
EXECUTE FUNCTION validate_user_role_profile();

CREATE TRIGGER trg_employees_role_profile
BEFORE INSERT OR UPDATE ON employees
FOR EACH ROW
EXECUTE FUNCTION validate_user_role_profile();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trg_users_role_profile ON users;
DROP TRIGGER IF EXISTS trg_customers_role_profile ON customers;
DROP TRIGGER IF EXISTS trg_employees_role_profile ON employees;

DROP FUNCTION IF EXISTS validate_user_role_profile;

DROP INDEX IF EXISTS users_email_lower_idx;
ALTER TABLE IF EXISTS customers DROP CONSTRAINT IF EXISTS customers_dob_check;
ALTER TABLE IF EXISTS employees DROP CONSTRAINT IF EXISTS employees_dob_check;
ALTER TABLE IF EXISTS employees DROP CONSTRAINT IF EXISTS employees_salary_check;
-- +goose StatementEnd
