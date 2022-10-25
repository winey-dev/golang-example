# GRPC-K8S 
GRPC를 통해 각 Mirco Service들이 통신을 한다.
이때 Dst(Server)의 Scale In/Out을 통해 삭제 또는 생성된 Pod에게는 GRPC Client가 어떤방식으로 동작하는지 
예제 코드를 작성하려고 한다.

# 시험 방식 
Server POD 3개 기동 후  
Client POD 1개 기동 

1. Client POD는 주기적으로 RPC 호출하여 Replication된 POD들에게 정상적으로 메시질을 수신 받는지 확인
2. Server의 POD 하나를 Delete
3. Client POD에서 새로 Recover된 Server POD에게도 RPC를 호출하는지 확인


