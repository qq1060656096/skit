# api设计标准

https://cloud.google.com/apis/design/standard_methods?hl=zh-cn

## 1.标准方法
```
下表描述了如何将标准方法映射到 HTTP 方法：
标准方法	HTTP 映射	HTTP 请求正文	HTTP 响应正文  示例
List	GET <collection URL>	无	资源*列表
Get	GET <resource URL>	无	资源*	http://DOMAIN_NAME/v1/shelves/1
Create	POST <collection URL>	资源	资源*
Update	PUT or PATCH <resource URL>	资源	资源*
Delete	DELETE <resource URL>	不适用	google.protobuf.Empty**
```
| 标准方法      | HTTP        | 映射              | HTTP 请求正文 | HTTP 响应正文 | 示例 |
| ----------- | ----------- |------------------ |------------ |-------------- |------------ |
| List        | GET         |  <collection URL> |  无         |    无        |  http://demo.com/v1/shelves      |
| Get         | GET         |  <resource URL>   |  无         |    无        |  http://demo.com/v1/shelves/1      |
| Create      | POST        |  <collection URL> |  无         |    无        |  http://demo.com/v1/shelves      |
| Update     | PUT or PATCH |  <resource URL>   |  无         |    无        |  http://demo.com/v1/shelves/1      |
| Delete      | DELETE      |  <resource URL>   |  无         |    无        |  http://demo.com/v1/shelves/1      |


| 自定义方法      | HTTP      | 映射   | 备注             |
| ----------- | ----------- |------ |----------------- |
| Cancel      | :cancel     |  POST | 取消一个未完成的操作 |
| BatchGet    | :batchGet   |  GET  | 批量获取多个资源    |
| Move        | :move	    |  POST | 将资源从一个父级移动到另一个父级 |
| Search      | :search     |  GET  | List 的替代方法，用于获取不符合 List 语义的数据 |
| Undelete    | :undelete   |  POST | 恢复之前删除的资源 |

示例
https://cloud.google.com/endpoints/docs/grpc/transcoding?hl=zh-cn

