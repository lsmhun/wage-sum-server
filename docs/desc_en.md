# WageSum demo application 

golang wage sum microservice

## Links
https://github.com/golang-standards/project-layout 

## Openapi
```shell
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/api/wagesum-openapi.yaml  -g go-server  -o /local/internal/pkg --additional-properties=noservice,enumClassPrefix=true,featureCORS=false,onlyInterfaces,outputAsLibrary=true,sourceFolder=openapi
```

https://medium.com/@ankur_anand/how-to-mock-in-your-go-golang-tests-b9eee7d7c266

https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/ 

https://swagger.io/docs/specification/authentication/oauth2/

## Cleanup 
```shell
go mod tidy  
go mod vendor
```

## DB
https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b

```shell
docker run -p 5432:5432 --name wagesum-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```


## Other
https://towardsdev.com/golang-productivity-hacks-part-3-auto-generating-test-4c8055dc7946
https://eli.thegreenplace.net/2021/rest-servers-in-go-part-4-using-openapi-and-swagger/

https://stackoverflow.com/questions/7106012/download-a-single-folder-or-directory-from-a-github-repo

https://ribice.medium.com/serve-swaggerui-within-your-golang-application-5486748a5ed4

https://github.com/GoogleCloudPlatform/golang-samples