package entity

// sceneData 定义请求参数结构

// 定义场景对象结构体
type DialoguePair struct {
	Teacher []string `json:"teacher" db:"teacher"` // teacher是一个字符串数组，存储该场景中教师的若干句话
	Student []string `json:"student" db:"student"` // student是一个字符串数组，存储该场景中学生的若干句话
}
