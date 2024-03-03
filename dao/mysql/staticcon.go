package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"vTeacher/entity"
)

var db *sqlx.DB

func insertScenesAndDialogues(sceneData entity.SceneData) error {
	for sceneName, dialoguePairs := range sceneData.Scenes {
		// 插入场景
		result, err := db.Exec("INSERT INTO Scenes (Name) VALUES (?)", sceneName)
		if err != nil {
			return fmt.Errorf("inserting scene %s failed: %v", sceneName, err)
		}
		sceneID, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("getting last insert ID for scene %s failed: %v", sceneName, err)
		}

		// 插入对话
		for _, pair := range dialoguePairs {
			for _, teacherDialogue := range pair.Teacher {
				_, err := db.Exec("INSERT INTO TeacherDialogues (SceneID, Dialogue) VALUES (?, ?)", sceneID, teacherDialogue)
				if err != nil {
					return fmt.Errorf("inserting teacher dialogue failed: %v", err)
				}
			}
			for _, studentDialogue := range pair.Student {
				_, err := db.Exec("INSERT INTO StudentDialogues (SceneID, Dialogue) VALUES (?, ?)", sceneID, studentDialogue)
				if err != nil {
					return fmt.Errorf("inserting student dialogue failed: %v", err)
				}
			}
		}
	}

	return nil
}

func insertDialogues() {
	// 示例JSON字符串
	jsonStr := `这里是你的JSON字符串`

	// 解析JSON数据
	var sceneData SceneData
	if err := json.Unmarshal([]byte(jsonStr), &sceneData); err != nil {
		log.Fatalf("JSON unmarshal error: %v", err)
	}

	// 插入数据到数据库
	if err := insertScenesAndDialogues(sceneData); err != nil {
		log.Fatalf("Insert scenes and dialogues error: %v", err)
	}
}
