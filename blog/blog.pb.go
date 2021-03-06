// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog.proto

package blogpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Slug                 string   `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	AuthorId             string   `protobuf:"bytes,5,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Post) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

type CreatePostRequest struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{1}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type CreatePostResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostResponse) Reset()         { *m = CreatePostResponse{} }
func (m *CreatePostResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePostResponse) ProtoMessage()    {}
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{2}
}

func (m *CreatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostResponse.Unmarshal(m, b)
}
func (m *CreatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostResponse.Marshal(b, m, deterministic)
}
func (m *CreatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostResponse.Merge(m, src)
}
func (m *CreatePostResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePostResponse.Size(m)
}
func (m *CreatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostResponse proto.InternalMessageInfo

func (m *CreatePostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type ReadPostByIdRequest struct {
	PostId               string   `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadPostByIdRequest) Reset()         { *m = ReadPostByIdRequest{} }
func (m *ReadPostByIdRequest) String() string { return proto.CompactTextString(m) }
func (*ReadPostByIdRequest) ProtoMessage()    {}
func (*ReadPostByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{3}
}

func (m *ReadPostByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPostByIdRequest.Unmarshal(m, b)
}
func (m *ReadPostByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPostByIdRequest.Marshal(b, m, deterministic)
}
func (m *ReadPostByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPostByIdRequest.Merge(m, src)
}
func (m *ReadPostByIdRequest) XXX_Size() int {
	return xxx_messageInfo_ReadPostByIdRequest.Size(m)
}
func (m *ReadPostByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPostByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPostByIdRequest proto.InternalMessageInfo

func (m *ReadPostByIdRequest) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

type ReadPostByIdResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadPostByIdResponse) Reset()         { *m = ReadPostByIdResponse{} }
func (m *ReadPostByIdResponse) String() string { return proto.CompactTextString(m) }
func (*ReadPostByIdResponse) ProtoMessage()    {}
func (*ReadPostByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{4}
}

func (m *ReadPostByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPostByIdResponse.Unmarshal(m, b)
}
func (m *ReadPostByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPostByIdResponse.Marshal(b, m, deterministic)
}
func (m *ReadPostByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPostByIdResponse.Merge(m, src)
}
func (m *ReadPostByIdResponse) XXX_Size() int {
	return xxx_messageInfo_ReadPostByIdResponse.Size(m)
}
func (m *ReadPostByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPostByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPostByIdResponse proto.InternalMessageInfo

func (m *ReadPostByIdResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type UpdatePostRequest struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostRequest) Reset()         { *m = UpdatePostRequest{} }
func (m *UpdatePostRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePostRequest) ProtoMessage()    {}
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{5}
}

func (m *UpdatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostRequest.Unmarshal(m, b)
}
func (m *UpdatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostRequest.Merge(m, src)
}
func (m *UpdatePostRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePostRequest.Size(m)
}
func (m *UpdatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostRequest proto.InternalMessageInfo

func (m *UpdatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type UpdatePostResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostResponse) Reset()         { *m = UpdatePostResponse{} }
func (m *UpdatePostResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePostResponse) ProtoMessage()    {}
func (*UpdatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{6}
}

func (m *UpdatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostResponse.Unmarshal(m, b)
}
func (m *UpdatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostResponse.Merge(m, src)
}
func (m *UpdatePostResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePostResponse.Size(m)
}
func (m *UpdatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostResponse proto.InternalMessageInfo

func (m *UpdatePostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type ListPostsRequests struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsRequests) Reset()         { *m = ListPostsRequests{} }
func (m *ListPostsRequests) String() string { return proto.CompactTextString(m) }
func (*ListPostsRequests) ProtoMessage()    {}
func (*ListPostsRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{7}
}

func (m *ListPostsRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsRequests.Unmarshal(m, b)
}
func (m *ListPostsRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsRequests.Marshal(b, m, deterministic)
}
func (m *ListPostsRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsRequests.Merge(m, src)
}
func (m *ListPostsRequests) XXX_Size() int {
	return xxx_messageInfo_ListPostsRequests.Size(m)
}
func (m *ListPostsRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsRequests.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsRequests proto.InternalMessageInfo

type ListPostsResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsResponse) Reset()         { *m = ListPostsResponse{} }
func (m *ListPostsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPostsResponse) ProtoMessage()    {}
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{8}
}

func (m *ListPostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsResponse.Unmarshal(m, b)
}
func (m *ListPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsResponse.Marshal(b, m, deterministic)
}
func (m *ListPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsResponse.Merge(m, src)
}
func (m *ListPostsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPostsResponse.Size(m)
}
func (m *ListPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsResponse proto.InternalMessageInfo

func (m *ListPostsResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "blog.Post")
	proto.RegisterType((*CreatePostRequest)(nil), "blog.CreatePostRequest")
	proto.RegisterType((*CreatePostResponse)(nil), "blog.CreatePostResponse")
	proto.RegisterType((*ReadPostByIdRequest)(nil), "blog.ReadPostByIdRequest")
	proto.RegisterType((*ReadPostByIdResponse)(nil), "blog.ReadPostByIdResponse")
	proto.RegisterType((*UpdatePostRequest)(nil), "blog.UpdatePostRequest")
	proto.RegisterType((*UpdatePostResponse)(nil), "blog.UpdatePostResponse")
	proto.RegisterType((*ListPostsRequests)(nil), "blog.ListPostsRequests")
	proto.RegisterType((*ListPostsResponse)(nil), "blog.ListPostsResponse")
}

func init() { proto.RegisterFile("blog.proto", fileDescriptor_6745b25902462fb1) }

var fileDescriptor_6745b25902462fb1 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0xa4, 0xf9, 0xa5, 0xff, 0xa6, 0x3f, 0x84, 0xbe, 0x16, 0xba, 0x46, 0x10, 0xc9, 0xc9, 0x53,
	0x91, 0x56, 0xbc, 0x16, 0xaa, 0x1e, 0x0a, 0x1e, 0xa4, 0xe2, 0xc5, 0x4b, 0x69, 0xbb, 0x4b, 0x5d,
	0x28, 0xdd, 0x98, 0xdd, 0x2a, 0x7e, 0x12, 0xbf, 0xae, 0xec, 0x26, 0x69, 0x23, 0x09, 0x18, 0x6f,
	0xfb, 0x66, 0xde, 0xec, 0xbc, 0x9d, 0xc7, 0x02, 0xab, 0xad, 0xda, 0x0c, 0xa3, 0x58, 0x19, 0x45,
	0xbe, 0x3d, 0x87, 0x1f, 0xf0, 0x1f, 0x95, 0x36, 0x74, 0x02, 0x4f, 0x72, 0x56, 0xbb, 0xa8, 0x5d,
	0xb6, 0xe7, 0x9e, 0xe4, 0xd4, 0x47, 0xdd, 0x48, 0xb3, 0x15, 0xcc, 0x73, 0x50, 0x52, 0x10, 0x43,
	0x73, 0xad, 0x76, 0x46, 0xec, 0x0c, 0xfb, 0xe7, 0xf0, 0xac, 0x24, 0x82, 0xaf, 0xb7, 0xfb, 0x0d,
	0xf3, 0x1d, 0xec, 0xce, 0x74, 0x86, 0xf6, 0x72, 0x6f, 0x5e, 0x55, 0xbc, 0x90, 0x9c, 0xd5, 0x1d,
	0xd1, 0x4a, 0x80, 0x19, 0x0f, 0xc7, 0xe8, 0xde, 0xc6, 0x62, 0x69, 0x84, 0xb5, 0x9f, 0x8b, 0xb7,
	0xbd, 0xd0, 0x86, 0xce, 0xe1, 0x47, 0x4a, 0x1b, 0x37, 0x47, 0x67, 0x84, 0xa1, 0x1b, 0xd7, 0x35,
	0x38, 0x3c, 0xbc, 0x06, 0xe5, 0x45, 0x3a, 0x52, 0x3b, 0x2d, 0x7e, 0x55, 0x0d, 0xd1, 0x9b, 0x8b,
	0x25, 0xb7, 0xc8, 0xf4, 0x73, 0xc6, 0x33, 0xb3, 0x01, 0x9a, 0x96, 0x5e, 0x1c, 0xde, 0xdd, 0xb0,
	0xe5, 0x8c, 0x87, 0x37, 0xe8, 0xff, 0xec, 0xaf, 0xe8, 0x33, 0x46, 0xf7, 0x39, 0xe2, 0x7f, 0x7f,
	0x52, 0x5e, 0x54, 0xd1, 0xaa, 0x87, 0xee, 0x83, 0xd4, 0xc6, 0x22, 0x3a, 0x75, 0xd2, 0xd6, 0x3f,
	0x07, 0x56, 0xbb, 0x69, 0xf4, 0xe5, 0xa1, 0x63, 0xcb, 0x27, 0x11, 0xbf, 0xcb, 0xb5, 0xa0, 0x09,
	0x70, 0x8c, 0x98, 0x06, 0x49, 0x7f, 0x61, 0x53, 0x01, 0x2b, 0x12, 0xa9, 0xe1, 0x3d, 0xfe, 0xe7,
	0xd2, 0xbb, 0xa3, 0xd3, 0xa4, 0xb3, 0x64, 0x03, 0x41, 0x50, 0x46, 0xa5, 0xd7, 0x4c, 0x80, 0x63,
	0x2e, 0xd9, 0x1c, 0x85, 0x78, 0xb3, 0x39, 0x4a, 0x22, 0x9c, 0xa0, 0x7d, 0x48, 0x23, 0xd3, 0x17,
	0x32, 0x0b, 0x8a, 0x44, 0x22, 0xbf, 0xaa, 0x4d, 0x5b, 0x2f, 0x0d, 0xcb, 0x45, 0xab, 0x55, 0xc3,
	0xfd, 0x98, 0xf1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x7e, 0x62, 0x9a, 0x3f, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	ReadPostByID(ctx context.Context, in *ReadPostByIdRequest, opts ...grpc.CallOption) (*ReadPostByIdResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error)
	ListPosts(ctx context.Context, in *ListPostsRequests, opts ...grpc.CallOption) (PostService_ListPostsClient, error)
}

type postServiceClient struct {
	cc *grpc.ClientConn
}

func NewPostServiceClient(cc *grpc.ClientConn) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/blog.PostService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ReadPostByID(ctx context.Context, in *ReadPostByIdRequest, opts ...grpc.CallOption) (*ReadPostByIdResponse, error) {
	out := new(ReadPostByIdResponse)
	err := c.cc.Invoke(ctx, "/blog.PostService/ReadPostByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error) {
	out := new(UpdatePostResponse)
	err := c.cc.Invoke(ctx, "/blog.PostService/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPosts(ctx context.Context, in *ListPostsRequests, opts ...grpc.CallOption) (PostService_ListPostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PostService_serviceDesc.Streams[0], "/blog.PostService/ListPosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &postServiceListPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PostService_ListPostsClient interface {
	Recv() (*ListPostsResponse, error)
	grpc.ClientStream
}

type postServiceListPostsClient struct {
	grpc.ClientStream
}

func (x *postServiceListPostsClient) Recv() (*ListPostsResponse, error) {
	m := new(ListPostsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PostServiceServer is the server API for PostService service.
type PostServiceServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	ReadPostByID(context.Context, *ReadPostByIdRequest) (*ReadPostByIdResponse, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error)
	ListPosts(*ListPostsRequests, PostService_ListPostsServer) error
}

// UnimplementedPostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (*UnimplementedPostServiceServer) CreatePost(ctx context.Context, req *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (*UnimplementedPostServiceServer) ReadPostByID(ctx context.Context, req *ReadPostByIdRequest) (*ReadPostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPostByID not implemented")
}
func (*UnimplementedPostServiceServer) UpdatePost(ctx context.Context, req *UpdatePostRequest) (*UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (*UnimplementedPostServiceServer) ListPosts(req *ListPostsRequests, srv PostService_ListPostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}

func RegisterPostServiceServer(s *grpc.Server, srv PostServiceServer) {
	s.RegisterService(&_PostService_serviceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.PostService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ReadPostByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPostByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ReadPostByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.PostService/ReadPostByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ReadPostByID(ctx, req.(*ReadPostByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.PostService/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ListPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPostsRequests)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostServiceServer).ListPosts(m, &postServiceListPostsServer{stream})
}

type PostService_ListPostsServer interface {
	Send(*ListPostsResponse) error
	grpc.ServerStream
}

type postServiceListPostsServer struct {
	grpc.ServerStream
}

func (x *postServiceListPostsServer) Send(m *ListPostsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "ReadPostByID",
			Handler:    _PostService_ReadPostByID_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostService_UpdatePost_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPosts",
			Handler:       _PostService_ListPosts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blog.proto",
}
