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

enum Status {
	NormalStatus = 0   
	BanStatus    = 1   
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

struct BatchCreateStudentRequest {
	1:            list<StudentInfo> studentList    
	255: optional base.Base         Base           
}

struct StudentInfo {
	1:  string studentNumber    
	2:  string studentName      
	3:  Gender gender           
	4:  string classNumber      
	5:  Major  major            
	6:  i8     grade            
}

struct BatchCreateStudentResponse {
	1:            i32           num         
	255: optional base.BaseResp BaseResp    
}

struct EditStudentRequest {
	1:            i64       id               
	2:            string    studentNumber    
	3:            string    studentName      
	4:            Gender    gender           
	5:            string    classNumber      
	6:            Major     major            
	7:            i8        grade            
	8:            Status    status           
	255: optional base.Base Base             
}

struct EditStudentResponse {
	255: optional base.BaseResp BaseResp    
}

/**  user login  **/
enum UserRole {
	Student = 1   
	Admin   = 2   
}

struct UserLoginRequest {
	1:            string    userAccount     
	2:            string    userPassword    
	255: optional base.Base Base            
}

struct UserInfo {
	1:  i64      id          
	2:  string   name        
	3:  UserRole role        
	4:  i64      createAt    
}

struct UserLoginResponse {
	1:            UserInfo      userInfo    
	2:            string        token       
	3:            i64           expireAt    
	255: optional base.BaseResp BaseResp    
}

service TeachingEvaluationService {
    PingResponse Ping(1: PingRequest req) (api.post="/api/v1/itmo/teaching/evaluation/ping")
    
    /**  user login  **/
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/api/v1/itmo/teaching/evaluation/user/login")
    
    /** student class  **/
    StudentClassCreateResponse CreateStudentClass(1: StudentClassCreateRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/class/create")
    StudentClassEditResponse EditStudentClass(1: StudentClassEditRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/class/edit")
    BatchCreateStudentClassResponse BatchCreateStudentClass(1: BatchCreateStudentClassRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/class/create/batch")

    /** student   **/
    CreateStudentResponse CreateStudent(1: CreateStudentRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/create")
    BatchCreateStudentResponse BatchCreateStudent(1: BatchCreateStudentRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/create/batch")
    EditStudentResponse EditStudent(1: EditStudentRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/edit")
}