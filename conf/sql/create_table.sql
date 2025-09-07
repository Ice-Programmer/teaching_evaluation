-- 创建库
CREATE DATABASE IF NOT EXISTS teaching_evaluation CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 切换库
USE teaching_evaluation;

-- 班级表
CREATE TABLE IF NOT EXISTS student_class
(
    id           BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    class_number varchar(128)      not null comment '班级编号',
    class_name   VARCHAR(128)      NOT NULL COMMENT '班级名称',
    create_at    BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    is_delete    TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除',
    INDEX idx_class_name (class_name)
) COMMENT '班级表' CHARSET = utf8mb4
                   COLLATE = utf8mb4_unicode_ci;

-- 学生表
CREATE TABLE IF NOT EXISTS student
(
    id             BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    student_number VARCHAR(128)      NOT NULL COMMENT '学生学号',
    password       VARCHAR(256)      NOT NULL COMMENT '学生密码',
    student_name   VARCHAR(128)      NOT NULL COMMENT '学生姓名',
    gender         TINYINT DEFAULT 0 NOT NULL COMMENT '学生性别（0 - 女 1 - 男）',
    class_id       BIGINT            NOT NULL COMMENT '学生班级id',
    major          TINYINT DEFAULT 0 NOT NULL COMMENT '专业（0 - 计算机 1 - 自动化）',
    grade          INT               NOT NULL COMMENT '学生年级',
    status         TINYINT DEFAULT 0 NOT NULL COMMENT '学生状态（0 - 正常使用 1 - 拒绝访问）',
    create_at      BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    is_delete      TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除',
    INDEX idx_student_number (student_number),
    INDEX idx_class_id (class_id),
    INDEX idx_grade_major (grade, major)
) COMMENT '学生表' CHARSET = utf8mb4
                   COLLATE = utf8mb4_unicode_ci;

-- 教师表
CREATE TABLE IF NOT EXISTS teacher
(
    id            BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    teaching_name VARCHAR(256)      NOT NULL COMMENT '教师名称',
    gender        TINYINT DEFAULT 0 NOT NULL COMMENT '性别（0-女，1-男）',
    major         TINYINT DEFAULT 0 NOT NULL COMMENT '专业（0-计算机，1-自动化）',
    email         VARCHAR(512)      NULL COMMENT '邮箱',
    nationality   TINYINT DEFAULT 0 NOT NULL COMMENT '国籍（0-俄罗斯，1-中国）',
    create_at     BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    is_delete     TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除',
    INDEX idx_nationality_major (nationality, major)
) COMMENT '教师表' CHARSET = utf8mb4
                   COLLATE = utf8mb4_unicode_ci;

-- 课程表
CREATE TABLE IF NOT EXISTS course
(
    id             BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    course_cn_name VARCHAR(256)      NOT NULL COMMENT '课程中文名称',
    course_en_name VARCHAR(256)      NOT NULL COMMENT '课程英文名',
    grade          INT               NOT NULL COMMENT '学生年级',
    major          TINYINT DEFAULT 0 NOT NULL COMMENT '专业（0-计算机，1-自动化）',
    create_at      BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    is_delete      TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除',
    INDEX idx_grade_major (grade, major)
) COMMENT '课程表' CHARSET = utf8mb4
                   COLLATE = utf8mb4_unicode_ci;

-- 管理员表
CREATE TABLE IF NOT EXISTS admin
(
    id        BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    username  VARCHAR(256)      NOT NULL COMMENT '账号',
    password  VARCHAR(256)      NOT NULL COMMENT '密码',
    create_at BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    is_delete TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除',
    UNIQUE INDEX idx_username (username)
) COMMENT '管理员表' CHARSET = utf8mb4
                     COLLATE = utf8mb4_unicode_ci;

