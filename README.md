# grpc-task-app

```zsh
# buf Linter と Formatter を実行する
buf lint
buf format -w
```

```zsh
buf generate
```

```zsh
sqlc generate
```

```zsh
sudo apt-get update

sudo apt-get install openssl

openssl genrsa -out private_key.pem 2048

chmod 400 private_key.pem

openssl ecparam -name prime256v1 -genkey -noout -out private_key.pem
```

```zsh
curl --header "Content-Type: application/json" \
--data '{"email": "testuser1@example.com", "password": "pass"}' \
http://localhost:8080/proto.auth.v1.AuthService/Login

grpcurl -plaintext \
  -H "Authorization: Bearer ★tokenをセット★" \
  -d '{"title":"食料品", "description":"食料品の説明", "status":"pending", "due_date":""}' \
  localhost:8080 proto.task.v1.TaskService.CreateTask

grpcurl -plaintext \
  -H "Authorization: Bearer ★tokenをセット★" \
  -d '{"title":"食料品-リンゴ", "description":"食料品の説明", "status":"pending", "due_date":"2025-03-10T00:00:00Z"}' \
  localhost:8080 proto.task.v1.TaskService.CreateTask

grpcurl -plaintext \
  -H "Authorization: Bearer eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkEyNTZHQ00ifQ.FGe98EeAg3XQlqC0GuOciCw1bTAAZLinddv_TL71WSlXS1o0eK1xUEXkfBCJ1NQowhxm_PS15EEResuSwD6MeV1V7sfwIj-h8Bvvc-M6xcUy4LVC0l2kZKz2Iqa4uiPwF4nhCpKKmyc9ELeyii0-geGmIhMMPuYES43qu4I-AcH2hcDGlMkbNTyen7R5CTv_0kjx1NEtPwJFIPIiLmc3nhFfsTCMz7t1QDS5hjJ3_LzVqbVGnZVdsYFkQJ0vRzOG3be3Jn03piS9gmzxNMPDX40IyEG7u1RK3nlnnrz4c220UJ3BNUXjP8FY4FhgN4CNAnv81x9VXmwgh-dk740xew._ZghiRxvwFSAcYzp.-mrL5eJ4AtM1XPJ-JXWl3iPsUW5bLiGQ_2esKj83mOuu4vwj0YhdXAlJSUslp06IdTIaV1jJ2lf9IwBszrrYl6XwPcHyIYXASNnbXdBuoiTjuMHd0CBYQ7xeD4V6kiELaA6XeXSj-poI5mrpZqKuFPEjM4RloteMP8Rr4KrKPGzv_TMKFA-CqeltzIZHJK0h5mO2D0_sSusUdRUEN0689LNs7ITfNxUt_HCMIiV3a1qjEtxTB7ex2bo2yMTcpCH6Ijaq7LnvN2tvjeo98FUqSBMTCxOd0l-xIYYRY0p1HVlAGcusbxzxanpFaVKze4wITEaF-7YziDimq9QpwT_OH3_gRGb2Cm3sZr0gGwAEylgLwHjAllud3Zh2Dkwpt3I4jUbjvhhUBeyKEI3XTL7wNfnT2v9qw6C44WKF4fD5wytHiKwtFTxiJiV04EKoBjdp5MM01fGnJOdnkulQsiLxgaHpTe9uS8B-1ndwIyuNAMMMXdQAndrauZUO4AW8d2JSC86ZIZh6DZXoly7OLw4D9av3VJv11vYfPSTO_vByYBsMovapiRh6mVqJzcXHOB68oHlTvcT0DFf2yCNPlL6rtBBqZJeLS3cCj-uvPWeJrtkrkowgveZjb3qhyA.AlloNCGBZy1fcQ1LbcYN-A" \
  localhost:8080 proto.task.v1.TaskService.GetTaskList
```

