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
    rpc DeleteCategory (DeleteCategoryRequest) returns (google.protobuf.Empty); // 删
    rpc UpdateCategory (CategoryInfoRequest) returns (google.protobuf.Empty); // 修改
    rpc FindCategoryById (FindCategoryRequest) returns (CategoryInfoResponse); // 根据id查找
    rpc QueryPageCategory (QueryPageCategoryRequest) returns (QueryPageCategoryResponse); // 分页
}
`