-- 评价指标表
CREATE TABLE IF NOT EXISTS target
(
    id             BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    target_cn_name VARCHAR(256)                NOT NULL COMMENT '指标中文名称',
    target_en_name VARCHAR(256)                NOT NULL COMMENT '指标英文名称',
    nationality    TINYINT        DEFAULT 0    NOT NULL COMMENT '国籍（0-俄罗斯，1-中国）',
    target_id      BIGINT                      NOT NULL COMMENT '指标id',
    weight         DECIMAL(10, 2) DEFAULT 0.00 NOT NULL COMMENT '指标权重',
    create_at      BIGINT         DEFAULT 0    NOT NULL COMMENT '创建时间',
    is_delete      TINYINT        DEFAULT 0    NOT NULL COMMENT '是否删除',
    INDEX idx_target_id (target_id),
    INDEX idx_nationality (nationality)
) COMMENT '评价指标表' CHARSET = utf8mb4
                       COLLATE = utf8mb4_unicode_ci;

-- 评测表
CREATE TABLE IF NOT EXISTS evaluation
(
    id              BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    evaluation_name VARCHAR(256)      NOT NULL COMMENT '评测名称',
    start_at        DATETIME          NOT NULL COMMENT '开始时间',
    end_at          DATETIME          NOT NULL COMMENT '结束时间',
    create_at       BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    is_delete       TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除',
    INDEX idx_evaluation_time (start_at, end_at)
) COMMENT '评测表' CHARSET = utf8mb4
                   COLLATE = utf8mb4_unicode_ci;

-- 一级指标分数表（修正表名）
CREATE TABLE IF NOT EXISTS target_score
(
    id            BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    student_id    BIGINT                      NOT NULL COMMENT '学生id',
    teacher_id    BIGINT                      NOT NULL COMMENT '教师id',
    course_id     BIGINT                      NOT NULL COMMENT '课程id',
    target_id     BIGINT                      NOT NULL COMMENT '一级指标id',
    evaluation_id BIGINT                      NOT NULL COMMENT '评测id',
    nationality   TINYINT                     NOT NULL COMMENT '教师国籍（0-俄罗斯，1-中国）',
    score         DECIMAL(10, 2) DEFAULT 0.00 NOT NULL COMMENT '一级分数',
    create_at     BIGINT         DEFAULT 0    NOT NULL COMMENT '创建时间',
    is_delete     TINYINT        DEFAULT 0    NOT NULL COMMENT '是否删除',
    INDEX idx_main_query (teacher_id, course_id, evaluation_id),
    INDEX idx_student_evaluation (student_id, evaluation_id)
) COMMENT '一级指标分数表' CHARSET = utf8mb4
                           COLLATE = utf8mb4_unicode_ci;

-- 总分表
CREATE TABLE IF NOT EXISTS score
(
    id                BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    teaching_id       BIGINT            NOT NULL COMMENT '教师id',
    course_id         BIGINT            NOT NULL COMMENT '课程id',
    evaluation_id     BIGINT            NOT NULL COMMENT '评测id',
    detail_score_list JSON              NOT NULL COMMENT '一级指标细则【json】',
    nationality       TINYINT           NOT NULL COMMENT '国籍（0-俄罗斯，1-中国）',
    score             INT     DEFAULT 0 NOT NULL COMMENT '总分',
    create_at         BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    is_delete         TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除',
    INDEX idx_evaluation_id (evaluation_id),
    INDEX idx_teacher_course (teaching_id, course_id)
) COMMENT '总分表' CHARSET = utf8mb4
                   COLLATE = utf8mb4_unicode_ci;

-- 红线表
CREATE TABLE IF NOT EXISTS red_line
(
    id            BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    teaching_id   BIGINT            NOT NULL COMMENT '教师id',
    course_id     BIGINT            NOT NULL COMMENT '课程id',
    target_id     BIGINT            NOT NULL COMMENT '指标id',
    evaluation_id BIGINT            NOT NULL COMMENT '评测id',
    score         INT               NOT NULL COMMENT '分数',
    create_at     BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    is_delete     TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除',
    INDEX idx_main_query (teaching_id, course_id, evaluation_id)
) COMMENT '红线表' CHARSET = utf8mb4
                   COLLATE = utf8mb4_unicode_ci;