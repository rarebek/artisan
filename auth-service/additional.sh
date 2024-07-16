go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


if ! grep -qxF 'export PATH=$PATH:$HOME/go/bin' ~/.bashrc; then
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
    echo 'Added PATH update to .bashrc'
else
    echo 'PATH update already exists in .bashrc'
fi


source ~/.bashrc

which protoc-gen-go
which protoc-gen-go-grpc
