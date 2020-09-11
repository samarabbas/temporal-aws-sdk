package awsclients

import (
	"github.com/aws/aws-sdk-go/service/appmesh"
	"go.temporal.io/sdk/workflow"
)

type AppMeshClient interface {
       CreateGatewayRoute(ctx workflow.Context, input *appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error)
       CreateGatewayRouteAsync(ctx workflow.Context, input *appmesh.CreateGatewayRouteInput) *AppmeshCreateGatewayRouteResult

       CreateMesh(ctx workflow.Context, input *appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error)
       CreateMeshAsync(ctx workflow.Context, input *appmesh.CreateMeshInput) *AppmeshCreateMeshResult

       CreateRoute(ctx workflow.Context, input *appmesh.CreateRouteInput) (*appmesh.CreateRouteOutput, error)
       CreateRouteAsync(ctx workflow.Context, input *appmesh.CreateRouteInput) *AppmeshCreateRouteResult

       CreateVirtualGateway(ctx workflow.Context, input *appmesh.CreateVirtualGatewayInput) (*appmesh.CreateVirtualGatewayOutput, error)
       CreateVirtualGatewayAsync(ctx workflow.Context, input *appmesh.CreateVirtualGatewayInput) *AppmeshCreateVirtualGatewayResult

       CreateVirtualNode(ctx workflow.Context, input *appmesh.CreateVirtualNodeInput) (*appmesh.CreateVirtualNodeOutput, error)
       CreateVirtualNodeAsync(ctx workflow.Context, input *appmesh.CreateVirtualNodeInput) *AppmeshCreateVirtualNodeResult

       CreateVirtualRouter(ctx workflow.Context, input *appmesh.CreateVirtualRouterInput) (*appmesh.CreateVirtualRouterOutput, error)
       CreateVirtualRouterAsync(ctx workflow.Context, input *appmesh.CreateVirtualRouterInput) *AppmeshCreateVirtualRouterResult

       CreateVirtualService(ctx workflow.Context, input *appmesh.CreateVirtualServiceInput) (*appmesh.CreateVirtualServiceOutput, error)
       CreateVirtualServiceAsync(ctx workflow.Context, input *appmesh.CreateVirtualServiceInput) *AppmeshCreateVirtualServiceResult

       DeleteGatewayRoute(ctx workflow.Context, input *appmesh.DeleteGatewayRouteInput) (*appmesh.DeleteGatewayRouteOutput, error)
       DeleteGatewayRouteAsync(ctx workflow.Context, input *appmesh.DeleteGatewayRouteInput) *AppmeshDeleteGatewayRouteResult

       DeleteMesh(ctx workflow.Context, input *appmesh.DeleteMeshInput) (*appmesh.DeleteMeshOutput, error)
       DeleteMeshAsync(ctx workflow.Context, input *appmesh.DeleteMeshInput) *AppmeshDeleteMeshResult

       DeleteRoute(ctx workflow.Context, input *appmesh.DeleteRouteInput) (*appmesh.DeleteRouteOutput, error)
       DeleteRouteAsync(ctx workflow.Context, input *appmesh.DeleteRouteInput) *AppmeshDeleteRouteResult

       DeleteVirtualGateway(ctx workflow.Context, input *appmesh.DeleteVirtualGatewayInput) (*appmesh.DeleteVirtualGatewayOutput, error)
       DeleteVirtualGatewayAsync(ctx workflow.Context, input *appmesh.DeleteVirtualGatewayInput) *AppmeshDeleteVirtualGatewayResult

       DeleteVirtualNode(ctx workflow.Context, input *appmesh.DeleteVirtualNodeInput) (*appmesh.DeleteVirtualNodeOutput, error)
       DeleteVirtualNodeAsync(ctx workflow.Context, input *appmesh.DeleteVirtualNodeInput) *AppmeshDeleteVirtualNodeResult

       DeleteVirtualRouter(ctx workflow.Context, input *appmesh.DeleteVirtualRouterInput) (*appmesh.DeleteVirtualRouterOutput, error)
       DeleteVirtualRouterAsync(ctx workflow.Context, input *appmesh.DeleteVirtualRouterInput) *AppmeshDeleteVirtualRouterResult

       DeleteVirtualService(ctx workflow.Context, input *appmesh.DeleteVirtualServiceInput) (*appmesh.DeleteVirtualServiceOutput, error)
       DeleteVirtualServiceAsync(ctx workflow.Context, input *appmesh.DeleteVirtualServiceInput) *AppmeshDeleteVirtualServiceResult

       DescribeGatewayRoute(ctx workflow.Context, input *appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error)
       DescribeGatewayRouteAsync(ctx workflow.Context, input *appmesh.DescribeGatewayRouteInput) *AppmeshDescribeGatewayRouteResult

       DescribeMesh(ctx workflow.Context, input *appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error)
       DescribeMeshAsync(ctx workflow.Context, input *appmesh.DescribeMeshInput) *AppmeshDescribeMeshResult

       DescribeRoute(ctx workflow.Context, input *appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error)
       DescribeRouteAsync(ctx workflow.Context, input *appmesh.DescribeRouteInput) *AppmeshDescribeRouteResult

       DescribeVirtualGateway(ctx workflow.Context, input *appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error)
       DescribeVirtualGatewayAsync(ctx workflow.Context, input *appmesh.DescribeVirtualGatewayInput) *AppmeshDescribeVirtualGatewayResult

       DescribeVirtualNode(ctx workflow.Context, input *appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error)
       DescribeVirtualNodeAsync(ctx workflow.Context, input *appmesh.DescribeVirtualNodeInput) *AppmeshDescribeVirtualNodeResult

       DescribeVirtualRouter(ctx workflow.Context, input *appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error)
       DescribeVirtualRouterAsync(ctx workflow.Context, input *appmesh.DescribeVirtualRouterInput) *AppmeshDescribeVirtualRouterResult

       DescribeVirtualService(ctx workflow.Context, input *appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error)
       DescribeVirtualServiceAsync(ctx workflow.Context, input *appmesh.DescribeVirtualServiceInput) *AppmeshDescribeVirtualServiceResult

       ListGatewayRoutes(ctx workflow.Context, input *appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error)
       ListGatewayRoutesAsync(ctx workflow.Context, input *appmesh.ListGatewayRoutesInput) *AppmeshListGatewayRoutesResult

       ListMeshes(ctx workflow.Context, input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error)
       ListMeshesAsync(ctx workflow.Context, input *appmesh.ListMeshesInput) *AppmeshListMeshesResult

       ListRoutes(ctx workflow.Context, input *appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error)
       ListRoutesAsync(ctx workflow.Context, input *appmesh.ListRoutesInput) *AppmeshListRoutesResult

       ListTagsForResource(ctx workflow.Context, input *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *appmesh.ListTagsForResourceInput) *AppmeshListTagsForResourceResult

       ListVirtualGateways(ctx workflow.Context, input *appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error)
       ListVirtualGatewaysAsync(ctx workflow.Context, input *appmesh.ListVirtualGatewaysInput) *AppmeshListVirtualGatewaysResult

       ListVirtualNodes(ctx workflow.Context, input *appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error)
       ListVirtualNodesAsync(ctx workflow.Context, input *appmesh.ListVirtualNodesInput) *AppmeshListVirtualNodesResult

       ListVirtualRouters(ctx workflow.Context, input *appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error)
       ListVirtualRoutersAsync(ctx workflow.Context, input *appmesh.ListVirtualRoutersInput) *AppmeshListVirtualRoutersResult

       ListVirtualServices(ctx workflow.Context, input *appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error)
       ListVirtualServicesAsync(ctx workflow.Context, input *appmesh.ListVirtualServicesInput) *AppmeshListVirtualServicesResult

       TagResource(ctx workflow.Context, input *appmesh.TagResourceInput) (*appmesh.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *appmesh.TagResourceInput) *AppmeshTagResourceResult

       UntagResource(ctx workflow.Context, input *appmesh.UntagResourceInput) (*appmesh.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *appmesh.UntagResourceInput) *AppmeshUntagResourceResult

       UpdateGatewayRoute(ctx workflow.Context, input *appmesh.UpdateGatewayRouteInput) (*appmesh.UpdateGatewayRouteOutput, error)
       UpdateGatewayRouteAsync(ctx workflow.Context, input *appmesh.UpdateGatewayRouteInput) *AppmeshUpdateGatewayRouteResult

       UpdateMesh(ctx workflow.Context, input *appmesh.UpdateMeshInput) (*appmesh.UpdateMeshOutput, error)
       UpdateMeshAsync(ctx workflow.Context, input *appmesh.UpdateMeshInput) *AppmeshUpdateMeshResult

       UpdateRoute(ctx workflow.Context, input *appmesh.UpdateRouteInput) (*appmesh.UpdateRouteOutput, error)
       UpdateRouteAsync(ctx workflow.Context, input *appmesh.UpdateRouteInput) *AppmeshUpdateRouteResult

       UpdateVirtualGateway(ctx workflow.Context, input *appmesh.UpdateVirtualGatewayInput) (*appmesh.UpdateVirtualGatewayOutput, error)
       UpdateVirtualGatewayAsync(ctx workflow.Context, input *appmesh.UpdateVirtualGatewayInput) *AppmeshUpdateVirtualGatewayResult

       UpdateVirtualNode(ctx workflow.Context, input *appmesh.UpdateVirtualNodeInput) (*appmesh.UpdateVirtualNodeOutput, error)
       UpdateVirtualNodeAsync(ctx workflow.Context, input *appmesh.UpdateVirtualNodeInput) *AppmeshUpdateVirtualNodeResult

       UpdateVirtualRouter(ctx workflow.Context, input *appmesh.UpdateVirtualRouterInput) (*appmesh.UpdateVirtualRouterOutput, error)
       UpdateVirtualRouterAsync(ctx workflow.Context, input *appmesh.UpdateVirtualRouterInput) *AppmeshUpdateVirtualRouterResult

       UpdateVirtualService(ctx workflow.Context, input *appmesh.UpdateVirtualServiceInput) (*appmesh.UpdateVirtualServiceOutput, error)
       UpdateVirtualServiceAsync(ctx workflow.Context, input *appmesh.UpdateVirtualServiceInput) *AppmeshUpdateVirtualServiceResult
}

