package gen_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter/gen"
)

func TestInt64Shrink(t *testing.T) {
	zeroShrinks := gen.Int64Shrinker(int64(0)).All()
	if !reflect.DeepEqual(zeroShrinks, []interface{}{}) {
		t.Errorf("Invalid zeroShrinks: %#v", zeroShrinks)
	}

	tenShrinks := gen.Int64Shrinker(int64(10)).All()
	if !reflect.DeepEqual(tenShrinks, []interface{}{int64(0), int64(5), int64(-5), int64(8), int64(-8), int64(9), int64(-9)}) {
		t.Errorf("Invalid tenShrinks: %#v", tenShrinks)
	}

	negTenShrinks := gen.Int64Shrinker(int64(-10)).All()
	if !reflect.DeepEqual(negTenShrinks, []interface{}{int64(0), int64(-5), int64(5), int64(-8), int64(8), int64(-9), int64(9)}) {
		t.Errorf("Invalid negTenShrinks: %#v", negTenShrinks)
	}

	leetShrink := gen.Int64Shrinker(int64(1337)).All()
	if !reflect.DeepEqual(leetShrink, []interface{}{
		int64(0), int64(669), int64(-669), int64(1003), int64(-1003), int64(1170), int64(-1170),
		int64(1254), int64(-1254), int64(1296), int64(-1296), int64(1317), int64(-1317),
		int64(1327), int64(-1327), int64(1332), int64(-1332), int64(1335), int64(-1335),
		int64(1336), int64(-1336)}) {
		t.Errorf("Invalid leetShrink: %#v", leetShrink)
	}
}
