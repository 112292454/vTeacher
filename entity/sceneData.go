package entity

type SceneData struct {
	Teacher   []string `json:"teacher" db:"teacher"` // teacher是一个字符串数组，存储该场景中教师的若干句话
	Student   []string `json:"student" db:"student"` // student是一个字符串数组，存储该场景中学生的若干句话
	SceneName string
}

// sceneData 定义请求参数结构
