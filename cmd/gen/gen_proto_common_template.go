package gen

const protoCommonTemplateContext = `
syntax = "proto3";
option go_package = ".;proto";

message Query {
    string key = 1; //表字段名称
    string value = 2; //表字段值
    Operator operator = 3; //判断条件
}

enum Operator {
    GT = 0; //大于
    EQUAL = 1; //等于
    LT = 2; //小于
    NEQ = 3; //不等于
    LIKE = 4; //模糊查询
    GTE = 5; // 大于等于
    LTE = 6; // 小于等于
}

message PageData {
    uint32 pageSize = 1; // 一页多少条数据
    uint32 page = 2; // 第几页数据
    uint32 totalNumber = 3; // 一共多少条数据
    uint32 totalPage = 4; // 一共多少页
}
`
