# golang_json_tool
一个用于生成golang结构体胶水代码，已经从json生成golang代码的库

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
