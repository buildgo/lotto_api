# Lotto API
Lotto 번호를 응답하는 API 요청을 받으면 생성된 Lotto 번호를 담아서 환경변수에 저장된 SLACK WEBHOOK 으로 API 요청을 한다.

- `GET /lotto`

## Learned

- `godotenv`  로 환경변수 주입하는 방법 - `err := godotenv.Load()`
- `map` 타입을 JSON 으로 변경하는 방법  - `json.Marshal(bodyMap)`
- `net/http`  로 POST 요청 보내는 방법
- 중복없이 랜덤하게  정해진 범위 안의 숫자를 추출하는 방법
- 함수 인자에 default 값을 사용하는 방법
- 숫자 배열을 문자로 변환하는 방법 -  ex {1, 2, 3} →  “1, 2, 3”
- 숫자 배열 정렬하는 방법
- 어떤 값이 배열에 포함된 원소인지 확인하는 방법 - `slices.Contains(xs, e)`