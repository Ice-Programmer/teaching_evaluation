include "../../base.thrift"

namespace go teaching_evaluation

struct PingRequest {
	255: optional base.Base Base    
}

struct PingResponse {
	1:   optional string        response    
	255: optional base.BaseResp BaseResp    
}

/** student class  **/
struct StudentClassCreateRequest {
	1:            string    classNumber    
	255: optional base.Base Base
}

struct StudentClassCreateResponse {
	255: optional base.BaseResp BaseResp    
}


service TeachingEvaluationService {
    PingResponse Ping(1: PingRequest req) (api.post="/api/v1/itmo/teaching/evaluation/ping")
    
    /** student class  **/
    StudentClassCreateResponse CreateStudentClass(1: StudentClassCreateRequest req) (api.post="/api/v1/itmo/teaching/evaluation/student/class/create")
}