# golang_json_tool
由于golang是强类型语言，本产品是用于生成golang结构体之间胶水代码，主要原理是靠字段名和json tag的匹配
相比于copier没有反射的性能损失，而且可以灵活的实现一些自己的特殊逻辑
## 文档： 【腾讯文档】Golang结构体胶水代码
https://docs.qq.com/doc/DUU55emFPeFFKQVNI

例子

赋值的结构体
```
type TestModelA struct{
    A int64 `json:"a,omitempty"`
    B float32 `json:"b"`
    C model.A 
}
```
被赋值结构体
```
Type TestModelB struct{
    A int `json:"b"`
    B float32 
    C model.B
    D int64 
}
```
生成的代码
```
 func(testModelA *TestModelA, testModelB *TestModelB) { 
    TestModelB.A = int(TestModelA.A)
    TestModelB.B = TestModelA.B
    // TestModelB.C = TestModelA.C
    // TestModelB.D = 
}
```

# API
GenCopyCode(srcStr string, dstStr string) 

# 支持指针判断和赋值
- 如果拷贝对象字段是指针，那么会先判断是否是nil，然后进行赋值
- 如果需要强制转化，那会先生成一个临时对象

用例
```
type TestA struct {
    A *int64
    B float32
    C int64 
    D *u 
}

type TestB struct {
    A int 
    B float64 
    C *int 
    D u 
}
```
生成代码
```
func(testA *TestA, testB *TestB) { 
     if TestA.A!= nil {
        ATmp := int(*TestA.A)
        TestB.A = ATmp
    }

    TestB.B = float64(TestA.B)
    CTmp := int(TestA.C)
    TestB.C = &CTmp

    if TestA.D!= nil {
        TestB.D = *TestA.D
    }

}
```