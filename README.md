# 自己用的 blog，写的非常难，就别看了

#### README 也给自己看的,有时候不记得

My personal blog backend(api) powerd by golang.

## Response model

`status_code` will allways be 200, unless the path is correct.

```json
{
    code:
    msg:
    error:
    data:
}
```

## API

| method | path         |     data     | meaning           |             auth required              |
| ------ | ------------ | :----------: | ----------------- | :------------------------------------: |
| POST   | /checkToken  |    token     | 检查 token 有效性 |                   no                   |
| GET    | /users/1     |      -       | 个人信息          |                   no                   |
| GET    | /posts       |      -       | 帖子列表          | no(如果不包含验证头，只会显示公开内容) |
| POST   | /posts       | 帖子所需字段 | 新增帖子          |                  yes                   |
| GET    | /posts/\<id> |      -       | 帖子详情          |                   no                   |
| PUT    | /posts/\<id> | 帖子所需字段 | 更新帖子          |                  yes                   |
| DELETE | /posts/\<id> |      -       | 删除帖子          |                  yes                   |

## code example

| code | meaning    |
| ---- | ---------- |
| 401  | 权限不足   |
| 404  | 内容不存在 |
| 500  | 服务器错误 |
