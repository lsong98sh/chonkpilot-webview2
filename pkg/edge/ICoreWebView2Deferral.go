package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// ICoreWebView2Deferral 是 WebView2 的「延迟完成」句柄。
//
// 用途：WebResourceRequested 回调默认在 **UI 线程**上同步执行；若在回调内做耗时工作
// （总线请求-响应、数据面读写等），整个窗口的 UI 线程会被阻塞 → 页面所有网络请求连带
// 排队（表现为「界面卡住、fetch 全部超时」）。GetDeferral() 取得本对象后，回调可以**先
// 返回**，在后台完成工作，再在 UI 线程 PutResponse + Complete()。
type _ICoreWebView2DeferralVtbl struct {
	_IUnknownVtbl
	Complete ComProc
}

type ICoreWebView2Deferral struct {
	vtbl *_ICoreWebView2DeferralVtbl
}

func (i *ICoreWebView2Deferral) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Deferral) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

// Complete 结束延迟：调用后 WebView2 认为本次请求处理完毕（必须在 UI 线程、且已
// PutResponse）。
func (i *ICoreWebView2Deferral) Complete() error {
	_, _, err := i.vtbl.Complete.Call(uintptr(unsafe.Pointer(i)))
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