```zsh
# 実行例
/opt/grpc_task_backend# grpcurl -plaintext \
>   -H "Authorization: Bearer eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkEyNTZHQ00ifQ.FGe98EeAg3XQlqC0GuOciCw1bTAAZLinddv_TL71WSlXS1o0eK1xUEXkfBC
J1NQowhxm_PS15EEResuSwD6MeV1V7sfwIj-h8Bvvc-M6xcUy4LVC0l2kZKz2Iqa4uiPwF4nhCpKKmyc9ELeyii0-geGmIhMMPuYES43qu4I-AcH2hcDGlMkbNTyen7R5CTv_0kj
x1NEtPwJFIPIiLmc3nhFfsTCMz7t1QDS5hjJ3_LzVqbVGnZVdsYFkQJ0vRzOG3be3Jn03piS9gmzxNMPDX40IyEG7u1RK3nlnnrz4c220UJ3BNUXjP8FY4FhgN4CNAnv81x9VXmw
gh-dk740xew._ZghiRxvwFSAcYzp.-mrL5eJ4AtM1XPJ-JXWl3iPsUW5bLiGQ_2esKj83mOuu4vwj0YhdXAlJSUslp06IdTIaV1jJ2lf9IwBszrrYl6XwPcHyIYXASNnbXdBuoiT
juMHd0CBYQ7xeD4V6kiELaA6XeXSj-poI5mrpZqKuFPEjM4RloteMP8Rr4KrKPGzv_TMKFA-CqeltzIZHJK0h5mO2D0_sSusUdRUEN0689LNs7ITfNxUt_HCMIiV3a1qjEtxTB7e
x2bo2yMTcpCH6Ijaq7LnvN2tvjeo98FUqSBMTCxOd0l-xIYYRY0p1HVlAGcusbxzxanpFaVKze4wITEaF-7YziDimq9QpwT_OH3_gRGb2Cm3sZr0gGwAEylgLwHjAllud3Zh2Dkw
pt3I4jUbjvhhUBeyKEI3XTL7wNfnT2v9qw6C44WKF4fD5wytHiKwtFTxiJiV04EKoBjdp5MM01fGnJOdnkulQsiLxgaHpTe9uS8B-1ndwIyuNAMMMXdQAndrauZUO4AW8d2JSC86
ZIZh6DZXoly7OLw4D9av3VJv11vYfPSTO_vByYBsMovapiRh6mVqJzcXHOB68oHlTvcT0DFf2yCNPlL6rtBBqZJeLS3cCj-uvPWeJrtkrkowgveZjb3qhyA.AlloNCGBZy1fcQ1L
bcYN-A" \
>   localhost:8080 proto.task.v1.TaskService.GetTaskList
{
  "tasks": [
    {
      "id": "task1",
      "userId": "user1",
      "title": "タスク 1",
      "description": "Description for Task 1",
      "status": "TASK_STATUS_IN_PROGRESS",
      "dueDate": "2025-03-10T00:00:00Z",
      "createdAt": "2025-02-17T22:04:04Z",
      "updatedAt": "2025-02-19T23:29:43Z"
    },
    {
      "id": "task4",
      "userId": "user1",
      "title": "タスク 4",
      "description": "Description for Task4",
      "status": "TASK_STATUS_IN_PROGRESS",
      "dueDate": "2025-04-01T00:00:00Z",
      "createdAt": "2025-02-17T22:04:04Z",
      "updatedAt": "2025-02-19T23:29:43Z"
    },
    {
      "id": "218cc6ad-884d-4a9e-93bd-84b9a12ca428",
      "userId": "user1",
      "title": "食料品",
      "description": "食料品の説明",
      "status": "TASK_STATUS_IN_PROGRESS",
      "createdAt": "2025-02-19T23:14:04Z",
      "updatedAt": "2025-02-19T23:29:43Z"
    },
    {
      "id": "de4305f6-ea83-46e1-b8e5-565e2c797534",
      "userId": "user1",
      "title": "食料品-かぼちゃ",
      "description": "食料品の説明",
      "status": "TASK_STATUS_IN_PROGRESS",
      "createdAt": "2025-02-19T23:14:46Z",
      "updatedAt": "2025-02-19T23:29:43Z"
    },
    {
      "id": "d182a7c3-1785-4d2f-8523-71307ac22a84",
      "userId": "user1",
      "title": "食料品-リンゴ",
      "description": "食料品の説明",
      "status": "TASK_STATUS_IN_PROGRESS",
      "dueDate": "2025-03-10T00:00:00Z",
      "createdAt": "2025-02-19T23:17:15Z",
      "updatedAt": "2025-02-19T23:29:43Z"
    },
    {
      "id": "task2",
      "userId": "user1",
      "title": "タスク 2",
      "description": "Description for Task 2",
      "status": "TASK_STATUS_IN_PROGRESS",
      "dueDate": "2025-03-15T00:00:00Z",
      "createdAt": "2025-02-17T22:04:04Z",
      "updatedAt": "2025-02-17T22:04:04Z"
    }
  ]
}
```