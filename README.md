# wage-sum-server
golang wage sum microservice

# Links
https://github.com/golang-standards/project-layout 

```shell
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/api/wagesum-openapi.yaml  -g go-server  -o /local/internal/pkg --additional-properties=noservice,enumClassPrefix=true,featureCORS=false,onlyInterfaces,outputAsLibrary=true,sourceFolder=openapi
```

https://medium.com/@ankur_anand/how-to-mock-in-your-go-golang-tests-b9eee7d7c266