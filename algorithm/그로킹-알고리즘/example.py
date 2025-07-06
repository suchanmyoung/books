# linear time > log time

def binary_search(arr, target):
    low = 0 
    high = len(arr) - 1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
    return -1

print('Binary Search: ' + str(binary_search([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 7)))


def findSmallest(arr):
    smallest = arr[0]
    smallest_index = 0
    for i in range(1, len(arr)):
        if arr[i] < smallest:
            smallest = arr[i]
            smallest_index = i
    return smallest_index

def selectionSort(arr):
    newArr = []
    for i in range(len(arr)):
        smallest = findSmallest(arr)
        newArr.append(arr.pop(smallest))
    return newArr

print('Selection Sort: ' + str(selectionSort([5, 3, 6, 2, 10])))