package hungry_mode

/**
单例模式-饿汉模式
*/

//借助go的init函数来实现

// 饿汉式单例
// 注意定义非导出类型
type databaseConn struct {
	//变量
}

var dbConn *databaseConn

func init() {
	dbConn = &databaseConn{}
}

// GetInstance 获取实例
func Db() *databaseConn {
	return dbConn
}
