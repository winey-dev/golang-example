# disk-use-by-bytes 는 통계를 기록하는 시점에서 현재 값을 기록한다

/*
    foo, bar 각 프로젝트가 사용하는 DISK 사용량을 Node 별로 조회 하기
*/
from(bucket: "sample")
  |> range(start: -1m)
  |> filter(fn: (r) => r["_measurement"] == "foo-101" or r["_measurement"] == "foo-102" or r["_measurement"] == "bar-101" or r["_measurement"] == "bar-102")
  |> filter(fn: (r) => r["_field"] == "disk-use-by-bytes")
  |> drop(columns: ["_start","_stop","app_name","pod_name","container_name"])
  |> group(columns: ["node_name"])
  |> aggregateWindow(every: 10s, fn: sum, createEmpty: false)
  |> group()
  |> pivot(rowKey: ["_time"], columnKey: ["node_name"], valueColumn: "_value")
---
/*
    foo, bar 각 프로젝트가 사용하는 DISK 사용량을 App 별로 조회 하기
*/
from(bucket: "sample")
  |> range(start: -1m)
  |> filter(fn: (r) => r["_measurement"] == "foo-101" or r["_measurement"] == "foo-102" or r["_measurement"] == "bar-101" or r["_measurement"] == "bar-102")
  |> filter(fn: (r) => r["_field"] == "disk-use-by-bytes")
  |> drop(columns: ["_start","_stop","node_name","pod_name","container_name"])
  |> group(columns: ["app_name"])
  |> aggregateWindow(every: 10s, fn: sum, createEmpty: false)
  |> group()
  |> pivot(rowKey: ["_time"], columnKey: ["app_name"], valueColumn: "_value")

---
/*
   아래 3개의 QUERY는 multi feild를 Query후 데이터를 연산 하는 방법
*/
from(bucket: "sample")
  |> range(start: -1m)
  |> filter(fn: (r) => r["_measurement"] == "bar-101")
  |> filter(fn: (r) => r["app_name"] == "alert")
  |> filter(fn: (r) => r["pod_name"] == "alert-0")
  |> filter(fn: (r) => r["node_name"] == "worker-2")
  |> filter(fn: (r) => r["container_name"] == "alert")
  |> filter(fn: (r) => r["_field"] == "response_succ" or r["_field"] == "request_count")
---
from(bucket: "sample")
  |> range(start: -1m)
  |> filter(fn: (r) => r["_measurement"] == "bar-101")
  |> filter(fn: (r) => r["app_name"] == "alert")
  |> filter(fn: (r) => r["pod_name"] == "alert-0")
  |> filter(fn: (r) => r["node_name"] == "worker-2")
  |> filter(fn: (r) => r["container_name"] == "alert")
  |> filter(fn: (r) => r["_field"] == "response_succ" or r["_field"] == "request_count")
  |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> drop(columns: ["_start", "_stop"])
--- 
from(bucket: "sample")
  |> range(start: -1m)
  |> filter(fn: (r) => r["_measurement"] == "bar-101")
  |> filter(fn: (r) => r["app_name"] == "alert")
  |> filter(fn: (r) => r["pod_name"] == "alert-0")
  |> filter(fn: (r) => r["node_name"] == "worker-2")
  |> filter(fn: (r) => r["container_name"] == "alert")
  |> filter(fn: (r) => r["_field"] == "response_succ" or r["_field"] == "request_count")
  |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> drop(columns: ["_start", "_stop"])
  |> map(fn: (r) => ({r with _value: (r.response_succ / r.request_count)*100.000}))
---
/*
   모든 namespace/app/pod/conatiner들의 계산 결과
 */
from(bucket: "sample")
  |> range(start: -1m)
  |> filter(fn: (r) => r["_field"] == "response_succ" or r["_field"] == "request_count")
  |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> drop(columns: ["_start", "_stop"])
  |> map(fn: (r) => ({r with _value: (r.response_succ / r.request_count)*100.000}))
---
/*
  app 별 계산 결과
  app으로 group 지어 같은 테이블에 위치 시킨 후 각 계산값의 평균값을 구하기
 */
from(bucket: "sample")
  |> range(start: -1m)
  |> filter(fn: (r) => r["_field"] == "response_succ" or r["_field"] == "request_count")
  |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> drop(columns: ["_start", "_stop"])
  |> map(fn: (r) => ({r with _value: (r.response_succ / r.request_count)*100.000}))
  |> group(columns:["app_name"])
  |> aggregateWindow(every: 10s, fn: mean, createEmpty: false)
---
/*
  app 별 계산 결과
  위 결과를 같은 시간대 별로 각 app 별로 확인하기 
 */
from(bucket: "sample")
  |> range(start: -1m)
  |> filter(fn: (r) => r["_field"] == "response_succ" or r["_field"] == "request_count")
  |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> drop(columns: ["_start", "_stop"])
  |> map(fn: (r) => ({r with _value: (r.response_succ / r.request_count)*100.000}))
  |> group(columns:["app_name"])
  |> aggregateWindow(every: 10s, fn: mean, createEmpty: false)
  |> pivot(rowKey: ["_time"], columnKey: ["app_name"], valueColumn: "_value")
