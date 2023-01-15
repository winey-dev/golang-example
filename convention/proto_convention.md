
# Interface Name
- 일반적으로는 직관적인 명사를 사용. ex) `Library`
- 이미 언어 런타임 라이브러리에서 확립된 개념과 충돌하면 안됨. `File`
- 인터페이스 이름이 충돌할 경우 서픽스를 사용 `Api` 또는 `Service`

# Field Name
- .proto 파일에서 필드를 정의할 때는 `lower_case_undersocre_separated_names`를 사용해야합니다.
- 필드 이름은 `for`, `during`,`at`등의 전치사를 포함하지 않아야합니다.
  `error_reason` -> `reason_for_error`
  `failure_time_cpu_usage` -> `cpu_usage_at_time_of_failure`
- 필드 이름에 후치 형용사를 사용하지 않아야합니다.
  `collected_items` -> `items_collected`
  `imported_object` -> `object_imported`
  
# Method Name 
- UpperCamelCase로 지정해야한다.
- 
|동사|명사|메소드|요청메시지|응답메시지|
|:----|:----|:------|:---------|:---------|
|List|Book|ListBooks|ListBooksRequest|ListBooksReseponse|
|Get|Book|GetBook|GetBookRequest|Book|
|Create|Book|CreateBook|CreateBookRequest|Book|
|Update|Book|UpdateBook|UpdateBookRequest|Book|
|Rename|Book|RenameBook|RenameBookRequest|RenameBookResponse|
|Delete|Book|DeleteBook|DeleteBookRequest|google.protobuf.Empty|

# Enums
- 열거형은 UpperCamelCase 이름을 사용해야합니다. 
- 열거형의 값은 `CAPITALIZED_NAMES_WITH_UNDERSCORES`을 사용해야합니다. 
- 예시
~~~
enum FooBar {
    FOO_BAR_UNSPECIFIED = 0;
    FIRST_VALUE = 1;
    SECOND_VALUE = 2;
}
~~~

- 만약 0이 UNSPECIFIED가 아닌 다른 의미를 갖는 경우 Wrapper 해야합니다.(proto2와 호환하기위해 ?)
~~~
enum OldEnum {
    INVALID = 0;
    VALID = 1;
}

message OldEnumValue {
    OldEnum value = 1;
}
~~~

# 반복되는 필드
- 복수형으로 표현해야 합니다. `book` -> `books`

# 시간과 지속
- 시간대 또는 캘린더와 상관없이 특정 시점을 나태내려면 `google.protobuf.Timestamp`를 사용해야합니다.
> 필드 이름은 time으로 끝나야 합니다. `start_time`, `end_time`

- 시간 활동을 나타낼때는 필드 이름이 `verb_time` 형태여야 합니다. 동사의 과거 시제를 사용하지 마세요.
`created_time` -> `create_time`
`updated_time` -> `update_time`

- 일/월 같은 캘린더 개념 없이 두 지점의 지속 시간을 나타내려면 `google.protobuf.Duration`을 사용해야합니다.
- 레거시 또는 지속 시간, 지연 시간 같은 필드를 정수 형식으로 나타내야 할 경우 필드 이름은 아래와 같은 형상이여야 합니다.
~~~
xxx_{time|duration|delay|latency}_{seconds/millis/mircors/nanos}
~~~

# 날짜와 시간
- 시간대 및 시간에 관계없는 날짜의 경우 `google.type.Date`를 사용해야 하며 서픽스 `_date`가 있어야 합니다.
- 날짜를 문자로 표현할 경우 ISO 8601 형식인 `YYYY-MM-DD`를 따라야 합니다.
- 시간대와 날짜와 관계없는 시간의 경우 `google.type.TimeOfDay`를 사용해야 하며 서픽스 `_time`이 있어야 합니다.
- 시간을 문자열로 표현할 경우 ISO 8601 형식인 `HH:MM:SS[.FFF]`를 따라야 합니다.

# 참고 자료
- [google Cloud API 디자인 가이드 문서](https://cloud.google.com/apis/design/naming_convention)
