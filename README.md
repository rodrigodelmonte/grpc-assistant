# grpc-assistant

> grpc-assistant is a demo code to explore how to handle grpc requests, and forward it to OpenAI API

## Running

```sh
export OPENAI_API_KEY="sk-00000"
export OPENAI_MODEL_NAME="gpt-3.5-turbo"
go run main.go
```

## How to make questions

Install [grpcurl](https://github.com/fullstorydev/grpcurl)

```sh
❯ grpcurl -d '{"request": "Who has more Football World Cups?"}' -plaintext -proto assistant.proto 0.0.0.0:8080 Assistant.Handle
{
  "result": "The country with the most Football World Cups is Brazil, with a total of 5 championships (1958, 1962, 1970, 1994, and 2002)."
}
```

## Updating proto files

```sh
protoc \
--go_out=pkg/assistant \
--go_opt=paths=source_relative \
--go-grpc_out=pkg/assistant \
--go-grpc_opt=paths=source_relative \
assistant.proto
```
