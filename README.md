# go-common
Core packages for micro service architecture

## About

MSA 환경에서 각 도메인 개발자들의 작업이 매번 필요한 Boilerplate 를 최대한 줄이고, 공통 로직의 재사용성을 높이고자 합니다. <br/>
그 동안 각 도메인마다 개발했던 모듈 중 공통 모듈을 한데 모아 공통 모듈로 의존성을 관리합니다.

## Examples

### ✨ rest pkg
#### Succeed
	response := rest.NewApiResponse()
	response.Succeed("", base64Encoding)

#### Error
	response := rest.NewApiResponse()
	response.Error(&rest.BAD_REQUEST_ERROR, err.Error())

#### CustomError
	&rest.CustomError{CodeDesc: &rest.ACCESS_DENIED_ACCOUNT_DISABLE, Message: "정보보안 담장자에게 문의하세요."}
	
