package main

import (
	"fmt"
	"vTeacher/config"
	"vTeacher/dao/mysql"
	"vTeacher/log"

	"vTeacher/pkg/snowflake"
	"vTeacher/routers"
)

// @host 127.0.0.1:8081
// @BasePath /api/v1/
func main() {
	// var confFile string
	// flag.StringVar(&confFile, "conf", "./conf/config.yaml", "配置文件")
	// flag.Parse()
	// 加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := log.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	/*
		if err := redis.Init(settings.Conf.RedisConfig); err != nil {
			fmt.Printf("init redis failed, err:%v\n", err)
			return
		}
		defer redis.Close()*/

	// 雪花算法生成分布式ID
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	/*	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator Trans failed,err:%v\n", err)
		return
	}*/
	// 注册路由
	r := routers.SetupRouter(config.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
