// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: grpc_game.proto

package mpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GameService_GetHiddenBoss_FullMethodName          = "/mpb.GameService/GetHiddenBoss"
	GameService_Fight_FullMethodName                  = "/mpb.GameService/Fight"
	GameService_FightPVP_FullMethodName               = "/mpb.GameService/FightPVP"
	GameService_GetPVPInfo_FullMethodName             = "/mpb.GameService/GetPVPInfo"
	GameService_GetPVPRanks_FullMethodName            = "/mpb.GameService/GetPVPRanks"
	GameService_GetPVPChallengeTargets_FullMethodName = "/mpb.GameService/GetPVPChallengeTargets"
	GameService_GetEnergy_FullMethodName              = "/mpb.GameService/GetEnergy"
	GameService_AddEnergy_FullMethodName              = "/mpb.GameService/AddEnergy"
	GameService_GetRandomHiddenBoss_FullMethodName    = "/mpb.GameService/GetRandomHiddenBoss"
	GameService_NewHiddenBoss_FullMethodName          = "/mpb.GameService/NewHiddenBoss"
	GameService_GetGameInfo_FullMethodName            = "/mpb.GameService/GetGameInfo"
	GameService_GetPVPHistory_FullMethodName          = "/mpb.GameService/GetPVPHistory"
	GameService_RandomBuffCards_FullMethodName        = "/mpb.GameService/RandomBuffCards"
	GameService_ChoseBuffCard_FullMethodName          = "/mpb.GameService/ChoseBuffCard"
	GameService_AdminRecoverEnergy_FullMethodName     = "/mpb.GameService/AdminRecoverEnergy"
)

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	GetHiddenBoss(ctx context.Context, in *ReqGetHiddenBoss, opts ...grpc.CallOption) (*ResGetHiddenBoss, error)
	Fight(ctx context.Context, in *ReqFight, opts ...grpc.CallOption) (*ResFight, error)
	FightPVP(ctx context.Context, in *ReqFightPVP, opts ...grpc.CallOption) (*ResFightPVP, error)
	GetPVPInfo(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*ResGetPVPInfo, error)
	GetPVPRanks(ctx context.Context, in *ReqGetPVPRanks, opts ...grpc.CallOption) (*ResGetPVPRanks, error)
	GetPVPChallengeTargets(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*ResGetPVPChallengeTargets, error)
	GetEnergy(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*ResGetEnergy, error)
	AddEnergy(ctx context.Context, in *ReqAddEnergy, opts ...grpc.CallOption) (*ResAddEnergy, error)
	GetRandomHiddenBoss(ctx context.Context, in *ReqGetHiddenBoss, opts ...grpc.CallOption) (*ResGetHiddenBoss, error)
	NewHiddenBoss(ctx context.Context, in *ReqNewHiddenBoss, opts ...grpc.CallOption) (*ResNewHiddenBoss, error)
	GetGameInfo(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*ResGetGameInfo, error)
	GetPVPHistory(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResGetPVPHistory, error)
	RandomBuffCards(ctx context.Context, in *ReqRandomBuffCards, opts ...grpc.CallOption) (*ResRandomBuffCards, error)
	ChoseBuffCard(ctx context.Context, in *ReqChoseBuffCard, opts ...grpc.CallOption) (*ResChoseBuffCard, error)
	AdminRecoverEnergy(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*Empty, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) GetHiddenBoss(ctx context.Context, in *ReqGetHiddenBoss, opts ...grpc.CallOption) (*ResGetHiddenBoss, error) {
	out := new(ResGetHiddenBoss)
	err := c.cc.Invoke(ctx, GameService_GetHiddenBoss_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) Fight(ctx context.Context, in *ReqFight, opts ...grpc.CallOption) (*ResFight, error) {
	out := new(ResFight)
	err := c.cc.Invoke(ctx, GameService_Fight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) FightPVP(ctx context.Context, in *ReqFightPVP, opts ...grpc.CallOption) (*ResFightPVP, error) {
	out := new(ResFightPVP)
	err := c.cc.Invoke(ctx, GameService_FightPVP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetPVPInfo(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*ResGetPVPInfo, error) {
	out := new(ResGetPVPInfo)
	err := c.cc.Invoke(ctx, GameService_GetPVPInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetPVPRanks(ctx context.Context, in *ReqGetPVPRanks, opts ...grpc.CallOption) (*ResGetPVPRanks, error) {
	out := new(ResGetPVPRanks)
	err := c.cc.Invoke(ctx, GameService_GetPVPRanks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetPVPChallengeTargets(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*ResGetPVPChallengeTargets, error) {
	out := new(ResGetPVPChallengeTargets)
	err := c.cc.Invoke(ctx, GameService_GetPVPChallengeTargets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetEnergy(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*ResGetEnergy, error) {
	out := new(ResGetEnergy)
	err := c.cc.Invoke(ctx, GameService_GetEnergy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) AddEnergy(ctx context.Context, in *ReqAddEnergy, opts ...grpc.CallOption) (*ResAddEnergy, error) {
	out := new(ResAddEnergy)
	err := c.cc.Invoke(ctx, GameService_AddEnergy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetRandomHiddenBoss(ctx context.Context, in *ReqGetHiddenBoss, opts ...grpc.CallOption) (*ResGetHiddenBoss, error) {
	out := new(ResGetHiddenBoss)
	err := c.cc.Invoke(ctx, GameService_GetRandomHiddenBoss_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) NewHiddenBoss(ctx context.Context, in *ReqNewHiddenBoss, opts ...grpc.CallOption) (*ResNewHiddenBoss, error) {
	out := new(ResNewHiddenBoss)
	err := c.cc.Invoke(ctx, GameService_NewHiddenBoss_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetGameInfo(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*ResGetGameInfo, error) {
	out := new(ResGetGameInfo)
	err := c.cc.Invoke(ctx, GameService_GetGameInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetPVPHistory(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResGetPVPHistory, error) {
	out := new(ResGetPVPHistory)
	err := c.cc.Invoke(ctx, GameService_GetPVPHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) RandomBuffCards(ctx context.Context, in *ReqRandomBuffCards, opts ...grpc.CallOption) (*ResRandomBuffCards, error) {
	out := new(ResRandomBuffCards)
	err := c.cc.Invoke(ctx, GameService_RandomBuffCards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ChoseBuffCard(ctx context.Context, in *ReqChoseBuffCard, opts ...grpc.CallOption) (*ResChoseBuffCard, error) {
	out := new(ResChoseBuffCard)
	err := c.cc.Invoke(ctx, GameService_ChoseBuffCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) AdminRecoverEnergy(ctx context.Context, in *ReqUserId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, GameService_AdminRecoverEnergy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	GetHiddenBoss(context.Context, *ReqGetHiddenBoss) (*ResGetHiddenBoss, error)
	Fight(context.Context, *ReqFight) (*ResFight, error)
	FightPVP(context.Context, *ReqFightPVP) (*ResFightPVP, error)
	GetPVPInfo(context.Context, *ReqUserId) (*ResGetPVPInfo, error)
	GetPVPRanks(context.Context, *ReqGetPVPRanks) (*ResGetPVPRanks, error)
	GetPVPChallengeTargets(context.Context, *ReqUserId) (*ResGetPVPChallengeTargets, error)
	GetEnergy(context.Context, *ReqUserId) (*ResGetEnergy, error)
	AddEnergy(context.Context, *ReqAddEnergy) (*ResAddEnergy, error)
	GetRandomHiddenBoss(context.Context, *ReqGetHiddenBoss) (*ResGetHiddenBoss, error)
	NewHiddenBoss(context.Context, *ReqNewHiddenBoss) (*ResNewHiddenBoss, error)
	GetGameInfo(context.Context, *ReqUserId) (*ResGetGameInfo, error)
	GetPVPHistory(context.Context, *Empty) (*ResGetPVPHistory, error)
	RandomBuffCards(context.Context, *ReqRandomBuffCards) (*ResRandomBuffCards, error)
	ChoseBuffCard(context.Context, *ReqChoseBuffCard) (*ResChoseBuffCard, error)
	AdminRecoverEnergy(context.Context, *ReqUserId) (*Empty, error)
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) GetHiddenBoss(context.Context, *ReqGetHiddenBoss) (*ResGetHiddenBoss, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHiddenBoss not implemented")
}
func (UnimplementedGameServiceServer) Fight(context.Context, *ReqFight) (*ResFight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fight not implemented")
}
func (UnimplementedGameServiceServer) FightPVP(context.Context, *ReqFightPVP) (*ResFightPVP, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FightPVP not implemented")
}
func (UnimplementedGameServiceServer) GetPVPInfo(context.Context, *ReqUserId) (*ResGetPVPInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPVPInfo not implemented")
}
func (UnimplementedGameServiceServer) GetPVPRanks(context.Context, *ReqGetPVPRanks) (*ResGetPVPRanks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPVPRanks not implemented")
}
func (UnimplementedGameServiceServer) GetPVPChallengeTargets(context.Context, *ReqUserId) (*ResGetPVPChallengeTargets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPVPChallengeTargets not implemented")
}
func (UnimplementedGameServiceServer) GetEnergy(context.Context, *ReqUserId) (*ResGetEnergy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnergy not implemented")
}
func (UnimplementedGameServiceServer) AddEnergy(context.Context, *ReqAddEnergy) (*ResAddEnergy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEnergy not implemented")
}
func (UnimplementedGameServiceServer) GetRandomHiddenBoss(context.Context, *ReqGetHiddenBoss) (*ResGetHiddenBoss, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomHiddenBoss not implemented")
}
func (UnimplementedGameServiceServer) NewHiddenBoss(context.Context, *ReqNewHiddenBoss) (*ResNewHiddenBoss, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewHiddenBoss not implemented")
}
func (UnimplementedGameServiceServer) GetGameInfo(context.Context, *ReqUserId) (*ResGetGameInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameInfo not implemented")
}
func (UnimplementedGameServiceServer) GetPVPHistory(context.Context, *Empty) (*ResGetPVPHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPVPHistory not implemented")
}
func (UnimplementedGameServiceServer) RandomBuffCards(context.Context, *ReqRandomBuffCards) (*ResRandomBuffCards, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RandomBuffCards not implemented")
}
func (UnimplementedGameServiceServer) ChoseBuffCard(context.Context, *ReqChoseBuffCard) (*ResChoseBuffCard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChoseBuffCard not implemented")
}
func (UnimplementedGameServiceServer) AdminRecoverEnergy(context.Context, *ReqUserId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminRecoverEnergy not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_GetHiddenBoss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetHiddenBoss)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetHiddenBoss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetHiddenBoss_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetHiddenBoss(ctx, req.(*ReqGetHiddenBoss))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_Fight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqFight)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).Fight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_Fight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).Fight(ctx, req.(*ReqFight))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_FightPVP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqFightPVP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).FightPVP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_FightPVP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).FightPVP(ctx, req.(*ReqFightPVP))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetPVPInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetPVPInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetPVPInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetPVPInfo(ctx, req.(*ReqUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetPVPRanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetPVPRanks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetPVPRanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetPVPRanks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetPVPRanks(ctx, req.(*ReqGetPVPRanks))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetPVPChallengeTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetPVPChallengeTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetPVPChallengeTargets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetPVPChallengeTargets(ctx, req.(*ReqUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetEnergy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetEnergy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetEnergy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetEnergy(ctx, req.(*ReqUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_AddEnergy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqAddEnergy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).AddEnergy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_AddEnergy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).AddEnergy(ctx, req.(*ReqAddEnergy))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetRandomHiddenBoss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetHiddenBoss)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetRandomHiddenBoss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetRandomHiddenBoss_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetRandomHiddenBoss(ctx, req.(*ReqGetHiddenBoss))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_NewHiddenBoss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqNewHiddenBoss)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).NewHiddenBoss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_NewHiddenBoss_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).NewHiddenBoss(ctx, req.(*ReqNewHiddenBoss))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetGameInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetGameInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetGameInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetGameInfo(ctx, req.(*ReqUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetPVPHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetPVPHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetPVPHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetPVPHistory(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_RandomBuffCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqRandomBuffCards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).RandomBuffCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_RandomBuffCards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).RandomBuffCards(ctx, req.(*ReqRandomBuffCards))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ChoseBuffCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqChoseBuffCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ChoseBuffCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_ChoseBuffCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ChoseBuffCard(ctx, req.(*ReqChoseBuffCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_AdminRecoverEnergy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).AdminRecoverEnergy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_AdminRecoverEnergy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).AdminRecoverEnergy(ctx, req.(*ReqUserId))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mpb.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHiddenBoss",
			Handler:    _GameService_GetHiddenBoss_Handler,
		},
		{
			MethodName: "Fight",
			Handler:    _GameService_Fight_Handler,
		},
		{
			MethodName: "FightPVP",
			Handler:    _GameService_FightPVP_Handler,
		},
		{
			MethodName: "GetPVPInfo",
			Handler:    _GameService_GetPVPInfo_Handler,
		},
		{
			MethodName: "GetPVPRanks",
			Handler:    _GameService_GetPVPRanks_Handler,
		},
		{
			MethodName: "GetPVPChallengeTargets",
			Handler:    _GameService_GetPVPChallengeTargets_Handler,
		},
		{
			MethodName: "GetEnergy",
			Handler:    _GameService_GetEnergy_Handler,
		},
		{
			MethodName: "AddEnergy",
			Handler:    _GameService_AddEnergy_Handler,
		},
		{
			MethodName: "GetRandomHiddenBoss",
			Handler:    _GameService_GetRandomHiddenBoss_Handler,
		},
		{
			MethodName: "NewHiddenBoss",
			Handler:    _GameService_NewHiddenBoss_Handler,
		},
		{
			MethodName: "GetGameInfo",
			Handler:    _GameService_GetGameInfo_Handler,
		},
		{
			MethodName: "GetPVPHistory",
			Handler:    _GameService_GetPVPHistory_Handler,
		},
		{
			MethodName: "RandomBuffCards",
			Handler:    _GameService_RandomBuffCards_Handler,
		},
		{
			MethodName: "ChoseBuffCard",
			Handler:    _GameService_ChoseBuffCard_Handler,
		},
		{
			MethodName: "AdminRecoverEnergy",
			Handler:    _GameService_AdminRecoverEnergy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_game.proto",
}