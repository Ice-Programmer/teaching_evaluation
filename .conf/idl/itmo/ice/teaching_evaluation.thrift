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

struct StudentClassEditRequest {
	1:   required i64       id             
	2:            string    classNumber    
	255: optional base.Base Base           
}

struct StudentClassEditResponse {
	255: optional base.BaseResp BaseResp    
}

struct BatchCreateStudentRequest {
	1:            list<string> classNumberList    
	255: optional base.Base    Base               
}

struct BatchCreateStudentResponse {
	1:            i32           num         
	255: optional base.BaseResp BaseResp    
}


service TeachingEvaluationService {
    PingResponse Ping(1: PingRequest req) (api.post="/api/v1/itmo/teaching/evaluation/ping")
    
    /** student class  **/
    StudentClassCreateResponse CreateStudentClass(1: StudentClassCreateRequest req) (api.post="/api/v1/itmo/teaching/evaluation/student/class/create")
    StudentClassEditResponse EditStudentClass(1: StudentClassEditRequest req) (api.post="/api/v1/itmo/teaching/evaluation/student/class/edit")
    BatchCreateStudentResponse BatchCreateStudentClass(1: BatchCreateStudentRequest req) (api.post="/api/v1/itmo/teaching/evaluation/student/class/create/batch")
}