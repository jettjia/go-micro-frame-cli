package gen

const protoSrvTemplateContext = `
syntax = "proto3";
import "google/protobuf/empty.proto";
import public "google/protobuf/timestamp.proto";
import "common.proto";
import "category.proto";

option go_package = ".;proto";

service GoodsSrv {
    // 分类
    rpc CreateCategory (CategoryInfoRequest) returns (CategoryInfoResponse); // 新建
    rpc DeleteCategory (CategoryDeleteRequest) returns (google.protobuf.Empty); // 删
    rpc UpdateCategory (CategoryInfoRequest) returns (google.protobuf.Empty); // 修改
    rpc FindCategoryById (CategoryFindByIdRequest) returns (CategoryInfoResponse); // 根据id查找
    rpc QueryPageCategory (CategoryQueryPageRequest) returns (CategoryQueryPageResponse); // 分页List
}
`
