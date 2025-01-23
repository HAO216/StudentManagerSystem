package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// StudentInterface 定义学生接口
type StudentInterface interface {
	GetID() int
	GetName() string
	GetGender() string
	GetClass() string
	GetScores() map[string]float64
	SetScores(scores map[string]float64)
}

// Student 结构体
type Student struct {
	Name      string             `json:"name"`
	StudentID int                `json:"id"`
	Gender    string             `json:"gender"`
	Class     string             `json:"class"`
	Scores    map[string]float64 `json:"scores"`
}

// Undergraduate 本科生结构体
type Undergraduate struct {
	Student
}

// Graduate 研究生结构体
type Graduate struct {
	Student
}

// 实现 StudentInterface 接口方法

// GetID 获取学生ID
func (u *Undergraduate) GetID() int {
	return u.StudentID
}

// GetName 获取学生姓名
func (u *Undergraduate) GetName() string {
	return u.Name
}

// GetGender 获取学生性别
func (u *Undergraduate) GetGender() string {
	return u.Gender
}

// GetClass 获取学生班级
func (u *Undergraduate) GetClass() string {
	return u.Class
}

// GetScores 获取学生成绩
func (u *Undergraduate) GetScores() map[string]float64 {
	return u.Scores
}

// SetScores 设置学生成绩
func (u *Undergraduate) SetScores(scores map[string]float64) {
	u.Scores = scores
}

// GetID 获取学生ID
func (g *Graduate) GetID() int {
	return g.StudentID
}

// GetName 获取学生姓名
func (g *Graduate) GetName() string {
	return g.Name
}

// GetGender 获取学生性别
func (g *Graduate) GetGender() string {
	return g.Gender
}

// GetClass 获取学生班级
func (g *Graduate) GetClass() string {
	return g.Class
}

// GetScores 获取学生成绩
func (g *Graduate) GetScores() map[string]float64 {
	return g.Scores
}

// SetScores 设置学生成绩
func (g *Graduate) SetScores(scores map[string]float64) {
	g.Scores = scores
}

// StudentManager 结构体
type StudentManager struct {
	students map[int]*Student
	mu       sync.Mutex
}

// NewStudentManager 初始化 StudentManager
// 初始化了一个空的学生映射，用于后续添加和管理学生信息
func NewStudentManager() *StudentManager {
	return &StudentManager{
		students: make(map[int]*Student),
	}
}

// AddStudent 添加学生信息
func (sm *StudentManager) AddStudent(student StudentInterface) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// 使用反射获取学生ID
	val := reflect.ValueOf(student)
	idField := val.MethodByName("GetID")
	studentID := idField.Call(nil)[0].Interface().(int)

	// 将学生信息添加到学生管理器的映射中，使用学生ID作为键
	sm.students[studentID] = &Student{
		Name:      student.GetName(),
		StudentID: student.GetID(),
		Gender:    student.GetGender(),
		Class:     student.GetClass(),
	}
}

// DeleteStudent 删除学生信息
func (sm *StudentManager) DeleteStudent(studentID int) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	// 检查学生ID是否存在于映射中
	if _, exists := sm.students[studentID]; exists {
		// 如果存在，则从映射中删除该学生
		delete(sm.students, studentID)
		return nil
	}
	// 如果不存在，返回错误信息
	return fmt.Errorf("student with id %d not found", studentID)
}

// ModifyStudent 修改学生信息
func (sm *StudentManager) ModifyStudent(studentID int, updates map[string]interface{}) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// 检查学生ID是否存在于映射中
	if student, exists := sm.students[studentID]; exists {
		// 更新学生信息
		if name, ok := updates["name"].(string); ok {
			student.Name = name
		}
		if gender, ok := updates["gender"].(string); ok {
			student.Gender = gender
		}
		if class, ok := updates["class"].(string); ok {
			student.Class = class
		}
		return nil
	}

	// 如果不存在，返回错误信息
	return fmt.Errorf("student with id %d not found", studentID)
}

