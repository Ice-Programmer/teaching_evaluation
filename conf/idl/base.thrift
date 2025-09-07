namespace py base
namespace go base
namespace rs base
namespace cpp base
namespace java com.bytedance.thrift.base

struct TrafficEnv {
    1: bool Open = false, // A boolean value indicating whether the traffic environment is open.
    2: string Env = "", // A string representing the name of the traffic environment.
}

struct Base {
    1: string LogID = "", // A string representing the log ID.
    2: string Caller = "", // A string representing the caller.
    3: string Addr = "", // A string representing the address.
    4: string Client = "", // A string representing the client.
    5: optional TrafficEnv TrafficEnv, // An optional TrafficEnv struct representing the traffic environment.
    6: optional map<string, string> Extra, // An optional map of strings to strings representing extra information.
}

struct BaseResp {
    1: string StatusMessage = "", // A string representing the status message.
    2: i32 StatusCode = 0, // An integer representing the status code.
    3: optional map<string, string> Extra, // An optional map of strings to strings representing extra information.
}