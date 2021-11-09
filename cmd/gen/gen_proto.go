package gen

func GenProto(req GenReq) {
	doGenAutoProto(req) // 生成自动编译proto文件
	doGenCommon()       // 生成proto公共部分
	doGenSrv(req)       // 生成proto server部分
	doGenMessage(req)   // 生成表操作的具体逻辑部分
}

func doGenCommon() {

}

// auto proto grpc
func doGenAutoProto(req GenReq) {

}

func doGenSrv(req GenReq) {

}

func doGenMessage(req GenReq) {
	
}