// AddScore 为学生添加成绩
func (sm *StudentManager) AddScore(studentID int, courseName string, score float64) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	// 检查学生ID是否存在于映射中
	if student, exists := sm.students[studentID]; exists {
		// 如果存在且学生的成绩记录为空，初始化
		if student.Scores == nil {
			student.Scores = make(map[string]float64)
		}
		// 将课程分数添加到学生的成绩记录中
		student.Scores[courseName] = score
		return nil
	}
	// 如果不存在，返回错误信息
	return fmt.Errorf("student with id %d not found", studentID)
}

// DeleteScore 删除学生成绩
func (sm *StudentManager) DeleteScore(studentID int, courseName string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	// 检查学生ID是否存在于映射中
	if student, exists := sm.students[studentID]; exists {
		// 检查学生是否有指定课程的成绩记录
		if _, exists := student.Scores[courseName]; exists {
			// 如果课程成绩存在，删除课程成绩记录
			delete(student.Scores, courseName)
			return nil
		}
		// 如果课程成绩不存在，返回错误信息
		return fmt.Errorf("score for course %s not found for student with id %d", courseName, studentID)
	}
	// 如果学生不存在，返回错误信息
	return fmt.Errorf("student with id %d not found", studentID)
}

// ModifyScore 修改学生成绩
func (sm *StudentManager) ModifyScore(studentID int, courseName string, score float64) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	// 检查学生ID是否存在于映射中
	if student, exists := sm.students[studentID]; exists {
		// 检查学生是否有指定课程的成绩记录
		if _, exists := student.Scores[courseName]; exists {
			// 如果课程成绩存在，更新课程成绩
			student.Scores[courseName] = score
			return nil
		}
		// 如果课程成绩不存在，返回错误信息
		return fmt.Errorf("score for course %s not found for student with id %d", courseName, studentID)
	}
	// 如果学生不存在，返回错误信息
	return fmt.Errorf("student with id %d not found", studentID)
}

// QueryStudent 查询学生信息
func (sm *StudentManager) QueryStudent(studentID int) (*Student, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	// 检查学生ID是否存在于映射中
	if student, exists := sm.students[studentID]; exists {
		// 如果存在，返回学生信息
		return student, nil
	}
	// 如果不存在，返回错误信息
	return nil, fmt.Errorf("student with id %d not found", studentID)
}

// QueryScore 查询学生成绩
func (sm *StudentManager) QueryScore(studentID int, courseName string) (float64, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	// 检查学生ID是否存在于映射中
	if student, exists := sm.students[studentID]; exists {
		// 检查课程成绩是否存在
		if score, exists := student.Scores[courseName]; exists {
			// 如果课程成绩存在，返回课程成绩
			return score, nil
		}
		// 如果课程成绩不存在，返回错误信息
		return 0, fmt.Errorf("score for course %s not found for student with id %d", courseName, studentID)
	}
	// 如果学生不存在，返回错误信息
	return 0, fmt.Errorf("student with id %d not found", studentID)
}

