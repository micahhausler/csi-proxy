// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-csi/csi-proxy/client/api/filesystem/v1alpha1"
	impl "github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl"
)

func autoConvert_v1alpha1_IsMountPointRequest_To_impl_IsMountPointRequest(in *v1alpha1.IsMountPointRequest, out *impl.IsMountPointRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_v1alpha1_IsMountPointRequest_To_impl_IsMountPointRequest is an autogenerated conversion function.
func Convert_v1alpha1_IsMountPointRequest_To_impl_IsMountPointRequest(in *v1alpha1.IsMountPointRequest, out *impl.IsMountPointRequest) error {
	return autoConvert_v1alpha1_IsMountPointRequest_To_impl_IsMountPointRequest(in, out)
}

func autoConvert_impl_IsMountPointRequest_To_v1alpha1_IsMountPointRequest(in *impl.IsMountPointRequest, out *v1alpha1.IsMountPointRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_impl_IsMountPointRequest_To_v1alpha1_IsMountPointRequest is an autogenerated conversion function.
func Convert_impl_IsMountPointRequest_To_v1alpha1_IsMountPointRequest(in *impl.IsMountPointRequest, out *v1alpha1.IsMountPointRequest) error {
	return autoConvert_impl_IsMountPointRequest_To_v1alpha1_IsMountPointRequest(in, out)
}

func autoConvert_v1alpha1_IsMountPointResponse_To_impl_IsMountPointResponse(in *v1alpha1.IsMountPointResponse, out *impl.IsMountPointResponse) error {
	out.IsMountPoint = in.IsMountPoint
	return nil
}

// Convert_v1alpha1_IsMountPointResponse_To_impl_IsMountPointResponse is an autogenerated conversion function.
func Convert_v1alpha1_IsMountPointResponse_To_impl_IsMountPointResponse(in *v1alpha1.IsMountPointResponse, out *impl.IsMountPointResponse) error {
	return autoConvert_v1alpha1_IsMountPointResponse_To_impl_IsMountPointResponse(in, out)
}

func autoConvert_impl_IsMountPointResponse_To_v1alpha1_IsMountPointResponse(in *impl.IsMountPointResponse, out *v1alpha1.IsMountPointResponse) error {
	out.IsMountPoint = in.IsMountPoint
	return nil
}

// Convert_impl_IsMountPointResponse_To_v1alpha1_IsMountPointResponse is an autogenerated conversion function.
func Convert_impl_IsMountPointResponse_To_v1alpha1_IsMountPointResponse(in *impl.IsMountPointResponse, out *v1alpha1.IsMountPointResponse) error {
	return autoConvert_impl_IsMountPointResponse_To_v1alpha1_IsMountPointResponse(in, out)
}

func autoConvert_v1alpha1_LinkPathRequest_To_impl_LinkPathRequest(in *v1alpha1.LinkPathRequest, out *impl.LinkPathRequest) error {
	out.SourcePath = in.SourcePath
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_v1alpha1_LinkPathRequest_To_impl_LinkPathRequest is an autogenerated conversion function.
func Convert_v1alpha1_LinkPathRequest_To_impl_LinkPathRequest(in *v1alpha1.LinkPathRequest, out *impl.LinkPathRequest) error {
	return autoConvert_v1alpha1_LinkPathRequest_To_impl_LinkPathRequest(in, out)
}

func autoConvert_impl_LinkPathRequest_To_v1alpha1_LinkPathRequest(in *impl.LinkPathRequest, out *v1alpha1.LinkPathRequest) error {
	out.SourcePath = in.SourcePath
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_impl_LinkPathRequest_To_v1alpha1_LinkPathRequest is an autogenerated conversion function.
func Convert_impl_LinkPathRequest_To_v1alpha1_LinkPathRequest(in *impl.LinkPathRequest, out *v1alpha1.LinkPathRequest) error {
	return autoConvert_impl_LinkPathRequest_To_v1alpha1_LinkPathRequest(in, out)
}

func autoConvert_v1alpha1_LinkPathResponse_To_impl_LinkPathResponse(in *v1alpha1.LinkPathResponse, out *impl.LinkPathResponse) error {
	return nil
}

// Convert_v1alpha1_LinkPathResponse_To_impl_LinkPathResponse is an autogenerated conversion function.
func Convert_v1alpha1_LinkPathResponse_To_impl_LinkPathResponse(in *v1alpha1.LinkPathResponse, out *impl.LinkPathResponse) error {
	return autoConvert_v1alpha1_LinkPathResponse_To_impl_LinkPathResponse(in, out)
}

func autoConvert_impl_LinkPathResponse_To_v1alpha1_LinkPathResponse(in *impl.LinkPathResponse, out *v1alpha1.LinkPathResponse) error {
	return nil
}

// Convert_impl_LinkPathResponse_To_v1alpha1_LinkPathResponse is an autogenerated conversion function.
func Convert_impl_LinkPathResponse_To_v1alpha1_LinkPathResponse(in *impl.LinkPathResponse, out *v1alpha1.LinkPathResponse) error {
	return autoConvert_impl_LinkPathResponse_To_v1alpha1_LinkPathResponse(in, out)
}

func autoConvert_v1alpha1_MkdirRequest_To_impl_MkdirRequest(in *v1alpha1.MkdirRequest, out *impl.MkdirRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_v1alpha1_MkdirRequest_To_impl_MkdirRequest is an autogenerated conversion function.
func Convert_v1alpha1_MkdirRequest_To_impl_MkdirRequest(in *v1alpha1.MkdirRequest, out *impl.MkdirRequest) error {
	return autoConvert_v1alpha1_MkdirRequest_To_impl_MkdirRequest(in, out)
}

func autoConvert_impl_MkdirRequest_To_v1alpha1_MkdirRequest(in *impl.MkdirRequest, out *v1alpha1.MkdirRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_impl_MkdirRequest_To_v1alpha1_MkdirRequest is an autogenerated conversion function.
func Convert_impl_MkdirRequest_To_v1alpha1_MkdirRequest(in *impl.MkdirRequest, out *v1alpha1.MkdirRequest) error {
	return autoConvert_impl_MkdirRequest_To_v1alpha1_MkdirRequest(in, out)
}

func autoConvert_v1alpha1_MkdirResponse_To_impl_MkdirResponse(in *v1alpha1.MkdirResponse, out *impl.MkdirResponse) error {
	return nil
}

// Convert_v1alpha1_MkdirResponse_To_impl_MkdirResponse is an autogenerated conversion function.
func Convert_v1alpha1_MkdirResponse_To_impl_MkdirResponse(in *v1alpha1.MkdirResponse, out *impl.MkdirResponse) error {
	return autoConvert_v1alpha1_MkdirResponse_To_impl_MkdirResponse(in, out)
}

func autoConvert_impl_MkdirResponse_To_v1alpha1_MkdirResponse(in *impl.MkdirResponse, out *v1alpha1.MkdirResponse) error {
	return nil
}

// Convert_impl_MkdirResponse_To_v1alpha1_MkdirResponse is an autogenerated conversion function.
func Convert_impl_MkdirResponse_To_v1alpha1_MkdirResponse(in *impl.MkdirResponse, out *v1alpha1.MkdirResponse) error {
	return autoConvert_impl_MkdirResponse_To_v1alpha1_MkdirResponse(in, out)
}

func autoConvert_v1alpha1_PathExistsRequest_To_impl_PathExistsRequest(in *v1alpha1.PathExistsRequest, out *impl.PathExistsRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_v1alpha1_PathExistsRequest_To_impl_PathExistsRequest is an autogenerated conversion function.
func Convert_v1alpha1_PathExistsRequest_To_impl_PathExistsRequest(in *v1alpha1.PathExistsRequest, out *impl.PathExistsRequest) error {
	return autoConvert_v1alpha1_PathExistsRequest_To_impl_PathExistsRequest(in, out)
}

func autoConvert_impl_PathExistsRequest_To_v1alpha1_PathExistsRequest(in *impl.PathExistsRequest, out *v1alpha1.PathExistsRequest) error {
	out.Path = in.Path
	return nil
}

// Convert_impl_PathExistsRequest_To_v1alpha1_PathExistsRequest is an autogenerated conversion function.
func Convert_impl_PathExistsRequest_To_v1alpha1_PathExistsRequest(in *impl.PathExistsRequest, out *v1alpha1.PathExistsRequest) error {
	return autoConvert_impl_PathExistsRequest_To_v1alpha1_PathExistsRequest(in, out)
}

func autoConvert_v1alpha1_PathExistsResponse_To_impl_PathExistsResponse(in *v1alpha1.PathExistsResponse, out *impl.PathExistsResponse) error {
	out.Exists = in.Exists
	return nil
}

// Convert_v1alpha1_PathExistsResponse_To_impl_PathExistsResponse is an autogenerated conversion function.
func Convert_v1alpha1_PathExistsResponse_To_impl_PathExistsResponse(in *v1alpha1.PathExistsResponse, out *impl.PathExistsResponse) error {
	return autoConvert_v1alpha1_PathExistsResponse_To_impl_PathExistsResponse(in, out)
}

func autoConvert_impl_PathExistsResponse_To_v1alpha1_PathExistsResponse(in *impl.PathExistsResponse, out *v1alpha1.PathExistsResponse) error {
	out.Exists = in.Exists
	return nil
}

// Convert_impl_PathExistsResponse_To_v1alpha1_PathExistsResponse is an autogenerated conversion function.
func Convert_impl_PathExistsResponse_To_v1alpha1_PathExistsResponse(in *impl.PathExistsResponse, out *v1alpha1.PathExistsResponse) error {
	return autoConvert_impl_PathExistsResponse_To_v1alpha1_PathExistsResponse(in, out)
}

func autoConvert_v1alpha1_RmdirRequest_To_impl_RmdirRequest(in *v1alpha1.RmdirRequest, out *impl.RmdirRequest) error {
	out.Path = in.Path
	out.Force = in.Force
	return nil
}

// Convert_v1alpha1_RmdirRequest_To_impl_RmdirRequest is an autogenerated conversion function.
func Convert_v1alpha1_RmdirRequest_To_impl_RmdirRequest(in *v1alpha1.RmdirRequest, out *impl.RmdirRequest) error {
	return autoConvert_v1alpha1_RmdirRequest_To_impl_RmdirRequest(in, out)
}

func autoConvert_impl_RmdirRequest_To_v1alpha1_RmdirRequest(in *impl.RmdirRequest, out *v1alpha1.RmdirRequest) error {
	out.Path = in.Path
	out.Force = in.Force
	return nil
}

// Convert_impl_RmdirRequest_To_v1alpha1_RmdirRequest is an autogenerated conversion function.
func Convert_impl_RmdirRequest_To_v1alpha1_RmdirRequest(in *impl.RmdirRequest, out *v1alpha1.RmdirRequest) error {
	return autoConvert_impl_RmdirRequest_To_v1alpha1_RmdirRequest(in, out)
}

func autoConvert_v1alpha1_RmdirResponse_To_impl_RmdirResponse(in *v1alpha1.RmdirResponse, out *impl.RmdirResponse) error {
	return nil
}

// Convert_v1alpha1_RmdirResponse_To_impl_RmdirResponse is an autogenerated conversion function.
func Convert_v1alpha1_RmdirResponse_To_impl_RmdirResponse(in *v1alpha1.RmdirResponse, out *impl.RmdirResponse) error {
	return autoConvert_v1alpha1_RmdirResponse_To_impl_RmdirResponse(in, out)
}

func autoConvert_impl_RmdirResponse_To_v1alpha1_RmdirResponse(in *impl.RmdirResponse, out *v1alpha1.RmdirResponse) error {
	return nil
}

// Convert_impl_RmdirResponse_To_v1alpha1_RmdirResponse is an autogenerated conversion function.
func Convert_impl_RmdirResponse_To_v1alpha1_RmdirResponse(in *impl.RmdirResponse, out *v1alpha1.RmdirResponse) error {
	return autoConvert_impl_RmdirResponse_To_v1alpha1_RmdirResponse(in, out)
}