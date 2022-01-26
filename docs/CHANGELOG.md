# CHANGELOG

## 1.0.2 (January 2022)

### Code Enhancements:

* renamed APIClient's field `TokenApi` to `TokensApi`, pointer to `TokensApiService`
* renamed function `GetAllTokens()` to `TokensGet()`
* renamed function `GetAllTokensExecute()` to `TokensGetExecute()`
* renamed function `GetTokenById()` to `TokensFindById()`
* renamed function `GetTokenByIdExecute()` to `TokensFindByIdExecute()`
* renamed function `CreateToken()` to `TokensGenerate()`
* renamed function `CreateTokenExecute()` to `TokensGenerateExecute()`
* renamed function `DeleteTokenById()` to `TokensDeleteById()`
* renamed function `DeleteTokenByIdExecute()` to `TokensDeleteByIdExecute()`
* renamed function `DeleteTokenByCriteria()` to `TokensDeleteByCriteria()`
* renamed function `DeleteTokenByCriteriaExecute()` to `TokensDeleteByCriteriaExecute()`

### Fixes:

* fixed bug to support both `https` and `http` schemas

### Dependency-updates:

* updated `golang.org/x/oauth2` go import to the latest version

## 1.0.1 (November, 2021)

### Code Enhancements:

* sorted imports
* removed `git_push.sh` script

## 1.0.0 (November, 2021)

### Features:

* first release 🎉
* added api(`TokenApi`) and models(`Jwt`, `Token`, `Tokens`, etc.) for Auth API
