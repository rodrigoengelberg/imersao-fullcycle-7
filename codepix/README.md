# Codefix

Microserviço desenvolvido em GO utilizando Docker, comunicando com Postgres e interface de comunicação via gRPC.

## Comando para gerar protifiles gRPC
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto