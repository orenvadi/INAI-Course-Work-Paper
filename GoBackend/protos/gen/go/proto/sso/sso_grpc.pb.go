// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/sso/sso.proto

package sso

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// ------------------------------------------------------------//
	//
	//	REFACTOR OLD ONES                      //
	//
	// ------------------------------------------------------------//
	RegisterTeacher(ctx context.Context, in *RegisterTeacherRequest, opts ...grpc.CallOption) (*RegisterTeacherResponse, error)
	RegisterStudent(ctx context.Context, in *RegisterStudentRequest, opts ...grpc.CallOption) (*RegisterStudentResponse, error)
	LoginTeacher(ctx context.Context, in *LoginTeacherRequest, opts ...grpc.CallOption) (*LoginTeacherResponse, error)
	LoginStudent(ctx context.Context, in *LoginStudentRequest, opts ...grpc.CallOption) (*LoginStudentResponse, error)
	UpdateTeacherInfo(ctx context.Context, in *UpdateTeacherInfoRequest, opts ...grpc.CallOption) (*UpdateTeacherInfoResponse, error)
	UpdateStudentInfo(ctx context.Context, in *UpdateStudentInfoRequest, opts ...grpc.CallOption) (*UpdateStudentInfoResponse, error)
	LogoutTeacher(ctx context.Context, in *LogoutTeacherRequest, opts ...grpc.CallOption) (*LogoutTeacherResponse, error)
	LogoutStudent(ctx context.Context, in *LogoutStudentRequest, opts ...grpc.CallOption) (*LogoutStudentResponse, error)
	GetTeacherProfileData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTeacherProfileDataResponse, error)
	GetStudentProfileData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentProfileDataResponse, error)
	// at the end of lesson
	GetConfirmCode(ctx context.Context, in *GetConfirmCodeRequest, opts ...grpc.CallOption) (Auth_GetConfirmCodeClient, error)
	GetAttendanceJournal(ctx context.Context, in *GetAttendanceJournalRequest, opts ...grpc.CallOption) (*GetAttendanceJournalResponse, error)
	// ----------------STUDENTS
	// both at the start and the end of lesson
	SubmitTeacherCode(ctx context.Context, in *SubmitCodeRequest, opts ...grpc.CallOption) (*SubmitCodeResponse, error)
	SubmitRoomCode(ctx context.Context, in *SubmitCodeRequest, opts ...grpc.CallOption) (*SubmitCodeResponse, error)
	GetAttendanceLessons(ctx context.Context, in *GetAttendanceLessonsRequest, opts ...grpc.CallOption) (*GetAttendanceLessonsResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) RegisterTeacher(ctx context.Context, in *RegisterTeacherRequest, opts ...grpc.CallOption) (*RegisterTeacherResponse, error) {
	out := new(RegisterTeacherResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RegisterTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RegisterStudent(ctx context.Context, in *RegisterStudentRequest, opts ...grpc.CallOption) (*RegisterStudentResponse, error) {
	out := new(RegisterStudentResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RegisterStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginTeacher(ctx context.Context, in *LoginTeacherRequest, opts ...grpc.CallOption) (*LoginTeacherResponse, error) {
	out := new(LoginTeacherResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/LoginTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginStudent(ctx context.Context, in *LoginStudentRequest, opts ...grpc.CallOption) (*LoginStudentResponse, error) {
	out := new(LoginStudentResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/LoginStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateTeacherInfo(ctx context.Context, in *UpdateTeacherInfoRequest, opts ...grpc.CallOption) (*UpdateTeacherInfoResponse, error) {
	out := new(UpdateTeacherInfoResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/UpdateTeacherInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateStudentInfo(ctx context.Context, in *UpdateStudentInfoRequest, opts ...grpc.CallOption) (*UpdateStudentInfoResponse, error) {
	out := new(UpdateStudentInfoResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/UpdateStudentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LogoutTeacher(ctx context.Context, in *LogoutTeacherRequest, opts ...grpc.CallOption) (*LogoutTeacherResponse, error) {
	out := new(LogoutTeacherResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/LogoutTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LogoutStudent(ctx context.Context, in *LogoutStudentRequest, opts ...grpc.CallOption) (*LogoutStudentResponse, error) {
	out := new(LogoutStudentResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/LogoutStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetTeacherProfileData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTeacherProfileDataResponse, error) {
	out := new(GetTeacherProfileDataResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetTeacherProfileData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetStudentProfileData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStudentProfileDataResponse, error) {
	out := new(GetStudentProfileDataResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetStudentProfileData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetConfirmCode(ctx context.Context, in *GetConfirmCodeRequest, opts ...grpc.CallOption) (Auth_GetConfirmCodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Auth_ServiceDesc.Streams[0], "/auth.Auth/GetConfirmCode", opts...)
	if err != nil {
		return nil, err
	}
	x := &authGetConfirmCodeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Auth_GetConfirmCodeClient interface {
	Recv() (*GetConfirmCodeResponse, error)
	grpc.ClientStream
}

type authGetConfirmCodeClient struct {
	grpc.ClientStream
}

func (x *authGetConfirmCodeClient) Recv() (*GetConfirmCodeResponse, error) {
	m := new(GetConfirmCodeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authClient) GetAttendanceJournal(ctx context.Context, in *GetAttendanceJournalRequest, opts ...grpc.CallOption) (*GetAttendanceJournalResponse, error) {
	out := new(GetAttendanceJournalResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetAttendanceJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SubmitTeacherCode(ctx context.Context, in *SubmitCodeRequest, opts ...grpc.CallOption) (*SubmitCodeResponse, error) {
	out := new(SubmitCodeResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SubmitTeacherCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SubmitRoomCode(ctx context.Context, in *SubmitCodeRequest, opts ...grpc.CallOption) (*SubmitCodeResponse, error) {
	out := new(SubmitCodeResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SubmitRoomCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAttendanceLessons(ctx context.Context, in *GetAttendanceLessonsRequest, opts ...grpc.CallOption) (*GetAttendanceLessonsResponse, error) {
	out := new(GetAttendanceLessonsResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetAttendanceLessons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// ------------------------------------------------------------//
	//
	//	REFACTOR OLD ONES                      //
	//
	// ------------------------------------------------------------//
	RegisterTeacher(context.Context, *RegisterTeacherRequest) (*RegisterTeacherResponse, error)
	RegisterStudent(context.Context, *RegisterStudentRequest) (*RegisterStudentResponse, error)
	LoginTeacher(context.Context, *LoginTeacherRequest) (*LoginTeacherResponse, error)
	LoginStudent(context.Context, *LoginStudentRequest) (*LoginStudentResponse, error)
	UpdateTeacherInfo(context.Context, *UpdateTeacherInfoRequest) (*UpdateTeacherInfoResponse, error)
	UpdateStudentInfo(context.Context, *UpdateStudentInfoRequest) (*UpdateStudentInfoResponse, error)
	LogoutTeacher(context.Context, *LogoutTeacherRequest) (*LogoutTeacherResponse, error)
	LogoutStudent(context.Context, *LogoutStudentRequest) (*LogoutStudentResponse, error)
	GetTeacherProfileData(context.Context, *emptypb.Empty) (*GetTeacherProfileDataResponse, error)
	GetStudentProfileData(context.Context, *emptypb.Empty) (*GetStudentProfileDataResponse, error)
	// at the end of lesson
	GetConfirmCode(*GetConfirmCodeRequest, Auth_GetConfirmCodeServer) error
	GetAttendanceJournal(context.Context, *GetAttendanceJournalRequest) (*GetAttendanceJournalResponse, error)
	// ----------------STUDENTS
	// both at the start and the end of lesson
	SubmitTeacherCode(context.Context, *SubmitCodeRequest) (*SubmitCodeResponse, error)
	SubmitRoomCode(context.Context, *SubmitCodeRequest) (*SubmitCodeResponse, error)
	GetAttendanceLessons(context.Context, *GetAttendanceLessonsRequest) (*GetAttendanceLessonsResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) RegisterTeacher(context.Context, *RegisterTeacherRequest) (*RegisterTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTeacher not implemented")
}
func (UnimplementedAuthServer) RegisterStudent(context.Context, *RegisterStudentRequest) (*RegisterStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterStudent not implemented")
}
func (UnimplementedAuthServer) LoginTeacher(context.Context, *LoginTeacherRequest) (*LoginTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginTeacher not implemented")
}
func (UnimplementedAuthServer) LoginStudent(context.Context, *LoginStudentRequest) (*LoginStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginStudent not implemented")
}
func (UnimplementedAuthServer) UpdateTeacherInfo(context.Context, *UpdateTeacherInfoRequest) (*UpdateTeacherInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeacherInfo not implemented")
}
func (UnimplementedAuthServer) UpdateStudentInfo(context.Context, *UpdateStudentInfoRequest) (*UpdateStudentInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudentInfo not implemented")
}
func (UnimplementedAuthServer) LogoutTeacher(context.Context, *LogoutTeacherRequest) (*LogoutTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutTeacher not implemented")
}
func (UnimplementedAuthServer) LogoutStudent(context.Context, *LogoutStudentRequest) (*LogoutStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutStudent not implemented")
}
func (UnimplementedAuthServer) GetTeacherProfileData(context.Context, *emptypb.Empty) (*GetTeacherProfileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeacherProfileData not implemented")
}
func (UnimplementedAuthServer) GetStudentProfileData(context.Context, *emptypb.Empty) (*GetStudentProfileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentProfileData not implemented")
}
func (UnimplementedAuthServer) GetConfirmCode(*GetConfirmCodeRequest, Auth_GetConfirmCodeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConfirmCode not implemented")
}
func (UnimplementedAuthServer) GetAttendanceJournal(context.Context, *GetAttendanceJournalRequest) (*GetAttendanceJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttendanceJournal not implemented")
}
func (UnimplementedAuthServer) SubmitTeacherCode(context.Context, *SubmitCodeRequest) (*SubmitCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTeacherCode not implemented")
}
func (UnimplementedAuthServer) SubmitRoomCode(context.Context, *SubmitCodeRequest) (*SubmitCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitRoomCode not implemented")
}
func (UnimplementedAuthServer) GetAttendanceLessons(context.Context, *GetAttendanceLessonsRequest) (*GetAttendanceLessonsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttendanceLessons not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_RegisterTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RegisterTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RegisterTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RegisterTeacher(ctx, req.(*RegisterTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RegisterStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RegisterStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RegisterStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RegisterStudent(ctx, req.(*RegisterStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LoginTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/LoginTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginTeacher(ctx, req.(*LoginTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LoginStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/LoginStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginStudent(ctx, req.(*LoginStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UpdateTeacherInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeacherInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UpdateTeacherInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/UpdateTeacherInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UpdateTeacherInfo(ctx, req.(*UpdateTeacherInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UpdateStudentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UpdateStudentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/UpdateStudentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UpdateStudentInfo(ctx, req.(*UpdateStudentInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LogoutTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LogoutTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/LogoutTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LogoutTeacher(ctx, req.(*LogoutTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LogoutStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LogoutStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/LogoutStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LogoutStudent(ctx, req.(*LogoutStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetTeacherProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetTeacherProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetTeacherProfileData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetTeacherProfileData(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetStudentProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetStudentProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetStudentProfileData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetStudentProfileData(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetConfirmCode_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetConfirmCodeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthServer).GetConfirmCode(m, &authGetConfirmCodeServer{stream})
}

type Auth_GetConfirmCodeServer interface {
	Send(*GetConfirmCodeResponse) error
	grpc.ServerStream
}

type authGetConfirmCodeServer struct {
	grpc.ServerStream
}

func (x *authGetConfirmCodeServer) Send(m *GetConfirmCodeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Auth_GetAttendanceJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttendanceJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAttendanceJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetAttendanceJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAttendanceJournal(ctx, req.(*GetAttendanceJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SubmitTeacherCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SubmitTeacherCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SubmitTeacherCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SubmitTeacherCode(ctx, req.(*SubmitCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SubmitRoomCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SubmitRoomCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SubmitRoomCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SubmitRoomCode(ctx, req.(*SubmitCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetAttendanceLessons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttendanceLessonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAttendanceLessons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetAttendanceLessons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAttendanceLessons(ctx, req.(*GetAttendanceLessonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterTeacher",
			Handler:    _Auth_RegisterTeacher_Handler,
		},
		{
			MethodName: "RegisterStudent",
			Handler:    _Auth_RegisterStudent_Handler,
		},
		{
			MethodName: "LoginTeacher",
			Handler:    _Auth_LoginTeacher_Handler,
		},
		{
			MethodName: "LoginStudent",
			Handler:    _Auth_LoginStudent_Handler,
		},
		{
			MethodName: "UpdateTeacherInfo",
			Handler:    _Auth_UpdateTeacherInfo_Handler,
		},
		{
			MethodName: "UpdateStudentInfo",
			Handler:    _Auth_UpdateStudentInfo_Handler,
		},
		{
			MethodName: "LogoutTeacher",
			Handler:    _Auth_LogoutTeacher_Handler,
		},
		{
			MethodName: "LogoutStudent",
			Handler:    _Auth_LogoutStudent_Handler,
		},
		{
			MethodName: "GetTeacherProfileData",
			Handler:    _Auth_GetTeacherProfileData_Handler,
		},
		{
			MethodName: "GetStudentProfileData",
			Handler:    _Auth_GetStudentProfileData_Handler,
		},
		{
			MethodName: "GetAttendanceJournal",
			Handler:    _Auth_GetAttendanceJournal_Handler,
		},
		{
			MethodName: "SubmitTeacherCode",
			Handler:    _Auth_SubmitTeacherCode_Handler,
		},
		{
			MethodName: "SubmitRoomCode",
			Handler:    _Auth_SubmitRoomCode_Handler,
		},
		{
			MethodName: "GetAttendanceLessons",
			Handler:    _Auth_GetAttendanceLessons_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetConfirmCode",
			Handler:       _Auth_GetConfirmCode_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/sso/sso.proto",
}