## 2.标准字段
| 字段名             | 字段类型     |  说明                     |
| ----------------- |----------- | ------------------------- |
name                | string	 | name 字段应包含相对资源名称。 |
parent              | string	 | 对于资源定义和 List/Create 请求，parent 字段应包含父级相对资源名称。 |
create_time         | Timestamp	 | 创建实体的时间戳。 |
update_time         | Timestamp	 | 最后更新实体的时间戳。注意：执行 create/patch/delete 操作时会更新 update_time。 |
delete_time         | Timestamp	 | 删除实体的时间戳，仅当它支持保留时才适用。 |
expire_time         | Timestamp	 | 实体到期时的到期时间戳。 |
start_time          | Timestamp	 | 标记某个时间段开始的时间戳。 |
end_time            | Timestamp	 | 标记某个时间段或操作结束的时间戳（无论其成功与否）。 |
read_time           | Timestamp	 | 应读取（如果在请求中使用）或已读取（如果在响应中使用）特定实体的时间戳。 |
time_zone           | string	 | 时区名称。它应该是 IANA TZ 名称，例如“America/Los_Angeles”。如需了解详情，请参阅 https://en.wikipedia.org/wiki/List_of_tz_database_time_zones。 |
region_code         | string	 | 位置的 Unicode 国家/地区代码 (CLDR)，例如“US”和“419”。如需了解详情，请访问 http://www.unicode.org/reports/tr35/#unicode_region_subtag。 |
language_code	    | string	 | BCP-47 语言代码，例如“en-US”或“sr-Latn”。如需了解详情，请参阅 http://www.unicode.org/reports/tr35/#Unicode_locale_identifier。 |
mime_type	        | string     | IANA 发布的 MIME 类型（也称为媒体类型）。如需了解详情，请参阅 https://www.iana.org/assignments/media-types/media-types.xhtml。 |
display_name	    | string	 | 实体的显示名称。 |
title	            | string	 | 实体的官方名称，例如公司名称。它应被视为 display_name 的正式版本。 |
description	        | string	 | 实体的一个或多个文本描述段落。 |
filter	            | string	 | List 方法的标准过滤器参数。请参阅 AIP-160。 |
query	            | string	 | 如果应用于搜索方法（即 :search），则与 filter 相同。 |
page_token	        | string	 | List 请求中的分页令牌。 |
page_size	        | int32	     | List 请求中的分页大小。 |
total_size	        | int32	     | 列表中与分页无关的项目总数。 |
next_page_token	    | string	 | List 响应中的下一个分页令牌。它应该用作后续请求的 page_token。空值表示不再有结果。 |
order_by	        | string	 | 指定 List 请求的结果排序。 |
progress_percent	| int32	     | 指定操作的进度百分比 (0-100)。值 -1 表示进度未知. |
request_id	        | string	 | 用于检测重复请求的唯一字符串 ID。 |
resume_token	    | string	 | 用于恢复流式传输请求的不透明令牌。 |
labels	            | map<string, string>  | 表示 Cloud 资源标签。 |
show_deleted	    | bool	     | 如果资源允许恢复删除行为，相应的 List 方法必须具有 show_deleted 字段，以便客户端可以发现已删除的资源。 |
update_mask	        | FieldMask	 | 它用于 Update 请求消息，该消息用于对资源执行部分更新。此掩码与资源相关，而不是与请求消息相关。 |
validate_only	    | bool	     | 如果为 true，则表示仅应验证给定请求，而不执行该请求。 |
```


## 3.错误

```
接口定义一小组标准错误，而不是定义不同类型的错误，使用Code返回错误类型，使用 Message 返回错误原因。
注意 Message 应该是对用户友好的提示。你不能假设 Message是 API开发员，它可能是 用户、客户端开发员，操作员，应用的最终用户。
示例：如购物车逻辑，可能返回商品找不到（GOODS_NOT_FOUND） 、活动找不到（PROMOTION_NOT_FOUND）、或者商品库存记录找不到（INVENTORY_NOT_FOUND）等。
不应该定义多种不同类型的 NOT_FOUND，只应该定义一种 google.rpc.Code.NOT_FOUND

