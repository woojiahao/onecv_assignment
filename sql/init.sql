CREATE TABLE Teachers
(
    email VARCHAR(256) PRIMARY KEY
);

CREATE TABLE Students
(
    email        VARCHAR(256) PRIMARY KEY,
    is_suspended BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE TeacherStudents
(
    teacher_email VARCHAR(256) NOT NULL REFERENCES Teachers,
    student_email VARCHAR(256) NOT NULL REFERENCES Students,
    PRIMARY KEY (teacher_email, student_email)
);