package gen

import (
	"github.com/jettjia/go-micro-frame-cli/util"
	"os"
)

// todo
func Run() {
	// 1. 获取表完整结构信息
	// SHOW FULL COLUMNS FROM product

	InitDB("10.4.7.71", "3307", "root", "root", "zhe_pms")

	genReq := GenInit("goods_srv", "category")

	//fmt.Println(TableColumns)

	// 2. 生成项目文件结构
	CreateDir(genReq)

	// 3. 生成 model
	GenModel(genReq)

	// 4. 生成 repository
	GenRepository(genReq)

	// 4. 生成 service

	// 5. 生成 handler

	// 6. 生成 initialize

	// 7. proto

	// 格式化代码
	util.GoFmt(genReq.BaseDir)
}

// 创建需要的文件夹
func CreateDir(req GenReq) {
	os.MkdirAll(req.ModelDir, os.ModePerm)
	os.MkdirAll(req.RepositoryDir, os.ModePerm)
	os.MkdirAll(req.ServiceDir, os.ModePerm)
	os.MkdirAll(req.HandlerDir, os.ModePerm)
	os.MkdirAll(req.InitializeDir, os.ModePerm)
	os.MkdirAll(req.ProtoDir, os.ModePerm)
}