```

```golang
type ErrorInfo struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Details interface{} `json:"details"`
}
```

```js
{
  "error": {
    "code": 404,
    "message": "商品不存在",
    "status": "NOT_FOUND",
    "details": [
      {
        "code": 404,
        "reason": "商品不存在",
        "domain": "googleapis.com",
        "metadata": {
          "service": "translate.googleapis.com"
        }
      }
    ]
  }
}
```

> 标准错误

| HTTP  | gRPC                | 说明                        |
| ----- | ------------------- | ------------------------- |
200	    | OK	              | 无错误。 |
400	    | INVALID_ARGUMENT	  | 客户端指定了无效参数。如需了解详情，请查看错误消息和错误详细信息。 |
400	    | FAILED_PRECONDITION | 请求无法在当前系统状态下执行，例如删除非空目录。 |
400	    | OUT_OF_RANGE	      | 客户端指定了无效范围。 |
401     | UNAUTHENTICATED	  | 由于 OAuth 令牌丢失、无效或过期，请求未通过身份验证。 |
403	    | PERMISSION_DENIED	  | 客户端权限不足。这可能是因为 OAuth 令牌没有正确的范围、客户端没有权限或者 API 尚未启用。 |
404	    | NOT_FOUND	          | 未找到指定的资源。 |
409	    | ABORTED	          | 并发冲突，例如读取/修改/写入冲突。 |
409	    | ALREADY_EXISTS	  | 客户端尝试创建的资源已存在。 |
429	    | RESOURCE_EXHAUSTED  | 资源配额不足或达到速率限制。如需了解详情，客户端应该查找 google.rpc.QuotaFailure 错误详细信息。 |
499	    | CANCELLED	          | 请求被客户端取消。 |
500	    | DATA_LOSS	          | 出现不可恢复的数据丢失或数据损坏。客户端应该向用户报告错误。 |
500	    | UNKNOWN	          | 出现未知的服务器错误。通常是服务器错误。 |
500	    | INTERNAL	          | 出现内部服务器错误。通常是服务器错误。 |
501	    | NOT_IMPLEMENTED	  | API 方法未通过服务器实现。 |
502	    | 不适用	              | 到达服务器前发生网络错误。通常是网络中断或配置错误。 |
503	    | UNAVAILABLE	      | 服务不可用。通常是服务器已关闭。 |
504	    | DEADLINE_EXCEEDED	  |  超出请求时限。仅当调用者设置的时限比方法的默认时限短（即请求的时限不足以让服务器处理请求）并且请求未在时限范围内完成时，才会发生这种情况。 |

## 命名

| API     | 示例                                | 说明                       |
| ------- | ---------------------------------- | ------------------------- |
| 产品名称  | Google Calendar API                | 产品名称是指 API 的产品营销名称，例如 Google Calendar API。API、界面、文档、服务条款、对帐单和商业合同等信息中使用的产品名称必须一致。Google API 必须使用产品团队和营销团队批准的产品名称 |
| 服务名称  | calendar.googleapis.com            | 服务名称应该是语法上有效的 DNS 名称（遵循 RFC 1035）可以解析为一个或多个网络地址。公开的 Google API 的服务名称采用 xxx.googleapis.com 格式。例如，Google 日历的服务名称是 calendar.googleapis.com |
| 软件包名称 | google.calendar.v3                 | 软件包名称 API .proto 文件中声明的软件包名称应该与产品名称和服务名称保持一致。软件包名称应该使用单数组件名称，以避免混合使用单数和复数组件名称。软件包名称不能使用下划线。进行版本控制的 API 的软件包名称必须以此版本结尾 |
| 接口名称  | google.calendar.v3.CalendarService | 接口名称应该使用直观的名词，例如 Calendar 或 Blob。该名称不应该与编程语言及其运行时库中的成熟概念（例如 File）相冲突。 在极少数情况下，接口名称会与 API 中的其他名称相冲突，此时应该使用后缀（例如 Api 或 Service）来消除歧义。 |
| 来源目录  | //google/calendar/v3               |  |
| API 名称 | calendar                           |  |

集合 ID 应采用复数和 lowerCamelCase（小驼峰式命名法）格式，并遵循美式英语拼写和语义。例如：events、children 或 deletedEvents。

### 方法名称
```
服务可以在其 IDL 规范中定义一个或多个远程过程调用 (RPC) 方法，这些方法需与集合和资源上的方法对应。方法名称应采用大驼峰式命名格式并遵循 VerbNoun 的命名惯例，其中 Noun（名词）通常是资源类型
```
| 动词     |  名词  |  方法名称    |  请求消息           |  响应消息            |
| ------- | ---- | ------------ | ------------------ | ------------------ |
| List    | Book  |  ListBooks  |  ListBooksRequest  | ListBooksResponse  |
| Get	  | Book  | GetBook	    | GetBookRequest	 | Book |
| Create  | Book  | CreateBook  | CreateBookRequest	 | Book |
| Update  | Book  | UpdateBook  | UpdateBookRequest	 | Book |
| Rename  | Book  | RenameBook  | RenameBoetokRequest	 | RenameBookResponse |
| Delete  | Book  | DeleteBook  | DeleBookRequest	 | google.protobuf.Empty |

### 枚举名称
```
枚举类型必须使用 UpperCamelCase 格式的名称。

enum FooBar {
  // The first value represents the default and must be == 0.
  FOO_BAR_UNSPECIFIED = 0;
  FIRST_VALUE = 1;
  SECOND_VALUE = 2;
}
```

## API文档
```
proto 文件中的注释格式
使用常用的 Protocol Buffers // 注释格式向 .proto 文件添加注释。