type AppMeshStub struct {
}

func NewAppMeshStub() AppMeshClient {
    return &AppMeshStub{}
}

type AppmeshCreateGatewayRouteResult struct {
	Result workflow.Future
}

func (r *AppmeshCreateGatewayRouteResult) Get(ctx workflow.Context) (*appmesh.CreateGatewayRouteOutput, error) {
    var output appmesh.CreateGatewayRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshCreateMeshResult struct {
	Result workflow.Future
}

func (r *AppmeshCreateMeshResult) Get(ctx workflow.Context) (*appmesh.CreateMeshOutput, error) {
    var output appmesh.CreateMeshOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshCreateRouteResult struct {
	Result workflow.Future
}

func (r *AppmeshCreateRouteResult) Get(ctx workflow.Context) (*appmesh.CreateRouteOutput, error) {
    var output appmesh.CreateRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshCreateVirtualGatewayResult struct {
	Result workflow.Future
}

func (r *AppmeshCreateVirtualGatewayResult) Get(ctx workflow.Context) (*appmesh.CreateVirtualGatewayOutput, error) {
    var output appmesh.CreateVirtualGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshCreateVirtualNodeResult struct {
	Result workflow.Future
}

func (r *AppmeshCreateVirtualNodeResult) Get(ctx workflow.Context) (*appmesh.CreateVirtualNodeOutput, error) {
    var output appmesh.CreateVirtualNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshCreateVirtualRouterResult struct {
	Result workflow.Future
}

func (r *AppmeshCreateVirtualRouterResult) Get(ctx workflow.Context) (*appmesh.CreateVirtualRouterOutput, error) {
    var output appmesh.CreateVirtualRouterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshCreateVirtualServiceResult struct {
	Result workflow.Future
}

func (r *AppmeshCreateVirtualServiceResult) Get(ctx workflow.Context) (*appmesh.CreateVirtualServiceOutput, error) {
    var output appmesh.CreateVirtualServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDeleteGatewayRouteResult struct {
	Result workflow.Future
}

func (r *AppmeshDeleteGatewayRouteResult) Get(ctx workflow.Context) (*appmesh.DeleteGatewayRouteOutput, error) {
    var output appmesh.DeleteGatewayRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDeleteMeshResult struct {
	Result workflow.Future
}

func (r *AppmeshDeleteMeshResult) Get(ctx workflow.Context) (*appmesh.DeleteMeshOutput, error) {
    var output appmesh.DeleteMeshOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDeleteRouteResult struct {
	Result workflow.Future
}

func (r *AppmeshDeleteRouteResult) Get(ctx workflow.Context) (*appmesh.DeleteRouteOutput, error) {
    var output appmesh.DeleteRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDeleteVirtualGatewayResult struct {
	Result workflow.Future
}

func (r *AppmeshDeleteVirtualGatewayResult) Get(ctx workflow.Context) (*appmesh.DeleteVirtualGatewayOutput, error) {
    var output appmesh.DeleteVirtualGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDeleteVirtualNodeResult struct {
	Result workflow.Future
}

func (r *AppmeshDeleteVirtualNodeResult) Get(ctx workflow.Context) (*appmesh.DeleteVirtualNodeOutput, error) {
    var output appmesh.DeleteVirtualNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDeleteVirtualRouterResult struct {
	Result workflow.Future
}

func (r *AppmeshDeleteVirtualRouterResult) Get(ctx workflow.Context) (*appmesh.DeleteVirtualRouterOutput, error) {
    var output appmesh.DeleteVirtualRouterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDeleteVirtualServiceResult struct {
	Result workflow.Future
}

func (r *AppmeshDeleteVirtualServiceResult) Get(ctx workflow.Context) (*appmesh.DeleteVirtualServiceOutput, error) {
    var output appmesh.DeleteVirtualServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDescribeGatewayRouteResult struct {
	Result workflow.Future
}

func (r *AppmeshDescribeGatewayRouteResult) Get(ctx workflow.Context) (*appmesh.DescribeGatewayRouteOutput, error) {
    var output appmesh.DescribeGatewayRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDescribeMeshResult struct {
	Result workflow.Future
}

func (r *AppmeshDescribeMeshResult) Get(ctx workflow.Context) (*appmesh.DescribeMeshOutput, error) {
    var output appmesh.DescribeMeshOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDescribeRouteResult struct {
	Result workflow.Future
}

func (r *AppmeshDescribeRouteResult) Get(ctx workflow.Context) (*appmesh.DescribeRouteOutput, error) {
    var output appmesh.DescribeRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDescribeVirtualGatewayResult struct {
	Result workflow.Future
}

func (r *AppmeshDescribeVirtualGatewayResult) Get(ctx workflow.Context) (*appmesh.DescribeVirtualGatewayOutput, error) {
    var output appmesh.DescribeVirtualGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDescribeVirtualNodeResult struct {
	Result workflow.Future
}

func (r *AppmeshDescribeVirtualNodeResult) Get(ctx workflow.Context) (*appmesh.DescribeVirtualNodeOutput, error) {
    var output appmesh.DescribeVirtualNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDescribeVirtualRouterResult struct {
	Result workflow.Future
}

func (r *AppmeshDescribeVirtualRouterResult) Get(ctx workflow.Context) (*appmesh.DescribeVirtualRouterOutput, error) {
    var output appmesh.DescribeVirtualRouterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshDescribeVirtualServiceResult struct {
	Result workflow.Future
}

func (r *AppmeshDescribeVirtualServiceResult) Get(ctx workflow.Context) (*appmesh.DescribeVirtualServiceOutput, error) {
    var output appmesh.DescribeVirtualServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshListGatewayRoutesResult struct {
	Result workflow.Future
}

func (r *AppmeshListGatewayRoutesResult) Get(ctx workflow.Context) (*appmesh.ListGatewayRoutesOutput, error) {
    var output appmesh.ListGatewayRoutesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshListMeshesResult struct {
	Result workflow.Future
}

func (r *AppmeshListMeshesResult) Get(ctx workflow.Context) (*appmesh.ListMeshesOutput, error) {
    var output appmesh.ListMeshesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshListRoutesResult struct {
	Result workflow.Future
}

func (r *AppmeshListRoutesResult) Get(ctx workflow.Context) (*appmesh.ListRoutesOutput, error) {
    var output appmesh.ListRoutesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *AppmeshListTagsForResourceResult) Get(ctx workflow.Context) (*appmesh.ListTagsForResourceOutput, error) {
    var output appmesh.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshListVirtualGatewaysResult struct {
	Result workflow.Future
}

func (r *AppmeshListVirtualGatewaysResult) Get(ctx workflow.Context) (*appmesh.ListVirtualGatewaysOutput, error) {
    var output appmesh.ListVirtualGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshListVirtualNodesResult struct {
	Result workflow.Future
}

func (r *AppmeshListVirtualNodesResult) Get(ctx workflow.Context) (*appmesh.ListVirtualNodesOutput, error) {
    var output appmesh.ListVirtualNodesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshListVirtualRoutersResult struct {
	Result workflow.Future
}

func (r *AppmeshListVirtualRoutersResult) Get(ctx workflow.Context) (*appmesh.ListVirtualRoutersOutput, error) {
    var output appmesh.ListVirtualRoutersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshListVirtualServicesResult struct {
	Result workflow.Future
}

func (r *AppmeshListVirtualServicesResult) Get(ctx workflow.Context) (*appmesh.ListVirtualServicesOutput, error) {
    var output appmesh.ListVirtualServicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshTagResourceResult struct {
	Result workflow.Future
}

func (r *AppmeshTagResourceResult) Get(ctx workflow.Context) (*appmesh.TagResourceOutput, error) {
    var output appmesh.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshUntagResourceResult struct {
	Result workflow.Future
}

func (r *AppmeshUntagResourceResult) Get(ctx workflow.Context) (*appmesh.UntagResourceOutput, error) {
    var output appmesh.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshUpdateGatewayRouteResult struct {
	Result workflow.Future
}

func (r *AppmeshUpdateGatewayRouteResult) Get(ctx workflow.Context) (*appmesh.UpdateGatewayRouteOutput, error) {
    var output appmesh.UpdateGatewayRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshUpdateMeshResult struct {
	Result workflow.Future
}

func (r *AppmeshUpdateMeshResult) Get(ctx workflow.Context) (*appmesh.UpdateMeshOutput, error) {
    var output appmesh.UpdateMeshOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshUpdateRouteResult struct {
	Result workflow.Future
}

func (r *AppmeshUpdateRouteResult) Get(ctx workflow.Context) (*appmesh.UpdateRouteOutput, error) {
    var output appmesh.UpdateRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshUpdateVirtualGatewayResult struct {
	Result workflow.Future
}

func (r *AppmeshUpdateVirtualGatewayResult) Get(ctx workflow.Context) (*appmesh.UpdateVirtualGatewayOutput, error) {
    var output appmesh.UpdateVirtualGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshUpdateVirtualNodeResult struct {
	Result workflow.Future
}

func (r *AppmeshUpdateVirtualNodeResult) Get(ctx workflow.Context) (*appmesh.UpdateVirtualNodeOutput, error) {
    var output appmesh.UpdateVirtualNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshUpdateVirtualRouterResult struct {
	Result workflow.Future
}

func (r *AppmeshUpdateVirtualRouterResult) Get(ctx workflow.Context) (*appmesh.UpdateVirtualRouterOutput, error) {
    var output appmesh.UpdateVirtualRouterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppmeshUpdateVirtualServiceResult struct {
	Result workflow.Future
}

func (r *AppmeshUpdateVirtualServiceResult) Get(ctx workflow.Context) (*appmesh.UpdateVirtualServiceOutput, error) {
    var output appmesh.UpdateVirtualServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) CreateGatewayRoute(ctx workflow.Context, input *appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error) {
    var output appmesh.CreateGatewayRouteOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.CreateGatewayRoute", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) CreateGatewayRouteAsync(ctx workflow.Context, input *appmesh.CreateGatewayRouteInput) *AppmeshCreateGatewayRouteResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.CreateGatewayRoute", input)
    return &AppmeshCreateGatewayRouteResult{Result: future}
}

func (a *AppMeshStub) CreateMesh(ctx workflow.Context, input *appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error) {
    var output appmesh.CreateMeshOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.CreateMesh", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) CreateMeshAsync(ctx workflow.Context, input *appmesh.CreateMeshInput) *AppmeshCreateMeshResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.CreateMesh", input)
    return &AppmeshCreateMeshResult{Result: future}
}

func (a *AppMeshStub) CreateRoute(ctx workflow.Context, input *appmesh.CreateRouteInput) (*appmesh.CreateRouteOutput, error) {
    var output appmesh.CreateRouteOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.CreateRoute", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) CreateRouteAsync(ctx workflow.Context, input *appmesh.CreateRouteInput) *AppmeshCreateRouteResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.CreateRoute", input)
    return &AppmeshCreateRouteResult{Result: future}
}

func (a *AppMeshStub) CreateVirtualGateway(ctx workflow.Context, input *appmesh.CreateVirtualGatewayInput) (*appmesh.CreateVirtualGatewayOutput, error) {
    var output appmesh.CreateVirtualGatewayOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.CreateVirtualGateway", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) CreateVirtualGatewayAsync(ctx workflow.Context, input *appmesh.CreateVirtualGatewayInput) *AppmeshCreateVirtualGatewayResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.CreateVirtualGateway", input)
    return &AppmeshCreateVirtualGatewayResult{Result: future}
}

func (a *AppMeshStub) CreateVirtualNode(ctx workflow.Context, input *appmesh.CreateVirtualNodeInput) (*appmesh.CreateVirtualNodeOutput, error) {
    var output appmesh.CreateVirtualNodeOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.CreateVirtualNode", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) CreateVirtualNodeAsync(ctx workflow.Context, input *appmesh.CreateVirtualNodeInput) *AppmeshCreateVirtualNodeResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.CreateVirtualNode", input)
    return &AppmeshCreateVirtualNodeResult{Result: future}
}

func (a *AppMeshStub) CreateVirtualRouter(ctx workflow.Context, input *appmesh.CreateVirtualRouterInput) (*appmesh.CreateVirtualRouterOutput, error) {
    var output appmesh.CreateVirtualRouterOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.CreateVirtualRouter", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) CreateVirtualRouterAsync(ctx workflow.Context, input *appmesh.CreateVirtualRouterInput) *AppmeshCreateVirtualRouterResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.CreateVirtualRouter", input)
    return &AppmeshCreateVirtualRouterResult{Result: future}
}

func (a *AppMeshStub) CreateVirtualService(ctx workflow.Context, input *appmesh.CreateVirtualServiceInput) (*appmesh.CreateVirtualServiceOutput, error) {
    var output appmesh.CreateVirtualServiceOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.CreateVirtualService", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) CreateVirtualServiceAsync(ctx workflow.Context, input *appmesh.CreateVirtualServiceInput) *AppmeshCreateVirtualServiceResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.CreateVirtualService", input)
    return &AppmeshCreateVirtualServiceResult{Result: future}
}

func (a *AppMeshStub) DeleteGatewayRoute(ctx workflow.Context, input *appmesh.DeleteGatewayRouteInput) (*appmesh.DeleteGatewayRouteOutput, error) {
    var output appmesh.DeleteGatewayRouteOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DeleteGatewayRoute", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DeleteGatewayRouteAsync(ctx workflow.Context, input *appmesh.DeleteGatewayRouteInput) *AppmeshDeleteGatewayRouteResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DeleteGatewayRoute", input)
    return &AppmeshDeleteGatewayRouteResult{Result: future}
}

func (a *AppMeshStub) DeleteMesh(ctx workflow.Context, input *appmesh.DeleteMeshInput) (*appmesh.DeleteMeshOutput, error) {
    var output appmesh.DeleteMeshOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DeleteMesh", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DeleteMeshAsync(ctx workflow.Context, input *appmesh.DeleteMeshInput) *AppmeshDeleteMeshResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DeleteMesh", input)
    return &AppmeshDeleteMeshResult{Result: future}
}

func (a *AppMeshStub) DeleteRoute(ctx workflow.Context, input *appmesh.DeleteRouteInput) (*appmesh.DeleteRouteOutput, error) {
    var output appmesh.DeleteRouteOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DeleteRoute", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DeleteRouteAsync(ctx workflow.Context, input *appmesh.DeleteRouteInput) *AppmeshDeleteRouteResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DeleteRoute", input)
    return &AppmeshDeleteRouteResult{Result: future}
}

func (a *AppMeshStub) DeleteVirtualGateway(ctx workflow.Context, input *appmesh.DeleteVirtualGatewayInput) (*appmesh.DeleteVirtualGatewayOutput, error) {
    var output appmesh.DeleteVirtualGatewayOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DeleteVirtualGateway", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DeleteVirtualGatewayAsync(ctx workflow.Context, input *appmesh.DeleteVirtualGatewayInput) *AppmeshDeleteVirtualGatewayResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DeleteVirtualGateway", input)
    return &AppmeshDeleteVirtualGatewayResult{Result: future}
}

func (a *AppMeshStub) DeleteVirtualNode(ctx workflow.Context, input *appmesh.DeleteVirtualNodeInput) (*appmesh.DeleteVirtualNodeOutput, error) {
    var output appmesh.DeleteVirtualNodeOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DeleteVirtualNode", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DeleteVirtualNodeAsync(ctx workflow.Context, input *appmesh.DeleteVirtualNodeInput) *AppmeshDeleteVirtualNodeResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DeleteVirtualNode", input)
    return &AppmeshDeleteVirtualNodeResult{Result: future}
}

func (a *AppMeshStub) DeleteVirtualRouter(ctx workflow.Context, input *appmesh.DeleteVirtualRouterInput) (*appmesh.DeleteVirtualRouterOutput, error) {
    var output appmesh.DeleteVirtualRouterOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DeleteVirtualRouter", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DeleteVirtualRouterAsync(ctx workflow.Context, input *appmesh.DeleteVirtualRouterInput) *AppmeshDeleteVirtualRouterResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DeleteVirtualRouter", input)
    return &AppmeshDeleteVirtualRouterResult{Result: future}
}

func (a *AppMeshStub) DeleteVirtualService(ctx workflow.Context, input *appmesh.DeleteVirtualServiceInput) (*appmesh.DeleteVirtualServiceOutput, error) {
    var output appmesh.DeleteVirtualServiceOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DeleteVirtualService", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DeleteVirtualServiceAsync(ctx workflow.Context, input *appmesh.DeleteVirtualServiceInput) *AppmeshDeleteVirtualServiceResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DeleteVirtualService", input)
    return &AppmeshDeleteVirtualServiceResult{Result: future}
}

func (a *AppMeshStub) DescribeGatewayRoute(ctx workflow.Context, input *appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error) {
    var output appmesh.DescribeGatewayRouteOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DescribeGatewayRoute", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DescribeGatewayRouteAsync(ctx workflow.Context, input *appmesh.DescribeGatewayRouteInput) *AppmeshDescribeGatewayRouteResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DescribeGatewayRoute", input)
    return &AppmeshDescribeGatewayRouteResult{Result: future}
}

func (a *AppMeshStub) DescribeMesh(ctx workflow.Context, input *appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error) {
    var output appmesh.DescribeMeshOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DescribeMesh", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DescribeMeshAsync(ctx workflow.Context, input *appmesh.DescribeMeshInput) *AppmeshDescribeMeshResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DescribeMesh", input)
    return &AppmeshDescribeMeshResult{Result: future}
}

func (a *AppMeshStub) DescribeRoute(ctx workflow.Context, input *appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error) {
    var output appmesh.DescribeRouteOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DescribeRoute", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DescribeRouteAsync(ctx workflow.Context, input *appmesh.DescribeRouteInput) *AppmeshDescribeRouteResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DescribeRoute", input)
    return &AppmeshDescribeRouteResult{Result: future}
}

func (a *AppMeshStub) DescribeVirtualGateway(ctx workflow.Context, input *appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error) {
    var output appmesh.DescribeVirtualGatewayOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DescribeVirtualGateway", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DescribeVirtualGatewayAsync(ctx workflow.Context, input *appmesh.DescribeVirtualGatewayInput) *AppmeshDescribeVirtualGatewayResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DescribeVirtualGateway", input)
    return &AppmeshDescribeVirtualGatewayResult{Result: future}
}

func (a *AppMeshStub) DescribeVirtualNode(ctx workflow.Context, input *appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error) {
    var output appmesh.DescribeVirtualNodeOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DescribeVirtualNode", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DescribeVirtualNodeAsync(ctx workflow.Context, input *appmesh.DescribeVirtualNodeInput) *AppmeshDescribeVirtualNodeResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DescribeVirtualNode", input)
    return &AppmeshDescribeVirtualNodeResult{Result: future}
}

func (a *AppMeshStub) DescribeVirtualRouter(ctx workflow.Context, input *appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error) {
    var output appmesh.DescribeVirtualRouterOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DescribeVirtualRouter", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DescribeVirtualRouterAsync(ctx workflow.Context, input *appmesh.DescribeVirtualRouterInput) *AppmeshDescribeVirtualRouterResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DescribeVirtualRouter", input)
    return &AppmeshDescribeVirtualRouterResult{Result: future}
}

func (a *AppMeshStub) DescribeVirtualService(ctx workflow.Context, input *appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error) {
    var output appmesh.DescribeVirtualServiceOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.DescribeVirtualService", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) DescribeVirtualServiceAsync(ctx workflow.Context, input *appmesh.DescribeVirtualServiceInput) *AppmeshDescribeVirtualServiceResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.DescribeVirtualService", input)
    return &AppmeshDescribeVirtualServiceResult{Result: future}
}

func (a *AppMeshStub) ListGatewayRoutes(ctx workflow.Context, input *appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error) {
    var output appmesh.ListGatewayRoutesOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.ListGatewayRoutes", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) ListGatewayRoutesAsync(ctx workflow.Context, input *appmesh.ListGatewayRoutesInput) *AppmeshListGatewayRoutesResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.ListGatewayRoutes", input)
    return &AppmeshListGatewayRoutesResult{Result: future}
}

func (a *AppMeshStub) ListMeshes(ctx workflow.Context, input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error) {
    var output appmesh.ListMeshesOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.ListMeshes", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) ListMeshesAsync(ctx workflow.Context, input *appmesh.ListMeshesInput) *AppmeshListMeshesResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.ListMeshes", input)
    return &AppmeshListMeshesResult{Result: future}
}

func (a *AppMeshStub) ListRoutes(ctx workflow.Context, input *appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error) {
    var output appmesh.ListRoutesOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.ListRoutes", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) ListRoutesAsync(ctx workflow.Context, input *appmesh.ListRoutesInput) *AppmeshListRoutesResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.ListRoutes", input)
    return &AppmeshListRoutesResult{Result: future}
}

func (a *AppMeshStub) ListTagsForResource(ctx workflow.Context, input *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error) {
    var output appmesh.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) ListTagsForResourceAsync(ctx workflow.Context, input *appmesh.ListTagsForResourceInput) *AppmeshListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.ListTagsForResource", input)
    return &AppmeshListTagsForResourceResult{Result: future}
}

func (a *AppMeshStub) ListVirtualGateways(ctx workflow.Context, input *appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error) {
    var output appmesh.ListVirtualGatewaysOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.ListVirtualGateways", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) ListVirtualGatewaysAsync(ctx workflow.Context, input *appmesh.ListVirtualGatewaysInput) *AppmeshListVirtualGatewaysResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.ListVirtualGateways", input)
    return &AppmeshListVirtualGatewaysResult{Result: future}
}

func (a *AppMeshStub) ListVirtualNodes(ctx workflow.Context, input *appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error) {
    var output appmesh.ListVirtualNodesOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.ListVirtualNodes", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) ListVirtualNodesAsync(ctx workflow.Context, input *appmesh.ListVirtualNodesInput) *AppmeshListVirtualNodesResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.ListVirtualNodes", input)
    return &AppmeshListVirtualNodesResult{Result: future}
}

func (a *AppMeshStub) ListVirtualRouters(ctx workflow.Context, input *appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error) {
    var output appmesh.ListVirtualRoutersOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.ListVirtualRouters", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) ListVirtualRoutersAsync(ctx workflow.Context, input *appmesh.ListVirtualRoutersInput) *AppmeshListVirtualRoutersResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.ListVirtualRouters", input)
    return &AppmeshListVirtualRoutersResult{Result: future}
}

func (a *AppMeshStub) ListVirtualServices(ctx workflow.Context, input *appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error) {
    var output appmesh.ListVirtualServicesOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.ListVirtualServices", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) ListVirtualServicesAsync(ctx workflow.Context, input *appmesh.ListVirtualServicesInput) *AppmeshListVirtualServicesResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.ListVirtualServices", input)
    return &AppmeshListVirtualServicesResult{Result: future}
}

func (a *AppMeshStub) TagResource(ctx workflow.Context, input *appmesh.TagResourceInput) (*appmesh.TagResourceOutput, error) {
    var output appmesh.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) TagResourceAsync(ctx workflow.Context, input *appmesh.TagResourceInput) *AppmeshTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.TagResource", input)
    return &AppmeshTagResourceResult{Result: future}
}

func (a *AppMeshStub) UntagResource(ctx workflow.Context, input *appmesh.UntagResourceInput) (*appmesh.UntagResourceOutput, error) {
    var output appmesh.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) UntagResourceAsync(ctx workflow.Context, input *appmesh.UntagResourceInput) *AppmeshUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.UntagResource", input)
    return &AppmeshUntagResourceResult{Result: future}
}

func (a *AppMeshStub) UpdateGatewayRoute(ctx workflow.Context, input *appmesh.UpdateGatewayRouteInput) (*appmesh.UpdateGatewayRouteOutput, error) {
    var output appmesh.UpdateGatewayRouteOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.UpdateGatewayRoute", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) UpdateGatewayRouteAsync(ctx workflow.Context, input *appmesh.UpdateGatewayRouteInput) *AppmeshUpdateGatewayRouteResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.UpdateGatewayRoute", input)
    return &AppmeshUpdateGatewayRouteResult{Result: future}
}

func (a *AppMeshStub) UpdateMesh(ctx workflow.Context, input *appmesh.UpdateMeshInput) (*appmesh.UpdateMeshOutput, error) {
    var output appmesh.UpdateMeshOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.UpdateMesh", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) UpdateMeshAsync(ctx workflow.Context, input *appmesh.UpdateMeshInput) *AppmeshUpdateMeshResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.UpdateMesh", input)
    return &AppmeshUpdateMeshResult{Result: future}
}

func (a *AppMeshStub) UpdateRoute(ctx workflow.Context, input *appmesh.UpdateRouteInput) (*appmesh.UpdateRouteOutput, error) {
    var output appmesh.UpdateRouteOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.UpdateRoute", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) UpdateRouteAsync(ctx workflow.Context, input *appmesh.UpdateRouteInput) *AppmeshUpdateRouteResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.UpdateRoute", input)
    return &AppmeshUpdateRouteResult{Result: future}
}

func (a *AppMeshStub) UpdateVirtualGateway(ctx workflow.Context, input *appmesh.UpdateVirtualGatewayInput) (*appmesh.UpdateVirtualGatewayOutput, error) {
    var output appmesh.UpdateVirtualGatewayOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.UpdateVirtualGateway", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) UpdateVirtualGatewayAsync(ctx workflow.Context, input *appmesh.UpdateVirtualGatewayInput) *AppmeshUpdateVirtualGatewayResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.UpdateVirtualGateway", input)
    return &AppmeshUpdateVirtualGatewayResult{Result: future}
}

func (a *AppMeshStub) UpdateVirtualNode(ctx workflow.Context, input *appmesh.UpdateVirtualNodeInput) (*appmesh.UpdateVirtualNodeOutput, error) {
    var output appmesh.UpdateVirtualNodeOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.UpdateVirtualNode", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) UpdateVirtualNodeAsync(ctx workflow.Context, input *appmesh.UpdateVirtualNodeInput) *AppmeshUpdateVirtualNodeResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.UpdateVirtualNode", input)
    return &AppmeshUpdateVirtualNodeResult{Result: future}
}

