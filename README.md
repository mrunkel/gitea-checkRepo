# gitea-checkRepo
A small go binary to check if a gitea repo exists

## usage

set API_TOKEN to a gitea API Token
run `checkRepo <repoName>` 
`<repoName>`  should be in the form of "organization/repository"

checkRepo exits with 0 if the repository exists, 1 if it doesn't and 2 if there is an error.