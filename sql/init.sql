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

INSERT INTO Teachers
VALUES ('teacherken@gmail.com'),
       ('teacherjoe@gmail.com');
INSERT INTO Students (email)
VALUES ('studentjon@gmail.com'),
       ('studenthon@gmail.com'),
       ('commonstudent1@gmail.com'),
       ('commonstudent2@gmail.com'),
       ('student_only_under_teacher_ken@gmail.com'),
       ('studentagnes@gmail.com'),
       ('studentbob@gmail.com'),
       ('studentmary@gmail.com'),
       ('studentmiche@gmail.com');
INSERT INTO TeacherStudents
VALUES ('teacherken@gmail.com', 'studentbob@gmail.com'),
       ('teacherken@gmail.com', 'commonstudent1@gmail.com'),
       ('teacherken@gmail.com', 'commonstudent2@gmail.com'),
       ('teacherken@gmail.com', 'student_only_under_teacher_ken@gmail.com'),
       ('teacherjoe@gmail.com', 'commonstudent1@gmail.com'),
       ('teacherjoe@gmail.com', 'commonstudent2@gmail.com');
