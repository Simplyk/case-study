# zeffy-case-study

## Backend

Go 1.22.0

### Exigences

go version 1.22.0 - https://go.dev/doc/install

### Étapes

#### GO

`cd backend_go && go run server.go`

OU

`cd backend_go && go build && ./zbackend`

#### Express

`cd backend_express && yarn install && yarn dev`

root api: localhost:1323/api/v1

### Tests

`go test -coverprofile=coverage.out ./...`

## Frontend

NextJS 14

### Exigences

yarn 1.22.21 (sinon npm)

node version >=18.17.0

### Étapes

`cd zfrontend && yarn install && yarn dev`

servi sur: localhost:3000

### Tests

<table style="text-align:center">
 <tr>
    <td><b style="font-size:20px">Jest</b></td>
    <td><b style="font-size:20px">Cypress</b></td>
 </tr>
 <tr>
    <td><pre>yarn test</pre></td>
    <td><pre>yarn e2e</pre></td>
 </tr>
</table>

## Suggestions

### Real world considerations

### Autres Questions
