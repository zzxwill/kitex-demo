namespace go api

struct GetMaximalRequest {
    1: list<i32> numbers
}

struct GetMaximalResponse {
    1: i32 result
}

service calculator {
    GetMaximalResponse GetMaximal(1: GetMaximalRequest request)
}
