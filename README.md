# wage-sum-server
golang wage sum microservice

# Links
https://github.com/golang-standards/project-layout 

# Openapi
```shell
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/api/wagesum-openapi.yaml  -g go-server  -o /local/internal/pkg --additional-properties=noservice,enumClassPrefix=true,featureCORS=false,onlyInterfaces,outputAsLibrary=true,sourceFolder=openapi
```

https://medium.com/@ankur_anand/how-to-mock-in-your-go-golang-tests-b9eee7d7c266

https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/ 

https://swagger.io/docs/specification/authentication/oauth2/

# Cleanup 
```shell
go mod tidy  
go mod vendor
```

# Other
https://towardsdev.com/golang-productivity-hacks-part-3-auto-generating-test-4c8055dc7946