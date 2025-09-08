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
	1:            i64           id          
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

struct BatchCreateStudentClassRequest {
	1:            list<string> classNumberList    
	255: optional base.Base    Base               
}

struct BatchCreateStudentClassResponse {
	1:            i32           num         
	255: optional base.BaseResp BaseResp    
}

/** student  **/

enum Major {
	Computer   = 0
	Automation = 1
}

enum Gender {
	Female = 0
	Male   = 1
}

struct CreateStudentRequest {
	1:            string    studentNumber    
	2:            string    studentName      
	3:            Gender    gender           
	4:            string    classNumber      
	5:            Major     major            
	6:            i8        grade            
	255: optional base.Base Base             
}

struct CreateStudentResponse {
	1:            i64           id          
	255: optional base.BaseResp BaseResp    
}


service TeachingEvaluationService {
    PingResponse Ping(1: PingRequest req) (api.post="/api/v1/itmo/teaching/evaluation/ping")
    
    /** student class  **/
    StudentClassCreateResponse CreateStudentClass(1: StudentClassCreateRequest req) (api.post="/api/v1/itmo/teaching/evaluation/student/class/create")
    StudentClassEditResponse EditStudentClass(1: StudentClassEditRequest req) (api.post="/api/v1/itmo/teaching/evaluation/student/class/edit")
    BatchCreateStudentClassResponse BatchCreateStudentClass(1: BatchCreateStudentClassRequest req) (api.post="/api/v1/itmo/teaching/evaluation/student/class/create/batch")

    /** student   **/
    CreateStudentResponse CreateStudent(1: CreateStudentRequest req) (api.post="/api/v1/itmo/teaching/evaluation/student/create")
}