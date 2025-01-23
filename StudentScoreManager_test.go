package main

import "testing"

// 测试 AddStudent 方法
func TestAddStudent(t *testing.T) {
	sm := NewStudentManager()

	// 创建一个本科生实例
	undergraduate := &Undergraduate{
		Student{
			Name:      "wei",
			StudentID: 1,
			Gender:    "male",
			Class:     "28",
		},
	}

	// 添加本科生
	sm.AddStudent(undergraduate)

	// 检查本科生是否正确添加
	student, err := sm.QueryStudent(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if student.Name != "wei" || student.StudentID != 1 || student.Gender != "male" || student.Class != "28" {
		t.Errorf("Expected student Alice, got %v", student)
	}

	// 创建一个研究生实例
	graduate := &Graduate{
		Student{
			Name:      "hao",
			StudentID: 2,
			Gender:    "female",
			Class:     "27",
		},
	}

	// 添加研究生
	sm.AddStudent(graduate)

	// 检查研究生是否正确添加
	student, err = sm.QueryStudent(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if student.Name != "hao" || student.StudentID != 2 || student.Gender != "female" || student.Class != "27" {
		t.Errorf("Expected student Bob, got %v", student)
	}
}

// TestDeleteStudent 测试 DeleteStudent 方法
func TestDeleteStudent(t *testing.T) {
	// 创建一个 StudentManager 实例
	sm := NewStudentManager()

	// 添加一些测试学生
	undergraduate := &Undergraduate{
		Student{
			Name:      "wei",
			StudentID: 1,
			Gender:    "male",
			Class:     "28",
		},
	}
	graduate := &Graduate{
		Student{
			Name:      "hao",
			StudentID: 2,
			Gender:    "female",
			Class:     "27",
		},
	}
	sm.AddStudent(undergraduate)
	sm.AddStudent(graduate)

	// 测试删除存在的学生
	err := sm.DeleteStudent(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	_, err = sm.QueryStudent(1)
	if err == nil {
		t.Errorf("Expected student with id 1 to be deleted, but found")
	}

	// 测试删除不存在的学生
	err = sm.DeleteStudent(3)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr := "student with id 3 not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}

	// 测试删除另一个存在的学生
	err = sm.DeleteStudent(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	_, err = sm.QueryStudent(2)
	if err == nil {
		t.Errorf("Expected student with id 2 to be deleted, but found")
	}
}

// TestModifyStudent 测试 ModifyStudent 方法
func TestModifyStudent(t *testing.T) {
	// 创建一个 StudentManager 实例
	sm := NewStudentManager()

	// 添加一些测试学生
	undergraduate := &Undergraduate{
		Student{
			Name:      "wei",
			StudentID: 1,
			Gender:    "male",
			Class:     "28",
		},
	}
	graduate := &Graduate{
		Student{
			Name:      "hao",
			StudentID: 2,
			Gender:    "female",
			Class:     "27",
		},
	}
	sm.AddStudent(undergraduate)
	sm.AddStudent(graduate)

	// 测试修改存在的学生
	updates := map[string]interface{}{
		"name":  "wei modified",
		"class": "28 modified",
	}
	err := sm.ModifyStudent(1, updates)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	updatedStudent, err := sm.QueryStudent(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if updatedStudent.Name != "wei modified" || updatedStudent.Gender != "male" || updatedStudent.Class != "28 modified" {
		t.Errorf("Expected student to be updated, got %v", updatedStudent)
	}

	// 测试修改不存在的学生
	err = sm.ModifyStudent(3, updates)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr := "student with id 3 not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}

	// 测试修改另一个存在的学生
	updates = map[string]interface{}{
		"name":  "hao modified",
		"class": "27 modified",
	}
	err = sm.ModifyStudent(2, updates)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	updatedStudent, err = sm.QueryStudent(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if updatedStudent.Name != "hao modified" || updatedStudent.Gender != "female" || updatedStudent.Class != "27 modified" {
		t.Errorf("Expected student to be updated, got %v", updatedStudent)
	}

	// 测试修改学生的性别
	updates = map[string]interface{}{
		"gender": "female",
	}
	err = sm.ModifyStudent(1, updates)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	updatedStudent, err = sm.QueryStudent(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if updatedStudent.Gender != "female" {
		t.Errorf("Expected student gender to be updated, got %v", updatedStudent.Gender)
	}

	// 测试修改学生的班级
	updates = map[string]interface{}{
		"class": "29",
	}
	err = sm.ModifyStudent(2, updates)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	updatedStudent, err = sm.QueryStudent(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if updatedStudent.Class != "29" {
		t.Errorf("Expected student class to be updated, got %v", updatedStudent.Class)
	}
}

// TestAddScore 测试 AddScore 方法
func TestAddScore(t *testing.T) {
	// 创建一个 StudentManager 实例
	sm := NewStudentManager()

	// 添加一些测试学生
	undergraduate := &Undergraduate{
		Student{
			Name:      "wei",
			StudentID: 1,
			Gender:    "male",
			Class:     "28",
		},
	}
	graduate := &Graduate{
		Student{
			Name:      "hao",
			StudentID: 2,
			Gender:    "female",
			Class:     "27",
		},
	}
	sm.AddStudent(undergraduate)
	sm.AddStudent(graduate)

	// 测试为存在的学生添加成绩
	err := sm.AddScore(1, "Math", 95.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	student, err := sm.QueryStudent(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if score, exists := student.Scores["Math"]; !exists || score != 95.0 {
		t.Errorf("Expected score 95.0 for course Math, got %v", student.Scores)
	}

	// 测试为不存在的学生添加成绩
	err = sm.AddScore(3, "Science", 88.0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr := "student with id 3 not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}

	// 测试为另一个存在的学生添加成绩
	err = sm.AddScore(2, "History", 85.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	student, err = sm.QueryStudent(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if score, exists := student.Scores["History"]; !exists || score != 85.0 {
		t.Errorf("Expected score 85.0 for course History, got %v", student.Scores)
	}
}

// TestDeleteScore 测试 DeleteScore 方法
func TestDeleteScore(t *testing.T) {
	// 创建一个 StudentManager 实例
	sm := NewStudentManager()

	// 添加一些测试学生
	undergraduate := &Undergraduate{
		Student{
			Name:      "wei",
			StudentID: 1,
			Gender:    "male",
			Class:     "28",
		},
	}
	graduate := &Graduate{
		Student{
			Name:      "hao",
			StudentID: 2,
			Gender:    "female",
			Class:     "27",
		},
	}
	sm.AddStudent(undergraduate)
	sm.AddStudent(graduate)

	// 为学生添加一些成绩
	err := sm.AddScore(1, "Math", 95.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	err = sm.AddScore(2, "Science", 88.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 测试删除存在的学生成绩
	err = sm.DeleteScore(1, "Math")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	student, err := sm.QueryStudent(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if _, exists := student.Scores["Math"]; exists {
		t.Errorf("Expected score for course Math to be deleted, got %v", student.Scores)
	}

	// 测试删除不存在的学生成绩
	err = sm.DeleteScore(1, "Science")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr := "score for course Science not found for student with id 1"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}

	// 测试删除另一个存在的学生成绩
	err = sm.DeleteScore(2, "Science")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	student, err = sm.QueryStudent(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if _, exists := student.Scores["Science"]; exists {
		t.Errorf("Expected score for course Science to be deleted, got %v", student.Scores)
	}

	// 测试删除不存在的学生的成绩
	err = sm.DeleteScore(3, "Math")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr = "student with id 3 not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}
}

// TestModifyScore 测试 ModifyScore 方法
func TestModifyScore(t *testing.T) {
	// 创建一个 StudentManager 实例
	sm := NewStudentManager()

	// 添加一些测试学生
	undergraduate := &Undergraduate{
		Student{
			Name:      "wei",
			StudentID: 1,
			Gender:    "male",
			Class:     "28",
		},
	}
	graduate := &Graduate{
		Student{
			Name:      "hao",
			StudentID: 2,
			Gender:    "female",
			Class:     "27",
		},
	}
	sm.AddStudent(undergraduate)
	sm.AddStudent(graduate)

	// 为学生添加一些成绩
	err := sm.AddScore(1, "Math", 95.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	err = sm.AddScore(2, "Science", 88.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 测试修改存在的学生成绩
	err = sm.ModifyScore(1, "Math", 90.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	student, err := sm.QueryStudent(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if score, exists := student.Scores["Math"]; !exists || score != 90.0 {
		t.Errorf("Expected score 90.0 for course Math, got %v", student.Scores)
	}

	// 测试修改不存在的学生成绩
	err = sm.ModifyScore(1, "Science", 85.0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr := "score for course Science not found for student with id 1"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}

	// 测试修改另一个存在的学生成绩
	err = sm.ModifyScore(2, "Science", 92.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	student, err = sm.QueryStudent(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if score, exists := student.Scores["Science"]; !exists || score != 92.0 {
		t.Errorf("Expected score 92.0 for course Science, got %v", student.Scores)
	}

	// 测试修改不存在的学生的成绩
	err = sm.ModifyScore(3, "Math", 80.0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr = "student with id 3 not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}
}

// TestQueryStudent 测试 QueryStudent 方法
func TestQueryStudent(t *testing.T) {
	// 创建一个 StudentManager 实例
	sm := NewStudentManager()

	// 添加一些测试学生
	undergraduate := &Undergraduate{
		Student{
			Name:      "wei",
			StudentID: 1,
			Gender:    "male",
			Class:     "28",
			Scores:    make(map[string]float64),
		},
	}
	graduate := &Graduate{
		Student{
			Name:      "hao",
			StudentID: 2,
			Gender:    "female",
			Class:     "27",
			Scores:    make(map[string]float64),
		},
	}
	sm.AddStudent(undergraduate)
	sm.AddStudent(graduate)

	// 为学生添加一些成绩
	err := sm.AddScore(1, "Math", 95.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	err = sm.AddScore(2, "Science", 88.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 修改学生的信息
	updates := map[string]interface{}{
		"name":   "wei modified",
		"gender": "female",
		"class":  "28 modified",
	}
	err = sm.ModifyStudent(1, updates)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	updates = map[string]interface{}{
		"name":   "hao modified",
		"gender": "male",
		"class":  "27 modified",
	}
	err = sm.ModifyStudent(2, updates)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 测试查询存在的学生
	student, err := sm.QueryStudent(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if student.Name != "wei modified" || student.StudentID != 1 || student.Gender != "female" || student.Class != "28 modified" {
		t.Errorf("Expected student to be updated, got %v", student)
	}

	// 测试查询另一个存在的学生
	student, err = sm.QueryStudent(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if student.Name != "hao modified" || student.StudentID != 2 || student.Gender != "male" || student.Class != "27 modified" {
		t.Errorf("Expected student to be updated, got %v", student)
	}

	// 测试查询不存在的学生
	_, err = sm.QueryStudent(3)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr := "student with id 3 not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}
}

// TestQueryScore 测试 QueryScore 方法
func TestQueryScore(t *testing.T) {
	// 创建一个 StudentManager 实例
	sm := NewStudentManager()

	// 添加一些测试学生
	undergraduate := &Undergraduate{
		Student{
			Name:      "wei",
			StudentID: 1,
			Gender:    "male",
			Class:     "28",
			Scores:    make(map[string]float64),
		},
	}
	graduate := &Graduate{
		Student{
			Name:      "hao",
			StudentID: 2,
			Gender:    "female",
			Class:     "27",
			Scores:    make(map[string]float64),
		},
	}
	sm.AddStudent(undergraduate)
	sm.AddStudent(graduate)

	// 为学生添加一些成绩
	err := sm.AddScore(1, "Math", 95.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	err = sm.AddScore(2, "Science", 88.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 修改学生的信息
	updates := map[string]interface{}{
		"name":   "wei modified",
		"gender": "female",
		"class":  "28 modified",
	}
	err = sm.ModifyStudent(1, updates)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	updates = map[string]interface{}{
		"name":   "hao modified",
		"gender": "male",
		"class":  "27 modified",
	}
	err = sm.ModifyStudent(2, updates)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 测试查询存在的学生成绩
	score, err := sm.QueryScore(1, "Math")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if score != 95.0 {
		t.Errorf("Expected score 95.0 for course Math, got %v", score)
	}

	// 测试查询另一个存在的学生成绩
	score, err = sm.QueryScore(2, "Science")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if score != 88.0 {
		t.Errorf("Expected score 88.0 for course Science, got %v", score)
	}

	// 测试查询不存在的学生成绩
	_, err = sm.QueryScore(1, "Science")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr := "score for course Science not found for student with id 1"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}

	// 测试查询不存在的学生的成绩
	_, err = sm.QueryScore(3, "Math")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedErr = "student with id 3 not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}
}