func main() {
	// 创建 Gin 引擎
	r := gin.Default()
	// 创建学生管理器
	sm := NewStudentManager()

	// 增加本科生信息
	r.POST("/undergraduates", func(c *gin.Context) {
		var undergraduate Undergraduate
		if err := c.ShouldBindJSON(&undergraduate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		sm.AddStudent(&undergraduate)
		c.JSON(http.StatusCreated, gin.H{"message": "Undergraduate added successfully"})
	})

	// 增加研究生信息
	r.POST("/graduates", func(c *gin.Context) {
		var graduate Graduate
		if err := c.ShouldBindJSON(&graduate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		sm.AddStudent(&graduate)
		c.JSON(http.StatusCreated, gin.H{"message": "Graduate added successfully"})
	})

	// 删除学生信息
	r.DELETE("/students/:id", func(c *gin.Context) {
		studentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student id"})
			return
		}
		if err := sm.DeleteStudent(studentID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
	})

	// 修改学生信息
	r.PUT("/students/:id", func(c *gin.Context) {
		studentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student id"})
			return
		}

		var updates map[string]interface{}
		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := sm.ModifyStudent(studentID, updates); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Student modified successfully"})
	})

	// 增加学生成绩
	r.POST("/students/:id/scores", func(c *gin.Context) {
		studentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student id"})
			return
		}
		var scoreData struct {
			CourseName string  `json:"course_name"`
			Score      float64 `json:"score"`
		}
		if err := c.ShouldBindJSON(&scoreData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := sm.AddScore(studentID, scoreData.CourseName, scoreData.Score); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Score added successfully"})
	})

	// 删除学生成绩
	r.DELETE("/students/:id/scores/:course", func(c *gin.Context) {
		studentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student id"})
			return
		}
		courseName := c.Param("course")
		if err := sm.DeleteScore(studentID, courseName); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Score deleted successfully"})
	})

	// 修改学生成绩
	r.PUT("/students/:id/scores", func(c *gin.Context) {
		studentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student id"})
			return
		}
		var scoreData struct {
			CourseName string  `json:"course_name"`
			Score      float64 `json:"score"`
		}
		if err := c.ShouldBindJSON(&scoreData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := sm.ModifyScore(studentID, scoreData.CourseName, scoreData.Score); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Score modified successfully"})
	})

	// 查询学生信息
	r.GET("/students/:id", func(c *gin.Context) {
		// 获取路径参数 "id"
		studentIDStr := c.Param("id")
		if studentIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Student ID is required"})
			return
		}

		// 将路径参数转换为整数
		studentID, err := strconv.Atoi(studentIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student id"})
			return
		}

		// 查询学生信息
		student, err := sm.QueryStudent(studentID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// 返回学生信息
		c.JSON(http.StatusOK, student)
	})

	// 查询学生成绩
	r.GET("/students/:id/scores/:course", func(c *gin.Context) {
		// 获取路径参数 "id"
		studentIDStr := c.Param("id")
		if studentIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Student ID is required"})
			return
		}

		// 将路径参数转换为整数
		studentID, err := strconv.Atoi(studentIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student id"})
			return
		}

		// 获取路径参数 "course"
		courseName := c.Param("course")
		if courseName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Course name is required"})
			return
		}

		// 查询学生成绩
		score, err := sm.QueryScore(studentID, courseName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// 返回学生成绩
		c.JSON(http.StatusOK, gin.H{"score": score})
	})

	// 并发导入 CSV 数据
	r.POST("/import", func(c *gin.Context) {
		// 获取上传的文件
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
			return
		}
		defer file.Close()
		// 创建一个通道
		reader := csv.NewReader(file)
		ch := make(chan StudentInterface)
		var wg sync.WaitGroup
		// 启动一个协程，负责读取和解析 CSV 数据
		go func() {
			for {
				record, err := reader.Read()
				if err == io.EOF {
					close(ch)
					break
				}
				if err != nil {
					fmt.Println("Error reading CSV:", err)
					continue
				}
				// 解析 CSV 记录
				studentType := record[0]
				studentID, _ := strconv.Atoi(record[1])
				name := record[2]
				gender := record[3]
				class := record[4]

				var student StudentInterface
				switch studentType {
				case "undergraduate":
					student = &Undergraduate{
						Student{
							Name:      name,
							StudentID: studentID,
							Gender:    gender,
							Class:     class,
						},
					}
				case "graduate":
					student = &Graduate{
						Student{
							Name:      name,
							StudentID: studentID,
							Gender:    gender,
							Class:     class,
						},
					}
				default:
					fmt.Println("Unknown student type:", studentType)
					continue
				}

				wg.Add(1)
				// 启动一个新的协程，将学生数据发送到通道
				go func(s StudentInterface) {
					defer wg.Done()
					ch <- s
				}(student)
			}
		}()
		// 启动一个协程，等待所有解析协程完成，然后关闭通道
		go func() {
			wg.Wait()
			close(ch)
		}()
		// 遍历通道，接收学生数据并添加到学生管理器中
		for student := range ch {
			sm.AddStudent(student)
		}
		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{"message": "CSV data imported successfully"})
	})

	// 启动服务器
	r.Run(":8080")
}
