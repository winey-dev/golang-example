# GRPC-K8S 
GRPC를 통해 각 Mirco Service들이 통신을 한다.
이때 Dst(Server)의 Scale In/Out을 통해 삭제 또는 생성된 Pod에게는 GRPC Client가 어떤방식으로 동작하는지 
예제 코드를 작성하려고 한다.

