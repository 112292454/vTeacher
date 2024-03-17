package redis

// redis key 注意使用命名空间的方式，方便查询和拆分
const (
	KeyPostInfoHashPrefix = "vTeacher-plus:post:"
	KeyPostTimeZSet       = "vTeacher-plus:post:time"  // zset;帖子及发帖时间定义
	KeyPostScoreZSet      = "vTeacher-plus:post:score" // zset;帖子及投票分数定义
	// KeyPostVotedUpSetPrefix   = "vTeacher-plus:post:voted:down:"
	// KeyPostVotedDownSetPrefix = "vTeacher-plus:post:voted:up:"
	KeyPostVotedZSetPrefix    = "vTeacher-plus:post:voted:" // zSet;记录用户及投票类型;参数是post_id
	KeyCommunityPostSetPrefix = "vTeacher-plus:community:"  // set保存每个分区下帖子的id
)
