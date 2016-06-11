// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// doc 包含了对文档各个结构的定义及对代码块的解析
package doc

import "sync"

// Doc 表示一个项目的完整文档列表。
type Doc struct {
	Title   string // 文档标题
	Version string // 文档的版本号
	Content string // 首页的简要介绍内容
	Apis    []*API
	mux     sync.Mutex // 控制 Apis 字段的多协程写入
}

// API 表示一个 API 文档。
type API struct {
	Method      string    // 请求的方法，GET，POST 等
	URL         string    // 请求地址
	Summary     string    // 简要描述
	Description string    // 详细描述
	Group       string    // 所属分组
	Queries     []*Param  // 查询参数
	Params      []*Param  // URL 参数
	Request     *Request  // 若是 GET，则使用此描述请求的具体数据
	Success     *Response // 成功时的响应内容
	Error       *Response // 出错时的响应内容
}

// Request 表示用户请求所表示的数据。
type Request struct {
	Type     string            // 请求的数据类型，多个用逗号分隔
	Headers  map[string]string // 请求必须携带的头
	Params   []*Param          //提交的各个字段的描述
	Examples []*Example        // 请求数据的示例
}

// Response 表示一次请求或是返回的数据。
// 一般为 @apiSuccess 或是 @apiError 所表示的数据。
type Response struct {
	Code     string            // HTTP 状态码
	Summary  string            // 该状态下的简要描述
	Headers  map[string]string // 返回的头信息。
	Params   []*Param          // 返回数据的各个字段描述
	Examples []*Example        // 返回数据的示例
}

// Param 用于描述提交和返回的参数信息。
type Param struct {
	Name    string // 参数名称
	Type    string // 类型
	Summary string // 参数介绍
}

// Example 示例代码
type Example struct {
	Type string // 示例代码的类型，xml 或是 json
	Code string // 示例代码
}

func New() *Doc {
	return &Doc{
		Apis: make([]*API, 0, 100),
	}
}
