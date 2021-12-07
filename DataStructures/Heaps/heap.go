package heaps

import "fmt"

//This structure is a complete three. It means that all the levels of this three are complete(They have 2 children)
//In case of a MaxHeap the largest key would be stored in the root of three. Every parent have a higher key than a choldren keys

//Example:
//				50
//		16				48
//	14		8 		34		20

//Min Heap - the root has a smallest key(Opposite structure)

//We can imagine heap as a arrya

//Index 0	1	2	3	4	5	6
//Key   50  16  48  14  8   34  20

//Calculating child index
//LeftChildIndex = ParentIndex * 2 + 1      RightChildIndex = ParentIndex * 2 + 2

//Examples
//Parrent = 16[1],  LeftChildIndex = 1 * 2 + 1 = 3 RightChildIndex = 1 * 2 + 2 = 4

//INSERT INTO HEAP
//1.Add new value to the bottom of the heap
//2.Rearange the heap

//Extract the root from Heap
//1.Extract(Delete) the root
//2.Rearange the heap
//1.Add last key to the root
//2.Swap if needed starting from the top

//Heap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

//Insert adds element to the heap
func (h *MaxHeap) Insert(key int) {
	//Insert key
	h.array = append(h.array, key)

	//Rearange heap
	h.maxHeapifyUp(len(h.array) - 1)
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parrent(index)] < h.array[index] {
		h.swap(parrent(index), index)
		index = parrent(index)
	}
}

//Get the parrent index
func parrent(i int) int {
	return (i - 1) / 2
}

// get left child index
func left(i int) int {
	return 2*i + 1
}

// get right child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

//Extract return the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	lenght := len(h.array) - 1

	//When the array is empty
	if len(h.array) == 0 {
		fmt.Println("Cant extract because array lenght is 0")
		return -1
	}

	//Take out the last index and put in in the root
	h.array[0] = h.array[lenght]
	h.array = h.array[:lenght]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childCompare := 0
	//loop while index has at least one child

	for l <= lastIndex {
		if l == lastIndex { // When left child is the only child
			childCompare = l
		} else if h.array[l] > h.array[r] { //When left child is larger
			childCompare = l
		} else { //When right child is larger
			childCompare = r
		}
		//compare array value of current index to larger and swap if smaller
		if h.array[index] < h.array[childCompare] {
			h.swap(index, childCompare)
			index = childCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}
