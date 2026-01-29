CREATE TABLE applications (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Required indexes
CREATE INDEX idx_applications_user_id ON applications(user_id);
CREATE INDEX idx_applications_created_at ON applications(created_at);
CREATE INDEX idx_applications_deleted_at ON applications(deleted_at);

CREATE TABLE application_status (
    id BIGSERIAL PRIMARY KEY,
    application_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_application_status_application
        FOREIGN KEY(application_id)
        REFERENCES applications(id)
        ON DELETE CASCADE
);

-- Required indexes
CREATE INDEX idx_application_status_user_id ON application_status(user_id);
CREATE INDEX idx_application_status_application_id ON application_status(application_id);
CREATE TABLE application_uploaded_file_type (
    id BIGSERIAL PRIMARY KEY,
    application_id BIGINT NOT NULL,
    file_type_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_application_file_type_application
        FOREIGN KEY(application_id)
        REFERENCES applications(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_application_file_type
        UNIQUE(application_id, file_type_name)
);
