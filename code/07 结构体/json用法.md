Go 的 `encoding/json` 包提供了许多函数和类型，用于 JSON 编码（序列化）和解码（反序列化）。以下是一些常见的用法和解释：

### 1. `json.Marshal` 和 `json.MarshalIndent`
- **`json.Marshal`**：将 Go 数据结构转换为 JSON 字节切片。
  ```go
  bytes, err := json.Marshal(data)
  ```

- **`json.MarshalIndent`**：与 `Marshal` 类似，但可以生成带缩进的 JSON 字符串，方便阅读。
  ```go
  bytes, err := json.MarshalIndent(data, "", "  ")
  ```
参数 `""` 是前缀，`"  "` 是缩进字符，表示每一级嵌套使用两个空格缩进。

**使用 json.Marshal**

调用 json.Marshal 时，传入一个变量（如结构体、切片、映射等），它会返回两个结果：

JSON 字节切片，表示编码后的 JSON 数据。

错误信息，若编码过程中出现错误，则返回错误信息，否则返回 nil。
```go
bytes, err := json.Marshal(spirit)
 ```
### 2. `json.Unmarshal`
- **`json.Unmarshal`**：将 JSON 字节数据解析（反序列化）到 Go 数据结构中。
  ```go
  err := json.Unmarshal(bytes, &data)
  ```
  - `&data` 表示数据结构的指针，`Unmarshal` 会把 JSON 中的内容填充到该结构中。
  - 支持的类型有结构体、映射、切片等。

#### 示例
```go
var data struct {
    Name string
    Age  int
}
err := json.Unmarshal([]byte(`{"Name":"Alice","Age":25}`), &data)
```

### 3. JSON 标签
- 使用结构体标签来控制字段名、忽略字段或设定字段编码方式：
  ```go
  type Person struct {
      Name     string `json:"name"`       // 字段名映射为 `name`
      Age      int    `json:"age"`        // 字段名映射为 `age`
      Password string `json:"-"`          // 跳过此字段，不包含在 JSON 中
      Height   int    `json:"height,omitempty"` // 为空时跳过
  }
  ```

### 4. `json.NewEncoder` 和 `json.NewDecoder`
- **`json.NewEncoder`**：将 JSON 数据写入 `io.Writer`，通常用于响应 HTTP 请求时输出 JSON。
  ```go
  err := json.NewEncoder(writer).Encode(data)
  ```

- **`json.NewDecoder`**：从 `io.Reader` 中读取 JSON 数据并解码，通常用于从请求体中读取 JSON。
  ```go
  err := json.NewDecoder(reader).Decode(&data)
  ```

### 5. `RawMessage`
- **`json.RawMessage`**：允许延迟解码或部分解码 JSON。它保存 JSON 编码的数据，可在需要时进一步解码。
  ```go
  var msg json.RawMessage
  err := json.Unmarshal(data, &msg)
  ```

### 6. 处理嵌套数据
- 对于嵌套数据，可以定义嵌套的结构体，或使用 `map[string]interface{}` 来处理动态结构。

  ```go
  var result map[string]interface{}
  json.Unmarshal([]byte(`{"name": "Alice", "address": {"city": "Wonderland"}}`), &result)
  ```

  `interface{}` 会根据 JSON 类型选择合适的 Go 类型：`float64` 数字，`string` 字符串，`map[string]interface{}` 对象，`[]interface{}` 数组。

### 示例代码
```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
}

func main() {
    user := User{ID: 1, Username: "Alice"}

    // 1. Marshal to JSON
    data, _ := json.Marshal(user)
    fmt.Println(string(data))

    // 2. Unmarshal JSON to Go struct
    var newUser User
    json.Unmarshal(data, &newUser)
    fmt.Println(newUser)

    // 3. Write JSON to file using Encoder
    file, _ := os.Create("user.json")
    defer file.Close()
    json.NewEncoder(file).Encode(user)
}
```