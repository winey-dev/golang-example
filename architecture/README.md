# Architecture
- Layer를 분리하는 이유! : 관심사 분리 
- SW 설계 관점에서의 Architecture로 Golang 국한되지 않는다. 
- 여러 Architecure를 Golang으로 구현 한다.

## 3~4 Layer 
- 일반적인 Access Layer / Application Layer / DB Access Layer
## 3~4 Layer - DIP (의존 역전 원칙)
- 일반적인 Access Layer / Application Layer / DB Access Layer
- Application Layer가 DB Access Layer에 의존하는 현상을 해결하기 위해 Application Layer에서 추상화 인터페이스를 정의하고 DB Access Layer에서 이를 구현한 것
## Clean Architecture
- Uncle Bob(Robert C.Martin) 아저씨가 말하는 Clean Architecture
- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
