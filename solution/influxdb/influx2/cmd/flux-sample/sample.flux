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

from(bucket: "sample")
  |> range(start: -1m)
  |> filter(fn: (r) => r["_measurement"] == "foo-102" or r["_measurement"] == "foo-101" or r["_measurement"] == "bar-102" or r["_measurement"] == "bar-101")
  |> filter(fn: (r) => r["_field"] == "response_count" or r["_field"] == "response_succ" or r["_field"] == "response_fail")
  |> aggregateWindow(every: 10s, fn: mean, createEmpty: false)
  |> yield(name: "mean")
