// 这里指定了 proto 文件的版本
syntax = "proto3";

// package 命名规则: product.application.version
package product.app.v1;

// go_package 生成 go 文件当中的包名
option go_package = "{{.Appname}}/api/product/app/v1";

import "google/api/annotations.proto";

service BlogService {

  // 创建文章
  rpc CreateArticle(CreateArticleReq) returns (CreateArticleResp) {
    option (google.api.http) = {
      post: "/article"
      body: "*"
    };
  }

  rpc GetArticle(GetArticleReq) returns (GetArticleResp) {
    option (google.api.http) = {
      get: "/article"
    };
  }
  
  // rpc ListArticleTags(ListArticleTagsReq) returns (ListArticleTagsResp) {

  // }
}

message CreateArticleReq {
  string title = 1;
  string content = 2;
}

message CreateArticleResp {}

message GetArticleReq {
  // @inject_tag: json:"id" form:"id" binding:"required"
  int32 id = 1;
}

message GetArticleResp {}

message ListArticleTagsReq {}

message ListArticleTagsResp {
  repeated Tag tags = 1;
}

message Tag {
  string key = 1;
  string value = 2;
}