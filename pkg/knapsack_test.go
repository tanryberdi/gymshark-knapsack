package pkg

import "testing"

func TestKnapsack(t *testing.T) {
	type args struct {
		items    []int
		capacity int
	}
	tests := []struct {
		name           string
		args           args
		minItemsAmount int
		minPacksAmount int
		usedItemsMap   map[int]int
	}{
		{
			name: "case-1",
			args: args{
				items:    []int{23, 31, 53},
				capacity: 263,
			},
			minItemsAmount: 263,
			minPacksAmount: 9,
			usedItemsMap: map[int]int{
				23: 2,
				31: 7,
			},
		},
		{
			name: "case-2",
			args: args{
				items:    []int{250, 500, 1000, 2000, 5000},
				capacity: 1,
			},
			minItemsAmount: 250,
			minPacksAmount: 1,
			usedItemsMap: map[int]int{
				250: 1,
			},
		},
		{
			name: "case-3",
			args: args{
				items:    []int{250, 500, 1000, 2000, 5000},
				capacity: 250,
			},
			minItemsAmount: 250,
			minPacksAmount: 1,
			usedItemsMap: map[int]int{
				250: 1,
			},
		},
		{
			name: "case-4",
			args: args{
				items:    []int{250, 500, 1000, 2000, 5000},
				capacity: 251,
			},
			minItemsAmount: 500,
			minPacksAmount: 1,
			usedItemsMap: map[int]int{
				500: 1,
			},
		},
		{
			name: "case-5",
			args: args{
				items:    []int{250, 500, 1000, 2000, 5000},
				capacity: 501,
			},
			minItemsAmount: 750,
			minPacksAmount: 2,
			usedItemsMap: map[int]int{
				250: 1,
				500: 1,
			},
		},
		{
			name: "case-6",
			args: args{
				items:    []int{250, 500, 1000, 2000, 5000},
				capacity: 12001,
			},
			minItemsAmount: 12250,
			minPacksAmount: 4,
			usedItemsMap: map[int]int{
				250:  1,
				2000: 1,
				5000: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itemsAmount, packsAmount, itemsMap := Knapsack(tt.args.items, tt.args.capacity)
			if itemsAmount != tt.minItemsAmount {
				t.Errorf("Knapsack() got = %v, want %v", itemsAmount, tt.minItemsAmount)
			}
			if packsAmount != tt.minPacksAmount {
				t.Errorf("Knapsack() got1 = %v, want %v", packsAmount, tt.minPacksAmount)
			}
			for k, v := range tt.usedItemsMap {
				if itemsMap[k] != v {
					t.Errorf("Knapsack() got2 = %v, want %v", itemsMap, tt.usedItemsMap)
				}
			}
		})
	}
}