// Creates a shelf in the library, and returns the new Shelf.
rpc CreateShelf(CreateShelfRequest) returns (Shelf) {
  option (google.api.http) = { post: "/v1/shelves" body: "shelf" };
}
```

```
服务配置中的注释
另一种向 .proto 文件添加文档注释的方法是，您可以在其 YAML 服务配置文件中向 API 添加内嵌文档。如果两个文件中都记录了相同的元素，则 YAML 文件中的文档将优先于 .proto 中的文档。

documentation:
  summary: Gets and lists social activities
  overview: A simple example service that lets you get and list possible social activities
  rules:
  - selector: google.social.Social.GetActivity
    description: Gets a social activity. If the activity does not exist, returns Code.NOT_FOUND.
...
```

## 版本控制

```
所有 Google API 接口都必须提供一个主要版本号，该版本号在 protobuf 软件包的末尾编码，并作为 REST API 的 URI 路径的第一部分。如果 API 引入了重大更改（例如删除或重命名字段），则该 API 必须增加其 API 版本号，以确保现有用户代码不会突然中断。

GET /v1/projects/my-project/topics HTTP/1.1
Host: pubsub.googleapis.com
Authorization: Bearer y29....
X-Goog-Visibilities: PREVIEW
```

## 兼容性

```
向后兼容的（非重大）更改
向 API 服务定义添加 API 接口
从协议的角度来看，这始终比较安全。唯一需要注意的是，客户端库可能已经使用了手写代码中的新 API 接口名称。如果您的新接口与现有接口完全正交，则不太可能实现；如果它是现有接口的简化版本，则更有可能导致冲突。

向 API 接口添加方法
除非您添加的方法与客户端库中已生成的方法发生冲突，否则这应该没问题。

（可能出现重大更改的示例：如果您有 GetFoo 方法，则表示 C# 代码生成器已经创建了 GetFoo 和 GetFooAsync 方法。因此，从客户端库的角度来看，在 API 接口中添加 GetFooAsync 方法将是一个重大更改。）

向方法添加 HTTP 绑定
假设绑定没有引入任何歧义，让服务器响应之前拒绝的网址就是安全的。将现有操作应用于新资源名称模式时，可以执行此操作。


向请求消息添加字段
添加请求字段可以是非重大更改，前提是未指定该字段的客户端将在新版本中采用与旧版本相同的处理方式。

可能错误地执行此操作的最明显示例是使用分页：如果 API 的 v1.0 不包含集合的分页，则无法在 v1.1 中添加它，除非将默认的 page_size 视为无限（这通常是一个坏主意）。否则，希望通过单个请求获得完整结果的 v1.0 客户端可能只收到部分结果，而且不会意识到该集合包含更多资源。

向响应消息添加字段
并非资源（例如 ListBooksResponse）的响应消息可在不影响客户端的情况下进行扩展，前提是这样不会改变其他响应字段的行为。之前在响应中填充的任何字段都应继续使用相同的语义填充，即使这会引入冗余也如此。

例如，在 1.0 版中的查询响应可能包含 contained_duplicates 的布尔字段，这表示某些结果由于复制而省略。在 1.1 版中，我们可能会在 duplicate_count 字段提供更详细的信息。尽管从 1.1 版角度看是多余的，但仍必须填充 contained_duplicates 字段。
```

```
向后不兼容的（重大）的更改

移除或重命名服务、字段、方法或枚举值
从根本上看，如果客户端代码可能引用某些内容，对其执行移除或重命名操作就是重大更改，必须通过新的主要版本进行。引用旧名称的代码，对于有些语言（例如 C＃和 Java）会导致编译失败，对于其他语言则可能导致执行失败或数据丢失。传输格式兼容性与此无关。


