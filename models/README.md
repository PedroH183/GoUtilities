## Features

- **GET Requests**: Easily make HTTP GET requests with custom headers.
- **POST Requests**: Send HTTP POST requests with JSON payloads and custom headers.
- **Error Handling**: Handles HTTP response errors and decodes JSON responses into Go structs.

## Installation

To use this module in your project, run:

```bash
go get github.com/PedroH183/GoUtilities@latest
```



## Usage

### Import the Package

```go
import "github.com/PedroH183/GoUtilities/models"
```

### Example: GET Request

```go
package main

import (
	"fmt"
	"github.com/PedroH183/GoUtilities/models"
	"net/http"
)

func main() {
	client := &models.HttpClient{
		Client: &http.Client{},
	}
	// ou uma struct qualquer
	var response interface{} 

	headers := map[string]string{
		"Authorization": "Bearer your_token",
	}

	err := client.MakeGetRequest("https://api.example.com/data", headers, &response)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response)
}
```

### Example: POST Request

```go
package main

import (
	"fmt"
	"github.com/PedroH183/GoUtilities/models"
	"net/http"
)

func main() {
	client := &models.HttpClient{
		Client: &http.Client{},
	}
	payload := map[string]interface{}{
		"key": "value",
	}
	headers := map[string]string{
		"Authorization": "Bearer your_token",
	}

	// ou uma struct qualquer
	var response interface{} 

	err := client.MakePostRequest("https://api.example.com/data", headers, &response, payload)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response)
}
```
