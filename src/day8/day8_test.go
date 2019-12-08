package day8

import (
    "testing"
    "reflect"
)

func TestExec(t *testing.T) {
    tables := []struct {
		imageData []int
        width int
        height int
		layers []Layer
	}{
        {[]int{1,2,3,4,5,6,7,8,9,0,1,2}, 3, 2,
         []Layer{
             Layer{[]int{1,2,3,4,5,6}, 3, 2},
             Layer{[]int{7,8,9,0,1,2}, 3, 2},
         }},
	}

    for _, table := range tables {
        got := imageToLayers(table.width, table.height, table.imageData)
        if !reflect.DeepEqual(got, table.layers) {
            t.Errorf("result was incorrect, got: %v, want: %v",
                got, table.layers)
        }
    }
}

