    CREATE TABLE tasks (
        id uuid NOT NULL DEFAULT gen_random_uuid(),
        title VARCHAR(255) NOT NULL,
        description TEXT ,
        status "task_status" NOT NULL DEFAULT 'PENDING',
        created_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP  NOT NULL  DEFAULT CURRENT_TIMESTAMP,
        user_id uuid NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users(id) 
        ON DELETE CASCADE
        ON UPDATE CASCADE
        );