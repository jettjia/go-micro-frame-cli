package gen

const tClientContext = `
package test

import (
	"google.golang.org/grpc"

	goodsProto "mall.com/mall-proto/goods"
)

var GrpcClient goodsProto.GoodsClient
var ClientConn *grpc.ClientConn

func Init() {
	var err error
	ClientConn, err = grpc.Dial("192.168.106.1:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	GrpcClient = goodsProto.NewGoodsClient(ClientConn)
}
`

const tContext = `
package test

import (
	"context"
	"fmt"
	"testing"

	goodsProto "mall.com/mall-proto/goods"
)

// 创建-属性
func Test_CreateCategoryAttr(t *testing.T) {
	Init()
	rsp, err := GrpcClient.CreateCategoryAttr(context.Background(), &goodsProto.CategoryAttrInfoRequest{
		{{db}}
	})

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(rsp)

	defer ClientConn.Close()
}

// 删除-属性
func Test_DeleteCategoryAttr(t *testing.T) {
	Init()
	_, err := GrpcClient.DeleteCategoryAttr(context.Background(), &goodsProto.DeleteCategoryAttrRequest{
		Id: 1,
	})
	if err != nil {
		t.Error(err.Error())
	}

	defer ClientConn.Close()
}

// 修改-属性
func Test_UpdateCategoryAttr(t *testing.T) {
	Init()
	rsp, err := GrpcClient.UpdateCategoryAttr(context.Background(), &goodsProto.CategoryAttrInfoRequest{
		{{db}}
	})

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(rsp)

	defer ClientConn.Close()
}

// 查找
func Test_FindCategoryAttrById(t *testing.T) {
	Init()
	rsp, err := GrpcClient.FindCategoryAttrById(context.Background(), &goodsProto.FindCategoryAttrRequest{
		Id: 1,
	})

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(rsp)

	defer ClientConn.Close()
}
`