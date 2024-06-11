 .
├──  cmd
│  ├──  migrator
│  │  └──  main.go
│  └──  sso
│     └──  main.go
├──  Dockerfile
├──  example.env
├──  go.mod
├──  go.sum
├──  grpc
│  └──  auth
│     ├──  attendance.go
│     ├──  code.go
│     ├──  getdata.go
│     ├──  login.go
│     ├──  logout.go
│     ├──  register.go
│     ├──  server.go
│     └──  update.go
├──  internal
│  ├──  app
│  │  ├──  app.go
│  │  └──  grpc
│  │     └──  app.go
│  ├──  config
│  │  └──  config.go
│  ├──  domain
│  │  └──  models
│  │     ├──  attendance.go
│  │     ├──  group.go
│  │     ├──  location.go
│  │     ├──  qrcode.go
│  │     ├──  schedule.go
│  │     ├──  subject.go
│  │     └──  user.go
│  ├──  lib
│  │  ├──  jwt
│  │  │  ├──  jwt.go
│  │  │  └──  logger
│  │  │     ├──  handlers
│  │  │     │  ├──  slogdiscard
│  │  │     │  │  └──  slogdiscard.go
│  │  │     │  └──  slogpretty
│  │  │     │     └──  slogpretty.go
│  │  │     └──  sl
│  │  │        └──  sl.go
│  │  └──  rnd
│  │     └──  rnd.go
│  └──  services
│     └──  auth
│        ├──  auth.go
│        ├──  getuserdata.go
│        ├──  journal.go
│        ├──  login.go
│        ├──  register.go
│        ├──  studscodes.go
│        ├──  teachrscodes.go
│        └──  update.go
├──  Makefile
├──  migrations
│  └──  surreal
│     ├──  1_init.down.surql
│     └──  1_init.up.surql
├──  protos
│  ├──  buf.gen.yaml
│  ├──  buf.lock
│  ├──  buf.yaml
│  ├──  gen
│  │  └──  go
│  │     └──  proto
│  │        └──  sso
│  │           ├──  sso.pb.go
│  │           └──  sso_grpc.pb.go
│  ├──  Makefile
│  └──  proto
│     └──  sso
│        └──  sso.proto
└──  README.md
