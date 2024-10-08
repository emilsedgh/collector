package util

import "time"

func StringPtrToString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func BoolPtrToBool(ptr *bool) bool {
	if ptr == nil {
		return false
	}
	return *ptr
}

func TimePtrToUnixTimestamp(ptr *time.Time) int64 {
	if ptr == nil {
		return 0
	}
	return ptr.Unix()
}

func TimePtrToTime(ptr *time.Time) time.Time {
	if ptr == nil {
		return time.Time{}
	}
	return *ptr
}

func IntPtrToInt(ptr *int64) int64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

func Int32PtrToInt(ptr *int32) int32 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

func StringCustomTypePtrToString[T ~string](ptr *T) string {
	if ptr == nil {
		return ""
	}
	return string(*ptr)
}
