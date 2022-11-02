BEGIN TRANSACTION;

DO
$$
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = 'jobstatus') THEN
            CREATE TYPE jobStatus AS ENUM
                (
                    'planned',
                    'file_uploaded',
                    'noise_wrap_started',
                    'noise_wrap_complete',
                    'noise_wrap_error',
                    'speech_recognize_started',
                    'speech_recognize_complete',
                    'speech_recognize_error',
                    'complete',
                    'cancelled'
                    );
        END IF;
    END
$$;

DO
$$
BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = 'strategytype') THEN
            CREATE TYPE strategyType AS ENUM
                (
                    'default',
                    'long_file'
                    );
        END IF;
    END
$$;

CREATE TABLE IF NOT EXISTS public.jobs
(
    id               SERIAL UNIQUE PRIMARY KEY ,
    uuid             varchar(36)  NOT NULL default md5(random()::text || clock_timestamp()::text)::uuid,
    status           jobStatus NOT NULL default 'planned',
    file_name        varchar(256) NOT NULL,
    strategy         strategyType NOT NULL default 'default',
    created_at       timestamptz  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       timestamptz  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT TRANSACTION;