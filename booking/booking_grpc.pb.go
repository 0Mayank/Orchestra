// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto_def/booking.proto

package booking

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BookingService_CreateBooking_FullMethodName        = "/booking.BookingService/CreateBooking"
	BookingService_GetBooking_FullMethodName           = "/booking.BookingService/GetBooking"
	BookingService_UpdateBooking_FullMethodName        = "/booking.BookingService/UpdateBooking"
	BookingService_CancelBooking_FullMethodName        = "/booking.BookingService/CancelBooking"
	BookingService_ListCustomerBookings_FullMethodName = "/booking.BookingService/ListCustomerBookings"
	BookingService_CreateRooms_FullMethodName          = "/booking.BookingService/CreateRooms"
	BookingService_GetRooms_FullMethodName             = "/booking.BookingService/GetRooms"
	BookingService_CreateHotelEntry_FullMethodName     = "/booking.BookingService/CreateHotelEntry"
	BookingService_UpdateHotelEntry_FullMethodName     = "/booking.BookingService/UpdateHotelEntry"
)

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	// Create a new booking
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	// Get details of an existing booking
	GetBooking(ctx context.Context, in *GetBookingRequest, opts ...grpc.CallOption) (*GetBookingResponse, error)
	// Update an existing booking
	UpdateBooking(ctx context.Context, in *UpdateBookingRequest, opts ...grpc.CallOption) (*UpdateBookingResponse, error)
	// Cancel an existing booking
	CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*Empty, error)
	// List all bookings for a given customer
	ListCustomerBookings(ctx context.Context, in *ListCustomerBookingsRequest, opts ...grpc.CallOption) (*ListCustomerBookingsResponse, error)
	CreateRooms(ctx context.Context, in *CreateRoomsRequest, opts ...grpc.CallOption) (*Empty, error)
	// Get available room types and details
	GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error)
	// Create a new hotel entry
	CreateHotelEntry(ctx context.Context, in *CreateHotelEntryRequest, opts ...grpc.CallOption) (*CreateHotelEntryResponse, error)
	// Update an existing hotel entry
	UpdateHotelEntry(ctx context.Context, in *UpdateHotelEntryRequest, opts ...grpc.CallOption) (*UpdateHotelEntryResponse, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_CreateBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBooking(ctx context.Context, in *GetBookingRequest, opts ...grpc.CallOption) (*GetBookingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_GetBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) UpdateBooking(ctx context.Context, in *UpdateBookingRequest, opts ...grpc.CallOption) (*UpdateBookingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_UpdateBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, BookingService_CancelBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) ListCustomerBookings(ctx context.Context, in *ListCustomerBookingsRequest, opts ...grpc.CallOption) (*ListCustomerBookingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCustomerBookingsResponse)
	err := c.cc.Invoke(ctx, BookingService_ListCustomerBookings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CreateRooms(ctx context.Context, in *CreateRoomsRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, BookingService_CreateRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoomsResponse)
	err := c.cc.Invoke(ctx, BookingService_GetRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CreateHotelEntry(ctx context.Context, in *CreateHotelEntryRequest, opts ...grpc.CallOption) (*CreateHotelEntryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateHotelEntryResponse)
	err := c.cc.Invoke(ctx, BookingService_CreateHotelEntry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) UpdateHotelEntry(ctx context.Context, in *UpdateHotelEntryRequest, opts ...grpc.CallOption) (*UpdateHotelEntryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateHotelEntryResponse)
	err := c.cc.Invoke(ctx, BookingService_UpdateHotelEntry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility.
type BookingServiceServer interface {
	// Create a new booking
	CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	// Get details of an existing booking
	GetBooking(context.Context, *GetBookingRequest) (*GetBookingResponse, error)
	// Update an existing booking
	UpdateBooking(context.Context, *UpdateBookingRequest) (*UpdateBookingResponse, error)
	// Cancel an existing booking
	CancelBooking(context.Context, *CancelBookingRequest) (*Empty, error)
	// List all bookings for a given customer
	ListCustomerBookings(context.Context, *ListCustomerBookingsRequest) (*ListCustomerBookingsResponse, error)
	CreateRooms(context.Context, *CreateRoomsRequest) (*Empty, error)
	// Get available room types and details
	GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error)
	// Create a new hotel entry
	CreateHotelEntry(context.Context, *CreateHotelEntryRequest) (*CreateHotelEntryResponse, error)
	// Update an existing hotel entry
	UpdateHotelEntry(context.Context, *UpdateHotelEntryRequest) (*UpdateHotelEntryResponse, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookingServiceServer struct{}

func (UnimplementedBookingServiceServer) CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetBooking(context.Context, *GetBookingRequest) (*GetBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooking not implemented")
}
func (UnimplementedBookingServiceServer) UpdateBooking(context.Context, *UpdateBookingRequest) (*UpdateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBooking not implemented")
}
func (UnimplementedBookingServiceServer) CancelBooking(context.Context, *CancelBookingRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedBookingServiceServer) ListCustomerBookings(context.Context, *ListCustomerBookingsRequest) (*ListCustomerBookingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCustomerBookings not implemented")
}
func (UnimplementedBookingServiceServer) CreateRooms(context.Context, *CreateRoomsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRooms not implemented")
}
func (UnimplementedBookingServiceServer) GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRooms not implemented")
}
func (UnimplementedBookingServiceServer) CreateHotelEntry(context.Context, *CreateHotelEntryRequest) (*CreateHotelEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHotelEntry not implemented")
}
func (UnimplementedBookingServiceServer) UpdateHotelEntry(context.Context, *UpdateHotelEntryRequest) (*UpdateHotelEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHotelEntry not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}
func (UnimplementedBookingServiceServer) testEmbeddedByValue()                        {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	// If the following call pancis, it indicates UnimplementedBookingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CreateBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBooking(ctx, req.(*GetBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_UpdateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).UpdateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_UpdateBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).UpdateBooking(ctx, req.(*UpdateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CancelBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CancelBooking(ctx, req.(*CancelBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_ListCustomerBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCustomerBookingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).ListCustomerBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_ListCustomerBookings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).ListCustomerBookings(ctx, req.(*ListCustomerBookingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CreateRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CreateRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateRooms(ctx, req.(*CreateRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetRooms(ctx, req.(*GetRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CreateHotelEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHotelEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateHotelEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CreateHotelEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateHotelEntry(ctx, req.(*CreateHotelEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_UpdateHotelEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHotelEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).UpdateHotelEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_UpdateHotelEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).UpdateHotelEntry(ctx, req.(*UpdateHotelEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "GetBooking",
			Handler:    _BookingService_GetBooking_Handler,
		},
		{
			MethodName: "UpdateBooking",
			Handler:    _BookingService_UpdateBooking_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _BookingService_CancelBooking_Handler,
		},
		{
			MethodName: "ListCustomerBookings",
			Handler:    _BookingService_ListCustomerBookings_Handler,
		},
		{
			MethodName: "CreateRooms",
			Handler:    _BookingService_CreateRooms_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _BookingService_GetRooms_Handler,
		},
		{
			MethodName: "CreateHotelEntry",
			Handler:    _BookingService_CreateHotelEntry_Handler,
		},
		{
			MethodName: "UpdateHotelEntry",
			Handler:    _BookingService_UpdateHotelEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto_def/booking.proto",
}