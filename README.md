## openaigo

A Go client library for the OpenAI Completions API.


### Installation

To install the `openaigo` package, run the following command:

```
go get github.com/immadiuttej/openaigo
```

### Usage

Here is an example of how to use the `Complete` function to generate completions for a given prompt:

```
package main

import (
    "context"
    "fmt"
    "net/http"
     openai "github.com/immadiuttej/openaigo"
)

func main() {
  // Create a new OpenAI client.
  client := openai.Client{
  HTTPClient: http.DefaultClient,
  APIKey:     "YOUR_OPENAI_API_TOKEN",
  Endpoint:   "https://api.openai.com/v1/completions",
  }

  req := openai.CompletionRequest{
      Prompt:      "how much food does a panda need?",
      Model:       "text-davinci-003",
      MaxTokens:   100,
      Temperature: 0.5,
  }
  resp, err := client.Complete(context.Background(), &req)
  if err != nil {
      fmt.Println(err)
      return
  }
  fmt.Println(resp)
}

```


## Response

The response contains the string in a list.
```
["\n\nA panda needs around 20 to 40 pounds of bamboo per day to meet its nutritional needs."]
```
