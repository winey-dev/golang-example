# Functional Options Pattern 
- 0개 이상의 함수를 인수로 받아들여 가변 생성자를 사용하여 복잡한 구조체를
  빌드 할 수있는 생성 디자인 패턴
- 도메인 객체를 생성할 때 Builder Pattern보다 좀더 유용하게 사용 할 수 있다.

# 단점
- 각각의 옵션에 대해 함수를 구현해줘야 한다는 단점이 존재한다. 
- 필수 필드의 경우 위치 인수로 전달 하는 것이 좋다
- 상호 연관이 있는 Option의 Validate 검사의 경우, 인수 순서에 따라 오류처리가 될 수도 있다.

# Compare 
- Builder Pattern
