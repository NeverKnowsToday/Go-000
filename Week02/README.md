学习笔记

学习笔记

#wrap基础

- 包含根因err，同时其中的withMessage结构体包含堆栈信息withStack（报错位置）
- 在应用代码中，使用errors.New()、errors.Errorf，如果是调用包内函数，通常直接返回
- 调用其他服务返回的错误要直接透穿，若再调用wrap会产生两倍的堆栈信息，直接返回错误，不应该到处打印日志
- 调用标准库、github、公司基础库，应该用warp保存根因，考虑使用error.warp()， error.warpf()保存堆栈信息
- 调用数据库产生的错误应该用warp，跟底层程序交互时使用呢warp，直接返回错误
- 在程序的顶部，或者工作的goroune顶部（请求入口），使用%+v把堆栈详情记录
- 使用error.Cause()获取根因，再跟sentinel error判定

#总结
- 制作基础库不能warp
- 若这层不准备处理错误，warp携带错误信息向上抛
- 打印入参和error解决99。9的问题
- 若这层处理过error，该错误不应该向上抛，返回nil

##fmt.Errorf存在问题
- 会丢弃原始错误中除文本外的所有内容

#go1.13 errors 和 fmt
标准库包引入新特性，以简化处理包含其他错误的错误

##errors
- errors.Is(err, ErrNotFound)   会对err递归找出根因，会查询err结构体是否实现了func(e *QueryError) Unwarp() error {return e.Err}方法,实现了就递归找根因。
- errors.As(err, &e)  递归找到err根因，赋给e

##fmt
- %v 丢根因
- %w 根因

