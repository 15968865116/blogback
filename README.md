# blogback
blog后端
# 12.16新增
## 新增category：
* 新增文章时需要新增category字段（完成）
* 新增category时需要判断category是否已存在，需要对category进行一个查询。（直接使用插入，若是失败则说明可能已存在）
* 在对文章进行初步获取时需要对文章按category分类（可以由前端完成）
* 文章的字段也需要添加一个category字段，字段还是用categoryID（完成）	
* category的增删查改（完成）
* 文章修改时也需要对category进行修改（完成）
## 需要测试：
* 修改category后，立刻进行文章的category修改，看是否会成功
* 前端添加相应功能
