package main

import "unsafe"

// growslice 用于切片扩容，是 Go runtime 内部在执行 append 时可能会调用的函数。
// 参数：
//
//	et   - slice 元素的类型信息（_type 是 Go runtime 内部的类型描述结构）
//	old  - 旧的 slice（包含底层数组指针、长度、容量）
//	cap  - 新 slice 的最低要求容量（旧长度 + append 添加的元素个数）
//
// 举例：
//
//	s := []int{1, 2, 3}
//	s = append(s, 4, 5)
//	此时：
//	  old.len = 3
//	  old.cap = 3
//	  cap = 5 (因为需要容纳原来的3个元素 + 追加的2个元素)
func growslice(et *_type, old slice, cap int) slice {
	// 如果要求的容量比旧的容量还小，说明调用方逻辑出错
	if cap < old.cap {
		panic(errorString("growslice: cap out of range"))
	}

	// 如果元素大小是 0（例如 struct{}），直接返回一个新的 slice，
	// 因为不需要实际分配空间。
	if et.size == 0 {
		return slice{unsafe.Pointer(&zerobase), old.len, cap}
	}

	// ====== 第一阶段：计算 newcap（扩容后的容量） ======
	newcap := old.cap
	doublecap := newcap + newcap // 旧容量的两倍
	if cap > doublecap {
		// 情况 1：最低需求 > 旧容量的两倍 => 直接使用需求值
		newcap = cap
	} else {
		if old.len < 1024 {
			// 情况 2：长度小于 1024，直接扩容一倍（性能优先）
			newcap = doublecap
		} else {
			// 情况 3：长度 >= 1024，逐步扩容 25%（内存利用率优先）
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			if newcap <= 0 { // 防止溢出
				newcap = cap
			}
		}
	}

	// ====== 第二阶段：计算内存大小并检查溢出 ======
	var overflow bool
	var lenmem, newlenmem, capmem uintptr
	switch {
	case et.size == 1:
		// 元素大小为 1 字节（如 byte）
		lenmem = uintptr(old.len)
		newlenmem = uintptr(cap)
		capmem = roundupsize(uintptr(newcap)) // 向上取整到适合内存分配器的大小
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem) // 调整 newcap
	case et.size == sys.PtrSize:
		// 元素大小等于指针大小（4 字节或 8 字节）
		lenmem = uintptr(old.len) * sys.PtrSize
		newlenmem = uintptr(cap) * sys.PtrSize
		capmem = roundupsize(uintptr(newcap) * sys.PtrSize)
		overflow = uintptr(newcap) > maxAlloc/sys.PtrSize
		newcap = int(capmem / sys.PtrSize)
	case isPowerOfTwo(et.size):
		// 元素大小是 2 的幂，使用位运算优化
		var shift uintptr
		if sys.PtrSize == 8 {
			shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
		}
		lenmem = uintptr(old.len) << shift
		newlenmem = uintptr(cap) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
	default:
		// 通用情况（元素大小不为 1、指针大小、或 2 的幂）
		lenmem = uintptr(old.len) * et.size
		newlenmem = uintptr(cap) * et.size
		capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.size)
	}

	// 溢出检查
	if overflow || capmem > maxAlloc {
		panic(errorString("growslice: cap out of range"))
	}

	// ====== 第三阶段：分配内存并拷贝旧数据 ======
	var p unsafe.Pointer
	if et.ptrdata == 0 {
		// 元素不含指针 => 可直接分配并清零尾部
		p = mallocgc(capmem, nil, false)
		memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
	} else {
		// 元素含指针 => 需要写屏障
		p = mallocgc(capmem, et, true)
		if lenmem > 0 && writeBarrier.enabled {
			bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(old.array), lenmem)
		}
	}

	// 将旧数据复制到新内存
	memmove(p, old.array, lenmem)

	// 返回新 slice（底层数组替换为新分配的数组）
	return slice{p, old.len, newcap}
}

func main() {

}
