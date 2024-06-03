package grpc

import (
	"context"

	"github.com/apecloud/kubeblocks/pkg/kb_agent/dcs"
	"github.com/apecloud/kubeblocks/pkg/kb_agent/plugin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
)

var _ = Describe("GRPC Handler", func() {
	var (
		handler *Handler
		cluster *dcs.Cluster
		ctx     context.Context
	)

	BeforeEach(func() {
		handler = &Handler{}
		cluster = &dcs.Cluster{}
		ctx = context.Background()
	})

	Describe("IsDBStartupReady", func() {
		Context("when DBStartupReady is true", func() {
			BeforeEach(func() {
				handler.DBStartupReady = true
			})

			It("should return true", func() {
				Expect(handler.IsDBStartupReady()).To(BeTrue())
			})
		})

		Context("when DBStartupReady is false", func() {
			BeforeEach(func() {
				handler.DBStartupReady = false
			})

			It("should return the response from dbClient", func() {
				mockIsServiceReady := &plugin.IsServiceReadyResponse{
					Ready: true,
				}
				mockIsServiceReadyFunc := func(ctx context.Context, req *plugin.IsServiceReadyRequest) (*plugin.IsServiceReadyResponse, error) {
					return mockIsServiceReady, nil
				}
				handler.dbClient = &mockServicePluginClient{
					mockIsServiceReady: mockIsServiceReadyFunc,
				}

				Expect(handler.IsDBStartupReady()).To(BeTrue())
			})
		})
	})

	Describe("JoinMember", func() {
		It("should call dbClient.JoinMember with the correct request", func() {
			mockJoinMember := &plugin.JoinMemberRequest{
				ServiceInfo: &plugin.ServiceInfo{},
				NewMember:   handler.CurrentMemberName,
				Members:     cluster.GetMemberAddrs(),
			}
			mockJoinMemberFunc := func(ctx context.Context, req *plugin.JoinMemberRequest) (*plugin.JoinMemberResponse, error) {
				Expect(req).To(Equal(mockJoinMember))
				return nil, nil
			}
			handler.dbClient = &mockServicePluginClient{
				mockJoinMember: mockJoinMemberFunc,
			}

			Expect(handler.JoinMember(ctx, cluster, "memberName")).To(Succeed())
		})
	})

	Describe("LeaveMember", func() {
		It("should call dbClient.LeaveMember with the correct request", func() {
			memberName := "test-pod-0"
			mockLeaveMember := &plugin.LeaveMemberRequest{
				ServiceInfo: &plugin.ServiceInfo{},
				LeaveMember: memberName,
				Members:     cluster.GetMemberAddrs(),
			}
			mockLeaveMemberFunc := func(ctx context.Context, req *plugin.LeaveMemberRequest) (*plugin.LeaveMemberResponse, error) {
				Expect(req).To(Equal(mockLeaveMember))
				return nil, nil
			}
			handler.dbClient = &mockServicePluginClient{
				mockLeaveMember: mockLeaveMemberFunc,
			}

			Expect(handler.LeaveMember(ctx, cluster, memberName)).To(Succeed())
		})
	})

	Describe("Lock", func() {
		It("should call dbClient.Readonly with the correct request", func() {
			reason := "reason for test lock"
			mockReadonly := &plugin.ReadonlyRequest{
				ServiceInfo: &plugin.ServiceInfo{},
				Reason:      reason,
			}
			mockReadonlyFunc := func(ctx context.Context, req *plugin.ReadonlyRequest) (*plugin.ReadonlyResponse, error) {
				Expect(req).To(Equal(mockReadonly))
				return nil, nil
			}
			handler.dbClient = &mockServicePluginClient{
				mockReadonly: mockReadonlyFunc,
			}

			Expect(handler.Lock(ctx, reason)).To(Succeed())
		})
	})

	Describe("Unlock", func() {
		It("should call dbClient.Readwrite with the correct request", func() {
			reason := "reason for test unlock"
			mockReadwrite := &plugin.ReadwriteRequest{
				ServiceInfo: &plugin.ServiceInfo{},
				Reason:      reason,
			}
			mockReadwriteFunc := func(ctx context.Context, req *plugin.ReadwriteRequest) (*plugin.ReadwriteResponse, error) {
				Expect(req).To(Equal(mockReadwrite))
				return nil, nil
			}
			handler.dbClient = &mockServicePluginClient{
				mockReadwrite: mockReadwriteFunc,
			}

			Expect(handler.Unlock(ctx, reason)).To(Succeed())
		})
	})
})

type mockServicePluginClient struct {
	plugin.ServicePluginClient

	mockIsServiceReady func(ctx context.Context, req *plugin.IsServiceReadyRequest) (*plugin.IsServiceReadyResponse, error)
	mockJoinMember     func(ctx context.Context, req *plugin.JoinMemberRequest) (*plugin.JoinMemberResponse, error)
	mockLeaveMember    func(ctx context.Context, req *plugin.LeaveMemberRequest) (*plugin.LeaveMemberResponse, error)
	mockReadonly       func(ctx context.Context, req *plugin.ReadonlyRequest) (*plugin.ReadonlyResponse, error)
	mockReadwrite      func(ctx context.Context, req *plugin.ReadwriteRequest) (*plugin.ReadwriteResponse, error)
}

func (m *mockServicePluginClient) IsServiceReady(ctx context.Context, req *plugin.IsServiceReadyRequest, opts ...grpc.CallOption) (*plugin.IsServiceReadyResponse, error) {
	if m.mockIsServiceReady != nil {
		return m.mockIsServiceReady(ctx, req)
	}
	return nil, nil
}

func (m *mockServicePluginClient) JoinMember(ctx context.Context, req *plugin.JoinMemberRequest, opts ...grpc.CallOption) (*plugin.JoinMemberResponse, error) {
	if m.mockJoinMember != nil {
		return m.mockJoinMember(ctx, req)
	}
	return nil, nil
}

func (m *mockServicePluginClient) LeaveMember(ctx context.Context, req *plugin.LeaveMemberRequest, opts ...grpc.CallOption) (*plugin.LeaveMemberResponse, error) {
	if m.mockLeaveMember != nil {
		return m.mockLeaveMember(ctx, req)
	}
	return nil, nil
}

func (m *mockServicePluginClient) Readonly(ctx context.Context, req *plugin.ReadonlyRequest, opts ...grpc.CallOption) (*plugin.ReadonlyResponse, error) {
	if m.mockReadonly != nil {
		return m.mockReadonly(ctx, req)
	}
	return nil, nil
}

func (m *mockServicePluginClient) Readwrite(ctx context.Context, req *plugin.ReadwriteRequest, opts ...grpc.CallOption) (*plugin.ReadwriteResponse, error) {
	if m.mockReadwrite != nil {
		return m.mockReadwrite(ctx, req)
	}
	return nil, nil
}
