package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vTeacher/entity"
)

// GetSceneDialogue 根据场景ID获取对话
func GetSceneDialogue(c *gin.Context) {
	// 从URL参数中获取场景ID
	sceneID, err := strconv.Atoi(c.Param("cid"))
	if err != nil {
		// 如果无法将cid转换为int，返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid scene ID"})
		return
	}

	// 假设从某处获取到所有场景的数据，这里以一个示例静态数据替代
	// 实际应用中，你可能需要从数据库或其他数据源中查询这些数据
	scenes := []entity.SceneData{
		// 示例数据
	}

	// 检查sceneID是否有效
	if sceneID < 0 || sceneID >= len(scenes) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "scene ID out of range"})
		return
	}

	// 获取指定的场景
	selectedScene := scenes[sceneID]

	// 创建一个切片存储对话顺序
	dialogues := make([]string, 0)

	// 交替添加教师和学生的对话到dialogues切片
	maxLen := max(len(selectedScene.Teacher), len(selectedScene.Student))
	for i := 0; i < maxLen; i++ {
		if i < len(selectedScene.Teacher) {
			dialogues = append(dialogues, "Teacher: "+selectedScene.Teacher[i])
		}
		if i < len(selectedScene.Student) {
			dialogues = append(dialogues, "Student: "+selectedScene.Student[i])
		}
	}

	// 将处理好的对话返回给前端
	c.JSON(http.StatusOK, gin.H{"dialogues": dialogues})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
