package entity

type SceneData struct {
	Scenes    map[string]DialoguePair `json:"scenes" db:"scenes"` // scenes是一个Scene对象的数组，每个Scene包含教师和学生的对话
	SceneName string
}

// sceneData 定义请求参数结构
