CREATE TABLE `users`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`name`)
VALUES ('Mateus'),
       ('Joao');

CREATE TABLE `badges`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `badges` (`name`)
VALUES ('specialist in chips'),
       ('specialist in trail visualization');

CREATE TABLE `companys`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `companys` (`name`)
VALUES ('aba'),
       ('austrian');

CREATE TABLE `jobs`
(
    id   bigint auto_increment,
    companyID bigint,
    title varchar(255) NOT NULL,
    job_description varchar(255) NOT NULL,
    istrial BOOLEAN,
    PRIMARY KEY (`id`),
    FOREIGN KEY (companyID) REFERENCES companys(id)
);

INSERT INTO `jobs` (`companyID`,`title`,`job_description`, `istrial`)
VALUES ('1', 'engineer in microeletronics', 'we need this guy to be the best engineer in microeletronics', true),
       ('2', 'specialist in analogic microeletronics', 'specialist in analogic microeletronics we will have a lot of difficulty jobs', false);

CREATE TABLE `coursers`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    companyID bigint ,
    workload bigint, 
    PRIMARY KEY (`id`),
    FOREIGN KEY (companyID) REFERENCES companys(id)
);

INSERT INTO `coursers` (`name`,`companyID`, workload)
VALUES ('chips construction', '1', 30),
       ('trail visualization', '2', 240);

CREATE TABLE `badges_users`
(
    id   bigint auto_increment,
    userID bigint ,
    badgeID bigint ,
    PRIMARY KEY (`id`),
    FOREIGN KEY (userID) REFERENCES users(id),
    FOREIGN KEY (badgeID) REFERENCES badges(id)
);

INSERT INTO `badges_users` (`userID`,`badgeID`)
VALUES ('1', '1'),
       ('1', '2'),
       ('2', '2');

CREATE TABLE `badges_jobs`
(
    id   bigint auto_increment,
    jobsID bigint ,
    badgeID bigint ,
    PRIMARY KEY (`id`),
    FOREIGN KEY (jobsID) REFERENCES jobs(id),
    FOREIGN KEY (badgeID) REFERENCES badges(id)
);

INSERT INTO `badges_jobs` (`jobsID`,`badgeID`)
VALUES ('1', '1'),
       ('1', '2'),
       ('2', '1'),
       ('2', '2');

CREATE TABLE `badges_coursers`
(
    id   bigint auto_increment,
    courserID bigint ,
    badgeID bigint ,
    PRIMARY KEY (`id`),
    FOREIGN KEY (courserID) REFERENCES coursers(id),
    FOREIGN KEY (badgeID) REFERENCES badges(id)
);

INSERT INTO `badges_coursers` (`courserID`,`badgeID`)
VALUES ('1', '1'),
       ('1', '2'),
       ('2', '2');


CREATE TABLE `users_coursers`
(
    id   bigint auto_increment,
    userID bigint ,
    courserID bigint ,
    workloadCompleted bigint,
    PRIMARY KEY (`id`),
    FOREIGN KEY (courserID) REFERENCES coursers(id),
    FOREIGN KEY (userID) REFERENCES users(id)
);

INSERT INTO `users_coursers` (`userID`,`courserID`, `workloadCompleted`)
VALUES ('1', '1', 20),
       ('2', '2', 40);


CREATE TABLE `tracks`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `tracks` (`name`)
VALUES ('analog'),
       ('digital');

CREATE TABLE `sub_tracks`
(
    id   bigint auto_increment,
    trackID bigint,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`, `trackID`),
    FOREIGN KEY (trackID) REFERENCES tracks(id)
);

INSERT INTO `sub_tracks` (`trackID`, `name`)
VALUES ('1', 'layout'),
       ('1', 'design'),
       ('2', 'bytesnbits');

CREATE TABLE `courses_sub_tracks`
(
    id   bigint auto_increment,
    courseID bigint,
    trackID bigint,
    subTrackID bigint,
    PRIMARY KEY (`id`, trackID, subTrackID),
    FOREIGN KEY (trackID, subTrackID) REFERENCES sub_tracks(trackID, id),
    FOREIGN KEY (courseID) REFERENCES coursers(id)
);

INSERT INTO `courses_sub_tracks` (`courseID`,`trackID`,`subTrackID`)
VALUES ('1', '1', '1'),
       ('2', '2', '3');
