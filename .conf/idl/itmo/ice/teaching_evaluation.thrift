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
	1:   required string    id
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

struct QueryClassCondition {
	1: optional string    id
	2: optional string    classNumber    
	3: optional list<i64> ids            
}

struct QueryStudentClassRequest {
	1:   optional QueryClassCondition condition
	2:            i32                 pageNum
	3:            i32                 pageSize
	255: optional base.Base           Base
}

struct QueryStudentClassResponse {
	1:            i64             total        
	2:            list<ClassInfo> classList    
	255: optional base.BaseResp   BaseResp     
}

struct ClassInfo {
	1:  string id
	2:  string classNumber    
	3:  i64    createAt       
}

struct DeleteStudentClassRequest {
	1:            string    id
	255: optional base.Base Base
}

struct DeleteStudentClassResponse {
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
	1:          string studentNumber    
	2:          string studentName      
	3:          Gender gender           
	4:          string classNumber      
	5:          Major  major            
	6:          i8     grade            
	7: optional i64    id               
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

struct QueryStudentCondition {
	1: optional i64       id         
	2: optional list<i64> idList     
	3: optional string    name       
	4: optional string    number     
	5: optional i64       classId    
	6: optional Major     major      
	7: optional i8        grade      
}

struct QueryStudentRequest {
	1:            QueryStudentCondition queryStudentCondition    
	2:            i32                   pageNum                  
	3:            i32                   pageSize                 
	255: optional base.Base             Base                     
}

struct QueryStudentResponse {
	1:            i64               total              
	2:            list<StudentInfo> studentInfoList    
	255: optional base.BaseResp     BaseResp           
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

struct GetCurrentUserRequest {
	255: optional base.Base Base    
}

struct GetCurrentUserResponse {
	1:            UserInfo      userInfo    
	255: optional base.BaseResp BaseResp    
}

service TeachingEvaluationService {
    PingResponse Ping(1: PingRequest req) (api.post="/api/v1/itmo/teaching/evaluation/ping")
    
    /**  user login  **/
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/api/v1/itmo/teaching/evaluation/user/login")
    GetCurrentUserResponse GetCurrentUser(1: GetCurrentUserRequest req) (api.post="/api/v1/itmo/teaching/evaluation/user/current")

    /** student class  **/
    StudentClassCreateResponse CreateStudentClass(1: StudentClassCreateRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/class/create")
    StudentClassEditResponse EditStudentClass(1: StudentClassEditRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/class/edit")
    BatchCreateStudentClassResponse BatchCreateStudentClass(1: BatchCreateStudentClassRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/class/create/batch")
    QueryStudentClassResponse QueryStudentClass(1: QueryStudentClassRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/class/query")
    DeleteStudentClassResponse DeleteStudentClass(1: DeleteStudentClassRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/class/delete")

    /** student   **/
    CreateStudentResponse CreateStudent(1: CreateStudentRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/create")
    BatchCreateStudentResponse BatchCreateStudent(1: BatchCreateStudentRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/create/batch")
    EditStudentResponse EditStudent(1: EditStudentRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/edit")
    QueryStudentResponse QueryStudent(1: QueryStudentRequest req) (api.post="/api/v1/itmo/teaching/evaluation/admin/student/query")
}