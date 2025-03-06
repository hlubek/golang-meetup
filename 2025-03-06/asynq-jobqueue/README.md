# asynq jobqueue

## Run worker

go run main.go worker

## Start asynqmon

go run github.com/hibiken/asynqmon/cmd/asynqmon

## Add tasks to the queue

go run main.go queue-resize https://unsplash.com/photos/S07ipcMrVLw/download?ixid=M3wxMjA3fDB8MXxhbGx8Mjl8fHx8fHx8fDE3NDEyNjc4NDJ8&force=true S07ipcMrVLw.jpg
go run main.go queue-resize https://unsplash.com/photos/GPw17rBXxpk/download?ixid=M3wxMjA3fDB8MXxhbGx8NDN8fHx8fHx8fDE3NDEyNzE0Mjl8&force=true GPw17rBXxpk.jpg
go run main.go queue-resize https://unsplash.com/photos/3TOcMH5MY0Q/download?ixid=M3wxMjA3fDB8MXxhbGx8fHx8fHx8fHwxNzQxMjc0NDg5fA&force=true 3TOcMH5MY0Q.jpg
go run main.go queue-resize https://unsplash.com/photos/Cwd-Qca1fAM/download?ixid=M3wxMjA3fDB8MXxhbGx8NTd8fHx8fHx8fDE3NDEyNzMyOTl8&force=true Cwd-Qca1fAM.jpg
go run main.go queue-resize https://unsplash.com/photos/kHvfT1cCU1U/download?ixid=M3wxMjA3fDB8MXxhbGx8ODJ8fHx8fHx8fDE3NDEyNzMzMTJ8&force=true kHvfT1cCU1U.jpg