更改 HTTP 绑定
此处的“更改”实际上是“删除和添加”。例如，如果您确定确实要支持 PATCH，但您发布的版本支持 PUT，或者您使用了错误的自定义动词名称，则可以添加新绑定，但不能因为相同原因而移除旧绑定，因为移除服务方法是一个重大更改。


更改字段的类型
即使新类型与传输格式兼容，这也可能更改客户端库生成的代码，因此必须通过新的主要版本进行。对于已编译的静态类型语言，这很容易引入编译时错误。


更改资源名称格式
资源不得更改其名称 - 这意味着不能更改集合名称。

与大多数重大更改不同，这也会影响主要版本：如果客户端可以使用 v2.0 访问在 v1.0 中创建的资源（反之亦然），则应在两个版本中使用相同的资源名称。

较容易忽略的是，由于以下原因，有效资源名称集也不应更改：

如果它的限制变得更严格，之前成功的请求现在将失败。
如果它没有之前记录的限制严格，基于先前文档做出假设的客户端可能无法正常工作。客户很可能采用对允许的字符集和名称长度敏感的方式，将资源名称存储在其他位置。或者，客户很可能执行自己的资源名称验证以遵循文档说明。（例如，在开始支持更长的 EC2 资源 ID 之前，亚马逊为客户提供了大量警告并且有一个迁移期。）
请注意，此类更改可能仅在原型文档中可见。 因此，在审核 CL 是否损坏时，仅查看非评论更改并不够。


更改现有请求的可见行为
客户通常依赖 API 行为和语义，即使此类行为没有得到明确支持或记录。因此，在大多数情况下，更改 API 数据的行为或语义造成的影响将被视为使用者的责任。如果行为未以加密方式隐藏，则应假设用户已发现并将依赖此行为。

由于这个原因（即使数据很无趣），对分页令牌加密也是一个好主意，可以防止用户创建自己的令牌，以及在令牌行为发生更改时影响令牌。
```

## 目录结构

```
API 服务通常使用 .proto 文件来定义 API 接口，并使用 .yaml 文件来配置 API 服务。每个 API 服务必须在 API 代码库中有一个 API 目录。API 目录应该包含所有 API 定义文件和构建脚本。

每个 API 目录应该具有以下标准布局：

API 目录

代码库必要条件

BUILD：构建文件。
METADATA：构建元数据文件。
OWNERS：API 目录所有者。
README.md：有关 API 服务的常规信息。
配置文件

{service}.yaml：基准服务配置文件，google.api.Service proto 消息的 YAML 表示法。
prod.yaml：生产环境增量服务配置文件。
staging.yaml：模拟环境增量服务配置文件。
test.yaml：测试环境增量服务配置文件。
local.yaml：本地环境增量服务配置文件。
文档文件

doc/*：技术文档文件。它们应采用 Markdown 格式。
接口定义

v[0-9]*/*：每个这样的目录都包含 API 的主要版本，主要是 proto 文件和构建脚本。
{subapi}/v[0-9]*/*：每个 {subapi} 目录都包含子 API 的接口定义。每个子 API 可以有自己的独立主要版本。
type/*：proto 文件，包含在不同 API 之间、同一 API 的不同版本之间或 API 与服务实现之间共享的类型。type/* 下的类型定义一旦发布就不应该有重大更改。
公共 Google API 定义在 GitHub 上发布，请参阅 Google API 代码库。如需详细了解目录结构，请参阅 Service Infrastructure 示例 API。
```

## 更新日志

```
2021-04
引入了基于公开范围的版本控制。

在术语表中引入了 API 标题。

2021-03
为仅限输出的字段添加了注释。

更新了枚举值指南，以始终包含明确的 _UNSPECIFIED 值。

添加有关如何生成和解析资源名称的指南。

向标准字段添加了 progress_percent。

2021-02
添加了有关 proto3 optional 原初字段的指南。
```


```pb
package helloworld;

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```