func (a *AppMeshStub) UpdateVirtualRouter(ctx workflow.Context, input *appmesh.UpdateVirtualRouterInput) (*appmesh.UpdateVirtualRouterOutput, error) {
    var output appmesh.UpdateVirtualRouterOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.UpdateVirtualRouter", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) UpdateVirtualRouterAsync(ctx workflow.Context, input *appmesh.UpdateVirtualRouterInput) *AppmeshUpdateVirtualRouterResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.UpdateVirtualRouter", input)
    return &AppmeshUpdateVirtualRouterResult{Result: future}
}

func (a *AppMeshStub) UpdateVirtualService(ctx workflow.Context, input *appmesh.UpdateVirtualServiceInput) (*appmesh.UpdateVirtualServiceOutput, error) {
    var output appmesh.UpdateVirtualServiceOutput
    err := workflow.ExecuteActivity(ctx, "AppMesh.UpdateVirtualService", input).Get(ctx, &output)
    return &output, err
}

func (a *AppMeshStub) UpdateVirtualServiceAsync(ctx workflow.Context, input *appmesh.UpdateVirtualServiceInput) *AppmeshUpdateVirtualServiceResult {
    future := workflow.ExecuteActivity(ctx, "AppMesh.UpdateVirtualService", input)
    return &AppmeshUpdateVirtualServiceResult{Result: future}
}
